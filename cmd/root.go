package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ErrMissingArg  = "missing argument"
	ConfigDebug    = "debug"
	ConfigFormat   = "format"
	OutputFormat   = "output"
	ConfigValidate = "validate"
	ConfigDryRun   = "dry-run"
)

var (
	Logger *zap.Logger

	// flag storage
	debug    bool
	format   string
	output   string
	dryRun   bool
	validate bool

	// Commands
	rootCmd = &cobra.Command{
		Use:   "manny",
		Short: "Argo CD tool to generate K8s manifests from GitOps repo",
		Long:  `Argo CD tool to generate K8s manifests from GitOps repo`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger)

	rootCmd.AddCommand(buildCmd)

	rootCmd.PersistentFlags().BoolVarP(&debug, ConfigDebug, "D", false, "sets debug mode")
	buildCmd.PersistentFlags().StringVarP(&format, ConfigFormat, "f", "yaml", "sets output format")
	buildCmd.PersistentFlags().StringVarP(&output, OutputFormat, "o", "stdout", "sets file location")
	buildCmd.PersistentFlags().BoolVarP(&validate, ConfigValidate, "", true, "validates the CloudFormation output")
	buildCmd.PersistentFlags().BoolVarP(&dryRun, ConfigDryRun, "", false, "does not output a CloudResource")
}

// initLogger reads in config file and ENV variables if set.
func initLogger() {
	cfg := zap.Config{
		Encoding:         "console",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
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

	if debug {
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	l, err := cfg.Build()
	if err != nil {
		fmt.Printf("Error setting up logger: %s", err)
	}

	Logger = l
}
