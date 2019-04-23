package ingest

import (
	"fmt"
	"io/ioutil"
	"os"

	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/constants"
	"github.com/solo-io/go-utils/protoutils"
)

func InitializeBuildRun() v1.BuildRun {
	buildSpec := parseSpec()
	buildRunConfig := getBuildRunConfigFromEnv(buildSpec)
	return v1.BuildRun{
		Spec:   buildSpec,
		Config: &buildRunConfig,
	}
}

// uses a config filename from env or default, in that order
func parseSpec() *v1.BuildSpec {
	filename := ""
	spec := &v1.BuildSpec{}
	envFile := os.Getenv(constants.EnvVarConfigFileName)
	if envFile != "" {
		filename = envFile
	} else {
		filename = constants.DefaultConfigFileName
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err := protoutils.UnmarshalYaml(b, spec); err != nil {
		panic(err)
	}
	return spec
}

func getBuildRunConfigFromEnv(spec *v1.BuildSpec) v1.BuildRunConfig {
	ev := &v1.BuildEnvVars{}
	ev.BuildId = os.Getenv(constants.EnvBuildId)
	ev.CommitSha = os.Getenv(constants.EnvCommitSha)
	ev.TagVersion = os.Getenv(constants.EnvTagVersion)
	cv := &v1.ComputedBuildVars{}
	cv.Release = isRelease(ev)
	cv.ImageTag = getImageTag(ev)
	cv.ContainerPrefix = getContainerPrefix(cv.Release, spec.Config)
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
	tag := ev.BuildId
	if isRelease(ev) {
		tag = imageTagFromTaggedVersion(ev.TagVersion)
	}
	if tag == "" {
		panic(fmt.Sprintf("must specify an image tag, none found for build env vars: %v", ev))
	}
	return tag
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

func getContainerPrefix(isRelease bool, config *v1.BuildConfig) string {
	targetRegistry := config.ReleaseContainerRegistry
	if !isRelease && config.TestContainerRegistry != nil {
		targetRegistry = config.TestContainerRegistry
	}
	prefix := ""
	if err := targetRegistry.GetPrefixFromContainerRegistry(&prefix); err != nil {
		panic(err)
	}
	return prefix
}
