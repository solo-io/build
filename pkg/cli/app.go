package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/solo-io/go-utils/protoutils"

	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/constants"
	"github.com/spf13/cobra"
)

type Options struct {
	Internal Internal
	Input    Input
	BuildRun v1.BuildRun
}

type Internal struct {
	ctx context.Context
}
type Input struct {
	Debug bool
}

func App(ctx context.Context, version string) *cobra.Command {
	o := &Options{
		Internal: Internal{ctx: ctx},
		BuildRun: InitializeBuildRun(),
	}
	app := &cobra.Command{
		Use:     "build",
		Short:   "CLI for solo.io's build tool",
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if o.Input.Debug {
				config := o.BuildRun.Config.BuildEnvVars
				contextutils.LoggerFrom(o.Internal.ctx).Infow("read build run config values",
					"build_id", config.BuildId,
					"commit_sha", config.CommitSha,
					"tag_version", config.TagVersion)
			}
			return nil
		},
	}

	app.AddCommand(
		o.parseBuildEnvArgs(),
	)
	app.PersistentFlags().BoolVar(&o.Input.Debug, "debug", false, "enable verbose debug output")
	app.ParseFlags([]string{})
	return app
}

func (o *Options) parseBuildEnvArgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse-env",
		Short: "read environment variables and return corresponding build values",
	}
	cmd.AddCommand(
		o.reportRelease(),
		o.reportImageTag(),
		o.reportContainerPrefix())
	return cmd
}

func (o *Options) reportRelease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release",
		Short: "reports if a build is a release build",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(stringForBoolToEnv(o.BuildRun.Config.ComputedBuildVars.Release))
			return nil
		},
	}
	return cmd
}

func (o *Options) reportImageTag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image-tag",
		Short: "reports the image tag to use for this build",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(o.BuildRun.Config.ComputedBuildVars.ImageTag)
			return nil
		},
	}
	return cmd
}

func (o *Options) reportContainerPrefix() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "container-prefix",
		Short: "reports the container repo and org spec (ex: gcr.io/solo-projects/)",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(o.BuildRun.Config.ComputedBuildVars.ContainerPrefix)
			return nil
		},
	}
	return cmd
}

func stringForBoolToEnv(b bool) string {
	if b {
		return constants.PrintEnvTrue
	}
	return constants.PrintEnvFalse
}

func getBuildRunConfigFromEnv() v1.BuildRunConfig {
	ev := &v1.BuildEnvVars{}
	ev.BuildId = os.Getenv(constants.EnvBuildId)
	ev.CommitSha = os.Getenv(constants.EnvCommitSha)
	ev.TagVersion = os.Getenv(constants.EnvTagVersion)
	cv := &v1.ComputedBuildVars{}
	cv.Release = isRelease(ev)
	cv.ImageTag = getImageTag(ev)
	cv.ContainerPrefix = "TODO-CONTAINER-PREFIX"
	return v1.BuildRunConfig{
		BuildEnvVars:      ev,
		ComputedBuildVars: cv,
	}
}

func isRelease(ev *v1.BuildEnvVars) bool {
	if ev.TagVersion == "" {
		return false
	}
	return true
}

func getImageTag(ev *v1.BuildEnvVars) string {
	if isRelease(ev) {
		return imageTagFromTaggedVersion(ev.TagVersion)
	}
	return ev.BuildId
}

func imageTagFromTaggedVersion(tv string) string {
	if len(tv) < 2 {
		panic("must have at least two characters in TaggedVersion")
	}
	if tv[0] != 'v' {
		panic(fmt.Sprintf("invalid tagged version: %v, must start with 'v'", tv))
	}
	return tv[1:]
}

func InitializeBuildRun() v1.BuildRun {
	buildRunConfig := getBuildRunConfigFromEnv()
	buildSpec := parseSpec(constants.ConfigFileName)
	return v1.BuildRun{
		Spec:   buildSpec,
		Config: &buildRunConfig,
	}
}

func parseSpec(filename string) *v1.BuildSpec {
	spec := &v1.BuildSpec{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err := protoutils.UnmarshalYaml(b, spec); err != nil {
		panic(err)
	}
	return spec
}
