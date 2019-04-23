package ingest

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/constants"
	"github.com/solo-io/go-utils/protoutils"
)

func InitializeBuildRun() (v1.BuildRun, error) {
	buildSpec, err := parseSpec()
	if err != nil {
		return v1.BuildRun{}, err
	}
	buildRunConfig, err := getBuildRunConfigFromEnv(buildSpec)
	if err != nil {
		return v1.BuildRun{}, err
	}
	return v1.BuildRun{
		Spec:   buildSpec,
		Config: &buildRunConfig,
	}, nil
}

// uses a config filename from env or default, in that order
func parseSpec() (*v1.BuildSpec, error) {
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
		return spec, errors.Wrapf(err, "could not parse build spec")
	}
	if err := protoutils.UnmarshalYaml(b, spec); err != nil {
		return spec, errors.Wrapf(err, "could not unmarshal build spec")
	}
	return spec, nil
}

func getBuildRunConfigFromEnv(spec *v1.BuildSpec) (v1.BuildRunConfig, error) {
	ev := &v1.BuildEnvVars{}
	ev.BuildId = os.Getenv(constants.EnvBuildId)
	ev.CommitSha = os.Getenv(constants.EnvCommitSha)
	ev.TagVersion = os.Getenv(constants.EnvTagVersion)
	cv := &v1.ComputedBuildVars{}
	cv.Release = isRelease(ev)
	if err := setImageTag(&cv.ImageTag, ev); err != nil {
		return v1.BuildRunConfig{}, errors.Wrapf(err, "could not set image tag")
	}
	if err := setContainerPrefix(&cv.ContainerPrefix, cv.Release, spec.Config); err != nil {
		return v1.BuildRunConfig{}, errors.Wrapf(err, "could not set container prefix")
	}
	return v1.BuildRunConfig{
		BuildEnvVars:      ev,
		ComputedBuildVars: cv,
	}, nil
}

func isRelease(ev *v1.BuildEnvVars) bool {
	if ev.TagVersion == "" {
		return false
	}
	return true
}

func setImageTag(tag *string, ev *v1.BuildEnvVars) error {
	*tag = ev.BuildId
	if isRelease(ev) {
		if err := setImageTagFromTaggedVersion(tag, ev.TagVersion); err != nil {
			return err
		}
	}
	if *tag == "" {
		return fmt.Errorf("must specify an image tag, none found for build env vars: %v", ev)
	}
	return nil
}

func setImageTagFromTaggedVersion(tag *string, tv string) error {
	if len(tv) < 2 {
		return fmt.Errorf("must have at least two characters in TaggedVersion")
	}
	if tv[0] != 'v' {
		return fmt.Errorf("invalid tagged version: %v, must start with 'v'", tv)
	}
	*tag = tv[1:]
	return nil
}

func setContainerPrefix(prefix *string, isRelease bool, config *v1.BuildConfig) error {
	targetRegistry := config.ReleaseContainerRegistry
	if !isRelease && config.TestContainerRegistry != nil {
		targetRegistry = config.TestContainerRegistry
	}
	if err := targetRegistry.SetPrefixFromContainerRegistry(prefix); err != nil {
		return errors.Wrapf(err, "could not set container prefix")
	}
	return nil
}
