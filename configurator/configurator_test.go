package configurator

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/mock"
	"github.com/keikoproj/cloudresource-manager/api/v1alpha1"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/keikoproj/manny/configurator/mocks"
)

func TestExamples(t *testing.T) {
	var table = []struct {
		name   string
		path   string
		length int
		error  bool
	}{
		{"nested configuration", "./examples/nested/aws/usw2/", 6, false},
		{"static configuration", "./examples/static/aws/usw2/", 3, false},
		{"static configuration", "./examples/static/aws/use2/", 2, false},
		{"s3 configuration", "./examples/s3/aws/usw2/", 2, false},
		{"http configuration", "./examples/http/aws/usw2/", 1, false},
		{"resolvers configuration", "./examples/resolvers/", 2, false},
		{"error on circular dependencies", "./tests/circular-dependency/", 0, true},
		{"error on greater than 10 base references", "./tests/base-reference-cap/", 0, true},
		{ "don't create stack for empty resources", "./examples/complex/missing-resources/", 0, false},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			// set working directory as manny root directory
			_, filename, _, _ := runtime.Caller(0)
			wd := filepath.Join(filename, "../..")
			err := os.Chdir(wd)

			// handle relative and absolute paths
			if strings.HasPrefix(tt.path, ".") {
				tt.path = filepath.Join(wd, tt.path)
			}

			// show debug logs if a test fails
			log, err := initTestLogger()
			if err != nil {
				t.Error(err)
			}

			// mock s3 api
			mockS3 := new(mocks.S3API)

			config := New(Config{
				Path:     tt.path,
				Logger:   log,
				S3Client: mockS3,
			})

			if tt.name == "s3 configuration" {
				ctx := context.Background()
				testFileBody := `{
									"Resources" : {
										"HelloBucket" : {
											"Type" : "AWS::S3::Bucket",
											"Properties" : {
											   "AccessControl" : "PublicRead"
											}
										}
									}
								}`
				contentLength := int64(len(testFileBody))
				s3GetObjectOutput := &s3.GetObjectOutput{
					Body:          ioutil.NopCloser(strings.NewReader(testFileBody)),
					ContentLength: &contentLength,
				}
				mockS3.On("GetObjectWithContext", ctx, mock.Anything, mock.Anything).
					Return(s3GetObjectOutput, nil)
			}

			deployments, err := config.CreateDeployments()
			if err != nil && !tt.error {
				t.Error(err)
			}

			if len(deployments) != tt.length {
				t.Errorf("Incorrect length got: %d want: %d\n", len(deployments), tt.length)
			}
			invalidConfig :=config.mannyConfigValidate("foo")
			if invalidConfig{
				t.Errorf("mannyConfigValidate failed")
			}
		})
	}
}

func Test_CloudResourcesValidate(t *testing.T) {
	var table = []struct {
		name      string
		resources CloudResources
		error     bool
	}{
		{"duplicate stack name", CloudResources{
			{
				Kind:       "",
				APIVersion: "",
				Metadata:   Metadata{},
				Spec: v1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &v1alpha1.StackSpec{
						Parameters: nil,
						Tags:       nil,
						Template:   "",
						Stackname:  "test1",
						CARole:     v1alpha1.AssumeRoleProvider{},
					},
				},
			},
			{
				Kind:       "",
				APIVersion: "",
				Metadata:   Metadata{},
				Spec: v1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &v1alpha1.StackSpec{
						Parameters: nil,
						Tags:       nil,
						Template:   "",
						Stackname:  "test1",
						CARole:     v1alpha1.AssumeRoleProvider{},
					},
				},
			},
		}, true},
		{"happy path", CloudResources{
			{
				Kind:       "",
				APIVersion: "",
				Metadata:   Metadata{},
				Spec: v1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &v1alpha1.StackSpec{
						Parameters: nil,
						Tags:       nil,
						Template:   "",
						Stackname:  "test1",
						CARole:     v1alpha1.AssumeRoleProvider{},
					},
				},
			},
			{
				Kind:       "",
				APIVersion: "",
				Metadata:   Metadata{},
				Spec: v1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &v1alpha1.StackSpec{
						Parameters: nil,
						Tags:       nil,
						Template:   "",
						Stackname:  "test2",
						CARole:     v1alpha1.AssumeRoleProvider{},
					},
				},
			},
		}, true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.resources.Validate(); err != nil && !tt.error {
				t.Error(err)
			}
		})
	}
}

func Test_CloudResourcesRender(t *testing.T) {
	var table = []struct {
		name       string
		resources  CloudResources
		jsonLength int
		yamlLength int
		error      bool
	}{
		{"happy path", CloudResources{
			{
				Kind:       "",
				APIVersion: "",
				Metadata:   Metadata{},
				Spec: v1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &v1alpha1.StackSpec{
						Parameters: nil,
						Tags:       nil,
						Template:   "",
						Stackname:  "test1",
						CARole: v1alpha1.AssumeRoleProvider{
							RoleARN:         "arn:aws:iam::123456789000:role/cfn-fa-role",
							RoleSessionName: "gitops-deployment",
						},
					},
				},
			},
			{
				Kind:       "",
				APIVersion: "",
				Metadata:   Metadata{},
				Spec: v1alpha1.CloudResourceDeploymentSpec{
					Cloudformation: &v1alpha1.StackSpec{
						Parameters: nil,
						Tags:       nil,
						Template:   "",
						Stackname:  "test2",
						CARole: v1alpha1.AssumeRoleProvider{
							RoleARN:         "arn:aws:iam::123456789000:role/cfn-fa-role",
							RoleSessionName: "gitops-deployment",
						},
					},
				},
			},
		}, 446, 522, true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := tt.resources.Render("yaml")
			if err != nil && !tt.error {
				t.Error(err)
			}

			// TODO: Should validate contents of the yaml doc, not just length
			if len(bytes) != tt.yamlLength {
				t.Logf(string(bytes))
				t.Errorf("YAML byte array length is incorrect; got: %d want %d", len(bytes), tt.yamlLength)
			}

			bytes, err = tt.resources.Render("json")
			if err != nil && !tt.error {
				t.Error(err)
			}

			// TODO: Should validate contents of the json doc, not just length
			if len(bytes) != tt.jsonLength {
				t.Logf(string(bytes))
				t.Errorf("JSON byte array length is incorrect; got: %d want %d", len(bytes), tt.jsonLength)
			}
		})
	}
}

func initTestLogger() (*zap.Logger, error) {
	cfg := zap.Config{
		Encoding:         "console",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	return cfg.Build()
}
