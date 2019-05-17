package cloudbuild

import (
	"fmt"
	"github.com/solo-io/build/pkg/api/v1"
)

const (
	GsutilContainer = "gcr.io/cloud-builders/gsutil"
	GolangContainer = "golang:1.12"
)

var (
	gopathVolume = &Volume{
		Name: "gopath",
		Path: "/go/pkg",
	}
)

func getGinkgoImage(goConfig *v1.GoConfig) string {
	switch goConfig.Type {
	case v1.GoProjectType_GO_MOD:
		return "gcr.io/$PROJECT_ID/go-mod-ginkgo:0.1.5"
	default:
		return "gcr.io/$PROJECT_ID/e2e-ginkgo:0.1.5"
	}
}

func getMountCacheCommand(cache *v1.GoCache) string {
	return fmt.Sprintf("mkdir -p /go/pkg && cd /go/pkg && gsutil cat gs://%s/%s | tar -xzf -",
		cache.Bucket, cache.Name)
}

func getGolangImage() string {
	return GolangContainer
}

func getKmsKeyName(gcloud *v1.GcloudConfig) string {
	return fmt.Sprintf("projects/%s/locations/global/keyRings/%s/cryptoKeys/%s",
		gcloud.ProjectId, gcloud.DecryptKeyring, gcloud.DecryptKey)
}