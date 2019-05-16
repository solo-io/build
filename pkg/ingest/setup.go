package ingest

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/pkg/errors"
	"github.com/solo-io/go-utils/versionutils"

	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/constants"
	"github.com/solo-io/go-utils/protoutils"
)

func InitializeBuildRun(configFilename string, explicitBuildEnvVars *v1.BuildEnvVars) (v1.BuildRun, error) {
	buildConfig, err := parseSpec(configFilename)
	if err != nil {
		return v1.BuildRun{}, err
	}
	buildRunConfig, err := getBuildRunConfig(buildConfig, explicitBuildEnvVars)
	if err != nil {
		return v1.BuildRun{}, err
	}
	return v1.BuildRun{
		Spec:   &v1.BuildSpec{
			Config: buildConfig,
		},
		Config: &buildRunConfig,
	}, nil
}

// uses a config filename from env or default, in that order
func parseSpec(filename string) (*v1.BuildConfig, error) {
	if filename == "" {
		contextutils.LoggerFrom(context.TODO()).Debugw("project filename not provided, checking env")
		envFile := os.Getenv(constants.EnvVarConfigFileName)
		if envFile != "" {
			filename = envFile
		} else {
			contextutils.LoggerFrom(context.TODO()).Debugw("project filename env var not found, using default filename")
			filename = constants.DefaultConfigFileName
		}
	}
	b, err := ioutil.ReadFile(filename)
	spec := &v1.BuildConfig{}
	if err != nil {
		return spec, errors.Wrapf(err, "could not parse build spec")
	}
	if err := protoutils.UnmarshalYaml(b, spec); err != nil {
		return spec, errors.Wrapf(err, "could not unmarshal build spec")
	}
	return spec, nil
}

func getBuildRunConfig(config *v1.BuildConfig, explicitBuildEnvVars *v1.BuildEnvVars) (v1.BuildRunConfig, error) {
	ev := resolveBuildEnvVars(explicitBuildEnvVars)
	cv := &v1.ComputedBuildVars{}
	cv.Release = isRelease(ev)
	var err error
	cv.Version, err = getVersion(cv.Release, ev.TaggedVersion, ev.BuildId)
	if err != nil {
		return v1.BuildRunConfig{}, errors.Wrapf(err, "could not get version")
	}
	if err := setImageTag(&cv.ImageTag, ev); err != nil {
		return v1.BuildRunConfig{}, errors.Wrapf(err, "could not set image tag")
	}
	if err := setContainerPrefix(&cv.ContainerPrefix, cv.Release, config); err != nil {
		return v1.BuildRunConfig{}, errors.Wrapf(err, "could not set container prefix")
	}
	if err := setHelmRepository(&cv.HelmRepository, cv.Release, config); err != nil {
		return v1.BuildRunConfig{}, errors.Wrapf(err, "could not set helm chart repository")
	}
	return v1.BuildRunConfig{
		BuildEnvVars:      ev,
		ComputedBuildVars: cv,
	}, nil
}

func resolveBuildEnvVars(explicitBuildEnvVars *v1.BuildEnvVars) *v1.BuildEnvVars {
	// copy values, if any
	ev := &v1.BuildEnvVars{
		BuildId:       explicitBuildEnvVars.BuildId,
		TaggedVersion: explicitBuildEnvVars.TaggedVersion,
	}
	if explicitBuildEnvVars.BuildId == "" {
		ev.BuildId = os.Getenv(constants.EnvBuildId)
	}
	if explicitBuildEnvVars.TaggedVersion == "" {
		ev.TaggedVersion = os.Getenv(constants.EnvTagVersion)
	}
	return ev
}

func getVersion(release bool, taggedVersion, buildId string) (string, error) {
	version := buildId
	if release {
		var err error
		version, err = versionutils.GetVersionFromTag(taggedVersion)
		if err != nil {
			return "", err
		}
	}
	return version, nil
}

func isRelease(ev *v1.BuildEnvVars) bool {
	if ev.TaggedVersion == "" {
		return false
	}
	return true
}

func setImageTag(tag *string, ev *v1.BuildEnvVars) error {
	*tag = ev.BuildId
	if isRelease(ev) {
		rv, err := versionutils.ParseVersion(ev.TaggedVersion)
		if err != nil {
			return err
		}
		rvString := rv.String()
		*tag = rvString[1:]
		return nil
	}
	if *tag == "" {
		return fmt.Errorf("must specify an image tag, none found for build env vars: %v", ev)
	}
	return nil
}

func setContainerPrefix(prefix *string, isRelease bool, config *v1.BuildConfig) error {
	if config.ReleaseContainerRegistry == nil {
		return fmt.Errorf("must provide a release_container_registry")
	}
	targetRegistry := config.ReleaseContainerRegistry
	if !isRelease && config.TestContainerRegistry != nil {
		targetRegistry = config.TestContainerRegistry
	}
	if err := targetRegistry.SetPrefixFromContainerRegistry(prefix); err != nil {
		return errors.Wrapf(err, "could not set container prefix")
	}
	return nil
}

func setHelmRepository(repository *string, isRelease bool, config *v1.BuildConfig) error {
	if config.ReleaseHelmRepository == nil {
		return fmt.Errorf("must provide a release_helm_repository")
	}
	if config.TestHelmRepository == nil {
		return fmt.Errorf("must provide a test_helm_repository")
	}
	targetRepo := config.TestHelmRepository
	if isRelease {
		targetRepo = config.ReleaseHelmRepository
	}
	// Currently the only supported repo type is GCS
	gcs := targetRepo.GetGoogleCloudStorage()
	if gcs == nil {
		return fmt.Errorf("unexpected helm repository type %T. We currently only support Google Cloud Storage", targetRepo)
	}
	*repository = gcs.BucketUrl
	return nil
}
