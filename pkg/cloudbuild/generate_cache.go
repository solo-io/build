package cloudbuild

import (
	"fmt"
	"github.com/solo-io/build/pkg/api/v1"
)

func GenerateCloudbuildCache(spec *v1.BuildSpec) (*Cloudbuild, error) {
	gcloud, err := getGcloudConfig(spec)
	if err != nil {
		return nil, err
	}
	return &Cloudbuild{
		Steps: getCacheSteps(gcloud),
	}, nil
}

func getCacheSteps(gcloud *v1.GcloudConfig) []*Step {
	var steps []*Step
	mountGoCache := getMountGoCacheStep(gcloud)
	if mountGoCache != nil {
		steps = append(steps, mountGoCache)
	}
	steps = append(steps, getGoModStep("download"))
	steps = append(steps, getGoModStep("tidy"))
	steps = append(steps, getTarStep(gcloud.GoCache))
	steps = append(steps, getUploadStep(gcloud.GoCache))
	return steps
}


func getGoModStep(goModAction string) *Step {
	return &Step{
		Name: "golang:1.12",
		Args: []string {"go", "mod", goModAction},
		Volumes: []*Volume {gopathVolume},
		Id: goModAction,
	}
}

func getTarStep(goCache *v1.GoCache) *Step {
	return &Step{
		Name: "golang:1.12",
		Args: []string {
			"-c",
			fmt.Sprintf("cd /go/pkg && tar -zvcf %s mod", goCache.Name),
		},
		Volumes: []*Volume {gopathVolume},
		Id: "tar-cache",
		Entrypoint: "bash",
	}
}

func getUploadStep(goCache *v1.GoCache) *Step {
	return &Step{
		Name: "gcr.io/cloud-builders/gsutil",
		Args: []string{
			"cp",
			fmt.Sprintf("/go/pkg/%s", goCache.Name),
			fmt.Sprintf("gs://%s/%s", goCache.Bucket, goCache.Name),
		},
		Volumes: []*Volume{gopathVolume},
		Id: "upload-cache",
	}
}

