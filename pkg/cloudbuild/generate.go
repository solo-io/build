package cloudbuild

import (
	"fmt"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/build/pkg/api/v1"
)

var (
	InvalidCiConfigError = errors.Errorf("CI config must be a Gcloud config to produce a cloudbuild")
	gopathVolume         = &Volume{
		Name: "gopath",
		Path: "/go/pkg",
	}
)

func GenerateCloudbuild(spec *v1.BuildSpec) (*Cloudbuild, error) {
	gcloud, err := getGcloudConfig(spec)
	if err != nil {
		return nil, err
	}
	return &Cloudbuild{
		Secrets: getSecrets(gcloud),
		Timeout: gcloud.GetTimeout(),
		Steps:   getSteps(gcloud, spec),
	}, nil
}

func getGcloudConfig(spec *v1.BuildSpec) (*v1.GcloudConfig, error) {
	gcloud, ok := spec.Config.CiConfig.(*v1.BuildConfig_Gcloud)
	if !ok {
		return nil, InvalidCiConfigError
	}
	return gcloud.Gcloud, nil
}

func getSteps(gcloud *v1.GcloudConfig, spec *v1.BuildSpec) []*Step {
	var steps []*Step
	mountGoCache := getMountGoCacheStep(gcloud)
	if mountGoCache != nil {
		steps = append(steps, mountGoCache)
	}
	goBuildStep := getGoBuildStep(spec)
	if goBuildStep != nil {
		steps = append(steps, goBuildStep)
	}
	steps = append(steps, getTestRunSteps(spec)...)
	return steps
}

func getTestRunSteps(spec *v1.BuildSpec) []*Step {
	var steps []*Step
	for _, testRun := range spec.TestRuns {
		if _, ok := testRun.TestRun.(*v1.TestRun_Ginkgo); ok {
			steps = append(steps, getGinkgoStep(testRun, spec))
		}
	}
	return steps
}

func getGinkgoStep(testRun *v1.TestRun, spec *v1.BuildSpec) *Step {
	return &Step{
		Name:      getGinkgoImage(spec),
		Volumes:   []*Volume{gopathVolume},
		SecretEnv: testRun.SecretEnv,
		Args:      testRun.Args,
	}
}

func getGinkgoImage(spec *v1.BuildSpec) string {
	switch spec.Config.GoConfig.Type {
	case v1.GoProjectType_GO_MOD:
		return "gcr.io/$PROJECT_ID/go-mod-ginkgo:0.1.5"
	default:
		return "gcr.io/$PROJECT_ID/e2e-ginkgo:0.1.5"
	}
}

func getMountGoCacheStep(gcloud *v1.GcloudConfig) *Step {
	if gcloud.GoCache == "" {
		return nil
	}
	return &Step{
		Name:       "gcr.io/cloud-builders/gsutil",
		Entrypoint: "bash",
		Args: []string{
			"-c",
			fmt.Sprintf("mkdir -p /go/pkg && cd /go/pkg && gsutil cat gs://%s | tar -xzf -", gcloud.GoCache),
		},
		Id:      "untar-go-cache",
		Volumes: []*Volume{gopathVolume},
	}
}

func getGoBuildStep(spec *v1.BuildSpec) *Step {
	if spec.Config.GoConfig == nil {
		return nil
	}
	if !spec.Config.GoConfig.BuildAll {
		return nil
	}
	return &Step{
		Name:    "golang:1.12",
		Volumes: []*Volume{gopathVolume},
		Id:      "go-build",
		Args:    []string{"go", "build", "./..."},
	}
}

func getSecrets(gcloud *v1.GcloudConfig) []*Secret {
	secret := getSecret(gcloud)
	if secret == nil {
		return nil
	}
	return []*Secret{secret}
}

func getSecret(gcloud *v1.GcloudConfig) *Secret {
	if gcloud.EncryptedSecrets == nil || len(gcloud.EncryptedSecrets) == 0 {
		return nil
	}
	return &Secret{
		KmsKeyName: getKmsKeyName(gcloud),
		SecretEnv:  gcloud.EncryptedSecrets,
	}
}

func getKmsKeyName(gcloud *v1.GcloudConfig) string {
	return fmt.Sprintf("projects/%s/locations/global/keyRings/%s/cryptoKeys/%s",
		gcloud.ProjectId, gcloud.DecryptKeyring, gcloud.DecryptKey)
}
