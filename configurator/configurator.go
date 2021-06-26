package configurator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/imdario/mergo"
	gitopsv1alpha1 "github.com/keikoproj/cloudresource-manager/api/v1alpha1"

	"github.com/keikoproj/manny/config"
	"github.com/keikoproj/manny/configurator/mocks"
)

const (
	// MannyConfigName is a special name in manny that denotes the name of the configuration file
	MannyConfigName = "config.yaml"

	// EmptyResourceRegExp Regex details : Should start with "Resources".
	//Can have space in between.
	//Should have ":" delimiter.
	//Can have space in between.
	//Can have "{}" or not. Space allowed between curly braces.
	EmptyResourceRegExp = "^(Resources(\"{0,1})(\\s*):(\\s*)({(\\s*)}){0,1})$"
)

var counter int
var baseDir []string
var AcceptedExtensions = []string{".yaml", ".yml", ".json"}

type SL struct {
	// stack lists
	StackList []map[string]string
}

var s SL

type CloudResources []*CloudResourceDeployment

// Validate runs various validations against the deployments
func (c CloudResources) Validate() error {
	lookupTable := map[string]bool{}

	for _, resource := range c {
		// find duplicate stack names
		name := resource.Spec.Cloudformation.Stackname
		if lookupTable[name] {
			return errors.New("duplicate stack name found: " + name)
		}

		lookupTable[name] = true
	}

	return nil
}



// Render returns a json or yaml byte array of the resources
func (c CloudResources) Render(format string) (bytes []byte, err error) {
	switch format {
	case "json":
		for _, d := range c {
			deployment, err := json.Marshal(d)
			if err != nil {
				return nil, err
			}

			bytes = append(bytes, deployment...)
		}
	default:
		for _, d := range c {
			bytes = append(bytes, []byte("---\n")...)

			deployment, err := yaml.Marshal(d)
			if err != nil {
				return nil, err
			}

			bytes = append(bytes, deployment...)
		}
	}

	return
}

// CloudResourceDeployment is a Custom Resource duplicated from the GitOps controller.
// We had to duplicate it here because the built-in Kubernetes types are made for machine processing
// with JSON. In order to produce a YAML manifest that someone can read we have to represent that
// structure in our code.
type CloudResourceDeployment struct {
	Kind       string                                     `yaml:"kind" json:"kind"`
	APIVersion string                                     `yaml:"apiVersion" json:"apiVersion"`
	Metadata   Metadata                                   `yaml:"metadata" json:"metadata"`
	Spec       gitopsv1alpha1.CloudResourceDeploymentSpec `yaml:"spec" json:"spec"`
}

type Metadata struct {
	Name        string            `yaml:"name" json:"name"`
	Namespace   string            `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty" json:"annotations,omitempty"`
}

type TemplatePath string

func (t TemplatePath) Parse() Template {
	var template Template

	switch true {
	// s3 handler
	case strings.HasPrefix(t.String(), "s3://"):
		template.Type = "s3"
		template.Path = t.String()
	// http(s) handler
	case strings.HasPrefix(t.String(), "http://"):
		template.Type = "http"
		template.Path = t.String()
	// default is the file handler
	default:
		template.Type = "file"
		template.Path = strings.TrimPrefix(t.String(), "file://")
	}

	return template
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (m *MannyConfig) CheckForEmptyResource() (bool,error) {
	regex,err := regexp.Compile(EmptyResourceRegExp)
	if err != nil {
		return true,err
	}
	// Don't generate manifests for empty resources
	indexOfResource := strings.Index(string(m.CloudFormation), "Resource")
	if indexOfResource == -1 {
		return false,nil
	}
	resourceSubString := m.CloudFormation[indexOfResource:]

	return regex.Match(resourceSubString),nil
}

func (m *MannyConfig) runResolvers(stackList []map[string]string, fPath string) error {
	m.OutputParameters = map[string]string{}

	for key, value := range m.InputParameters {
		var actualValue string
		m.OutputParameters[key] = value.Value
		switch value.Tag {
		case "!environment_variable":
			actualValue = os.Getenv(value.Value)
			m.OutputParameters[key] = actualValue
		case "!file_contents":
			target := filepath.Join(filepath.Dir(fPath), value.Value)
			data, err := ioutil.ReadFile(target)
			if err != nil {
				return err
			}
			m.OutputParameters[key] = string(data)
		case "!stack_output":
			sOutput := strings.Split(value.Value, "::")
			sOutput[0] = filepath.Base(sOutput[0])
			for _, v := range stackList {
				if v[sOutput[0]] != "" {
					m.OutputParameters[key] = value.Tag + " " + v[sOutput[0]] + "::" + sOutput[1]
					break
				} else {
					m.OutputParameters[key] = value.Tag + " " + value.Value
				}
			}
		}
	}

	return nil
}

func (c Configurator) mannyConfigValidate(path string) bool {
	inputConfigData, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}

	_, err = config.ConfigValidate(inputConfigData, config.Schemaconfig)
	if err != nil {
		c.logger.Info("mannyConfig validation failed", zap.Error(err))
		return false
	}

	c.logger.Debug("mannyConfig validation successful", zap.String("mannyConfig", path))
	return true
}

func (t TemplatePath) IsEmpty() bool {
	return t == ""
}

func (t TemplatePath) String() string {
	return string(t)
}

type Template struct {
	Type    string `yaml:"type"`
	Path    string `yaml:"path"`
	Version string `yaml:"version"`
	Name    string `yaml:"name"`
}

// MannyConfig is a manny config. This is only read in yaml but can the Custom Resource that it gets written to can
// output in YAML or JSON.
type MannyConfig struct {
	// FoundAt is the directory where Manny found the config
	FoundAt string
	// Template is the CloudFormation template provider config.
	// Can be one of "file", "s3"
	Template Template `yaml:"template"`
	// TemplatePath provides the same functionality as Template but in one parseable string
	TemplatePath TemplatePath `yaml:"template_path"`
	// SyncWave
	SyncWave int `yaml:"syncwave,omitempty"`
	// InputParameters map to CloudFormation parameter values
	InputParameters map[string]yaml.Node `yaml:"parameters,omitempty"`
	// OutputParameters are derived from input parameters when using resolvers. These values are the ones that show up
	// in the custom resource.
	OutputParameters map[string]string
	// Tags are CloudFormation tags to apply
	Tags map[string]string `yaml:"tags"`
	// StackName is the stack name to be used during execution
	StackName string `yaml:"stackname"`
	// RoleArn is the role to execute the CloudFormation with
	RoleArn string `yaml:"rolearn"`
	// ServiceRole that Cloudformation will assume. This is to have more security.
	ServiceRoleARN string `yaml:"servicerolearn"`
	// Expiry duration of the STS credentials. Defaults to 15 minutes if not set.
	Duration time.Duration `yaml:"duration"`
	// Optional ExternalID to pass along, defaults to nil if not set.
	ExternalID string `yaml:"externalID,omitempty"`
	// Optional AccountID.
	AccountID string `yaml:"acctnum"`
	// Optional Environment.
	Environment string `yaml:"env,omitempty"`
	// If ExpiryWindow is 0 or less it will be ignored.
	ExpiryWindow time.Duration `yaml:"expirywindow,omitempty"`
	// Timeout support
	Timeout int `yaml:"timeout"`
	//Region support
	Region string `yaml:"region"`
	// Base refers to additional configuration that should be loaded
	Base           string `yaml:"base"`
	CloudFormation []byte
}

// Configurator builds final configuration from many configurations
type Configurator struct {
	// Bases is the list of configs to be merged
	Bases []MannyConfig
	// Global is the config with all bases consumed
	Global MannyConfig
	// Stacks are the stack specific configs
	Stacks []MannyConfig
	// Origin is the original path that contains a config.yaml
	Origin string
	// GitURL is the remote URL used by Git
	GitURL string
	// References is a list of bases used
	References []string
	// StackPrefix is the generated prefix of the stack
	StackPrefix string
	// StackTable is a reference table for stacks
	StackTable map[string]bool

	logger   *zap.Logger
	s3Client s3iface.S3API
}

type Config struct {
	Path     string
	GitURL   string
	Logger   *zap.Logger
	S3Client s3iface.S3API
}

// New creates a new configurator
func New(config Config) Configurator {
	c := Configurator{
		logger: config.Logger,
		GitURL: config.GitURL,
	}

	// Set the origin to an absolute path
	path, _ := filepath.Abs(config.Path)
	c.Origin = path + "/"

	switch config.S3Client.(type) {
	case *mocks.S3API:
		c.s3Client = config.S3Client
	default:
		// use user credentials unless otherwise specified
		config.Logger.Debug("S3Client not set, setting default")

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			Config: aws.Config{Region: aws.String("us-west-2")},
			SharedConfigState: session.SharedConfigEnable,
		}))

		c.s3Client = s3.New(sess)
	}

	return c
}

func (c *Configurator) CreateDeployments() (CloudResources, error) {
	// build the config
	err := c.loadBases()
	if err != nil {
		return nil, err
	}

	// render a manifest for deployment
	return c.loadStacks()
}

// loadBases resolves base configs into the Global config
func (c *Configurator) loadBases() error {
	// Abs() does not include a trailing slash
	config := c.Origin + MannyConfigName

	c.logger.Debug("Finding config", zap.String("path", config))

	// check to see if the config file exists
	_, err := os.Stat(config)
	if !os.IsNotExist(err) {
		// Unmarshal the given config to determine if there are bases and recurses them
		if err := c.unmarshal(config); err != nil {
			return err
		}
	}

	// Merge all the configs into one, starting with the bases
	c.Global, err = c.mergeBases(c.Bases)
	if err != nil {
		return err
	}

	c.determineStackPrefix()

	c.logger.Debug("Merged base configs", zap.Int("BaseConfigs", len(c.Bases)),
		zap.Any("Global Config", c.Global))

	return nil
}

func (c *Configurator) determineStackPrefix() {
	if counter == 0 {
		baseDir = strings.Split(filepath.Clean(c.Origin), "/")
	}
	counter++
	dir := strings.Split(filepath.Clean(c.Origin), "/")
	a := len(dir) - len(baseDir)
	if a >= 0{
		c.StackPrefix = strings.Join(dir[len(dir)-(a):], "-")
	}
	c.logger.Debug("Stack prefix determined", zap.String("StackPrefix", c.StackPrefix))
}

func (c Configurator) mergeBases(bases []MannyConfig) (MannyConfig, error) {
	var global MannyConfig

	for _, config := range bases {
		if err := mergo.Merge(&global, config); err != nil {
			return MannyConfig{}, err
		}
	}

	return global, nil
}

func (c *Configurator) unmarshal(parentPath string) error {
	c.logger.Debug("Reading file", zap.String("path", parentPath))

	data, err := ioutil.ReadFile(parentPath)
	if err != nil {
		return err
	}

	var config MannyConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	config.FoundAt = filepath.Clean(filepath.Dir(parentPath))
	c.logger.Debug("Storing base location", zap.String("path", config.FoundAt))

	// Recurse bases
	if config.Base != "" {
		c.logger.Debug("Base detected", zap.String("path", config.Base))

		// determine absolute path
		target := filepath.Join(filepath.Dir(parentPath), config.Base)

		c.logger.Debug("Determining target", zap.String("ConfigPath", config.Base),
			zap.String("TargetPath", target))

		// @ToDo: Convert to map/lookup table
		for _, basePath := range c.References {
			if target == basePath {
				return errors.New("circular dependency found in " + target)
			}
		}

		if len(c.References)+1 > 10 {
			return errors.New("more than 10 referenced bases")
		}

		// track references
		c.References = append(c.References, target)

		if err := c.unmarshal(target); err != nil {
			return err
		}
	}

	c.Bases = append(c.Bases, config)

	return nil
}

// loadStacks loads stacks from the local directory and creates CloudResourceDeployments from them
func (c *Configurator) loadStacks() (CloudResources, error) {
	var resources CloudResources

	c.logger.Debug("Looking for stack configs", zap.String("path", c.Origin))

	// find stack files in the origin directory
	files, err := ioutil.ReadDir(c.Origin)
	if err != nil {
		return nil, err
	}
	// process VPC stacks first and then dir
	var files_inorder, d []os.FileInfo
	for i, f := range files {
		if f.IsDir() {
			d = append(d, files[i])
		}else {
			files_inorder = append(files_inorder, f)
		}
	}
	files_inorder = append(files_inorder, d...)

	// load the stack configs
	for _, f := range files_inorder {
		// Resolve relative directories
		target := filepath.Join(filepath.Dir(c.Origin+"/"), f.Name())

		if f.Name() == MannyConfigName {
			continue
		}

		// Recurse sub directories
		if f.IsDir() {
			c.logger.Debug("Walking directory for config", zap.String("path", target))

			config := New(Config{
				Path:   target + "/",
				Logger: c.logger,
				GitURL: c.GitURL,
			})

			if err := config.loadBases(); err != nil {
				c.logger.Info("Unable to load base from higher level directory", zap.Error(err))
				continue
			}

			deployments, err := config.loadStacks()
			if err != nil {
				c.logger.Info("Unable to load stacks from higher level directory", zap.Error(err))
				continue
			}

			resources = append(resources, deployments...)

			continue
		}

		extension := filepath.Ext(target)
		if contains(AcceptedExtensions, extension) && c.mannyConfigValidate(target) {
			// Handle stack generation
			c.logger.Debug("Reading stack config", zap.String("path", target))

			data, err := ioutil.ReadFile(target)
			if err != nil {
				return nil, err
			}

			var config MannyConfig
			// unmarshal the stack config
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				return nil, err
			}

			// If no stack name is found, generate one
			m := make(map[string]string)
			if config.StackName == "" {
				config.StackName = strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				if c.StackPrefix != "" {
					config.StackName = c.StackPrefix + "-" + config.StackName
				}
				m[f.Name()] = config.StackName
				s.StackList = append(s.StackList, m)
			} else {
				m[f.Name()] = config.StackName
				s.StackList = append(s.StackList, m)
			}

			// determine whether the template block or template_path is used
			if !config.TemplatePath.IsEmpty() {
				config.Template = config.TemplatePath.Parse()
			}

			switch config.Template.Type {
			case "file":
				// Resolve relative directories
				target := filepath.Join(filepath.Dir(target), config.Template.Path)

				// Unmarshal the CloudFormation
				config.CloudFormation, err = ioutil.ReadFile(target)
				if err != nil {
					return nil, err
				}
			case "http":
				resp, err := http.Get(config.Template.Path)
				if err != nil {
					return nil, err
				}
				defer resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					config.CloudFormation, err = ioutil.ReadAll(resp.Body)
					if err != nil {
						return nil, err
					}
				} else {
					c.logger.Debug("http handler template download: http_status_code", zap.String("http_status", resp.Status))
				}
			case "s3":
				//Parsing s3 bucket and key
				u, _ := url.Parse(config.Template.Path)
				if u.Host == "" || u.Path == "/" || u.Path == "" {
					return nil, errors.New("s3 Bucket or Key not found: " + config.Template.Path)
				}

				// Support s3 http url format
				if strings.Contains(u.Host, ".s3.amazonaws.com") {
					u.Host = strings.ReplaceAll(u.Host, ".s3.amazonaws.com", "")
				}

				// tmpfile
				fPath, err := ioutil.TempFile("", "s3-"+path.Base(config.Template.Path))
				if err != nil {
					return nil, err
				}
				defer os.Remove(fPath.Name()) // clean up

				//Download cfn template from s3
				downloader := s3manager.NewDownloaderWithClient(c.s3Client)
				_, err = downloader.Download(fPath,
					&s3.GetObjectInput{
						Bucket: aws.String(u.Host),
						Key:    aws.String(u.Path),
					})
				if err != nil {
					return nil, err
				}

				// Unmarshal the CloudFormation
				config.CloudFormation, err = ioutil.ReadFile(fPath.Name())
				if err != nil {
					return nil, err
				}
			}
			// Don't generate manifests for empty resources
			resourceEmpty,err := config.CheckForEmptyResource()
			if err != nil {
				c.logger.Error("Error while check for resource count")
			}
			if !resourceEmpty {
				c.Stacks = append(c.Stacks, config)
			}

		} else {
			c.logger.Debug("Skipping file", zap.String("path", target))
		}
	}

	r, err := c.generateCR(c.Stacks)
	if err != nil {
		return nil, err
	}

	resources = append(resources, r...)

	return resources, nil
}

func (c Configurator) generateCR(stacks []MannyConfig) (CloudResources, error) {
	var manifests CloudResources

	c.logger.With(zap.String("GitRemote", c.GitURL)).Debug("Writing git remote")

	for _, stack := range stacks {
		// Run resolvers
		if err := stack.runResolvers(s.StackList, c.Origin); err != nil {
			return nil, err
		}

		// Merge the Global config and the Stack config
		if err := mergo.Merge(&stack, c.Global); err != nil {
			return nil, err
		}
		if stack.SyncWave != 0 {
			manifests = append(manifests, &CloudResourceDeployment{
				Kind:       "CloudResourceDeployment",
				APIVersion: "cloudresource.keikoproj.io/v1alpha1",
				Metadata: Metadata{
					Name: stack.StackName,
					Annotations: map[string]string{
						"source": c.GitURL,
						"argocd.argoproj.io/sync-wave": strconv.Itoa(stack.SyncWave),
					},
				},
				Spec: gitopsv1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &gitopsv1alpha1.StackSpec{
						Parameters: stack.OutputParameters,
						Tags:       stack.Tags,
						Template:   fmt.Sprintf("%s", stack.CloudFormation),
						Stackname:  stack.StackName,
						CARole: gitopsv1alpha1.AssumeRoleProvider{
							RoleARN:         stack.RoleArn,
							RoleSessionName: "gitops-deployment",
							ServiceRoleARN:	stack.ServiceRoleARN,
							ExternalID:	stack.ExternalID,
							AccountID:	stack.AccountID,
							Environment: stack.Environment,
							Duration:	stack.Duration,
							ExpiryWindow: stack.ExpiryWindow,
						},
						Timeout: stack.Timeout,
						Region: stack.Region,
					},
				},
			})
		} else {
			manifests = append(manifests, &CloudResourceDeployment{
				Kind:       "CloudResourceDeployment",
				APIVersion: "cloudresource.keikoproj.io/v1alpha1",
				Metadata: Metadata{
					Name: stack.StackName,
					Annotations: map[string]string{
						"source": c.GitURL,
					},
				},
				Spec: gitopsv1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &gitopsv1alpha1.StackSpec{
						Parameters: stack.OutputParameters,
						Tags:       stack.Tags,
						Template:   fmt.Sprintf("%s", stack.CloudFormation),
						Stackname:  stack.StackName,
						CARole: gitopsv1alpha1.AssumeRoleProvider{
							RoleARN:         stack.RoleArn,
							RoleSessionName: "gitops-deployment",
							ServiceRoleARN:	stack.ServiceRoleARN,
							ExternalID:	stack.ExternalID,
							AccountID:	stack.AccountID,
							Environment: stack.Environment,
							Duration:	stack.Duration,
							ExpiryWindow: stack.ExpiryWindow,
						},
						Timeout: stack.Timeout,
						Region: stack.Region,
					},
				},
			})
		}
	}

	return manifests, nil
}