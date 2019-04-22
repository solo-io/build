package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

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
	fmt.Println(o)
	fmt.Println(o.BuildRun)
	fmt.Println(o.BuildRun.Spec)
	app := &cobra.Command{
		Use:     "build",
		Short:   "CLI for solo.io's build tool",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	app.AddCommand(
		o.parseBuildEnvArgs(),
	)
	return app
}

func (o *Options) parseBuildEnvArgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse-env",
		Short: "read environment variables and return corresponding build values",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

func getBuildRunConfigFromEnv() v1.BuildRunConfig {
	brc := v1.BuildRunConfig{}
	brc.BuildId = os.Getenv(constants.EnvBuildId)
	brc.CommitSha = os.Getenv(constants.EnvCommitSha)
	brc.TagVersion = os.Getenv(constants.EnvTagVersion)
	return brc
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
