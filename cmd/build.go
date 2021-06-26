package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/keikoproj/manny/configurator"
	"github.com/keikoproj/manny/utils"
)

var (
	buildCmd = &cobra.Command{
		Use:     "build path/to/stacks",
		Short:   "Builds a manny deployment",
		Long:    "Builds a manny deployment",
		Example: "manny build usw2",
		RunE:    buildConfig,
	}
)

func buildConfig(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New(ErrMissingArg)
	}

	path := args[0]

	Logger.Debug("Git repository path", zap.String("repo-path", filepath.Join(path, git.GitDirName)))

	gitURL, err := utils.GitRepoRemote(path)
	if err != nil {
		return err
	}

	// create a new configurator with defaults
	config := configurator.New(configurator.Config{
		Path:   path,
		Logger: Logger,
		GitURL: gitURL,
	})

	deployments, err := config.CreateDeployments()
	if err != nil {
		return err
	}

	Logger.Debug("Deployments created", zap.Any("CloudResourceDeployments", len(deployments)))

	if validate {
		if err := deployments.Validate(); err != nil {
			Logger.Error("Validation failed", zap.Error(err))
			return err
		}
	}

	// early return for dry run
	if dryRun {
		return nil
	}

	// render the manifest in a given format
	bytes, err := deployments.Render(format)
	if err != nil {
		return err
	}

	// output to location
	if output != "stdout" {
		// File location validation
		ok, err := utils.ValidateAndWrite(output, bytes)
		if !ok {
			return err
		}

		// early return
		return nil
	}

	// write to stdout
	fmt.Printf("%s", bytes)

	return nil
}
