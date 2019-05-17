package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/cloudbuild"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/protoutils"
	"go.uber.org/zap"
)

func main() {
	ctx := context.TODO()
	bytes, err := ioutil.ReadFile("examples/sample-cloudbuild-gen.yaml")
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Error reading config", zap.Error(err))
	}
	spec, err := ReadBuildSpec(bytes)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Error parsing config", zap.Error(err))
	}
	printCloudbuildYaml(ctx, spec)
	printCloudbuildCacheYaml(ctx, spec)
}

func printCloudbuildYaml(ctx context.Context, spec *v1.BuildSpec) {
	build, err := cloudbuild.GenerateCloudbuild(spec)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Error converting to cloudbuild", zap.Error(err))
	}
	cloudbuildYaml, err := GetCloudbuildString(build)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Error marshalling to yaml", zap.Error(err))
	}
	fmt.Printf("cloudbuild.yaml:\n")
	fmt.Printf("----------------\n")
	fmt.Printf(cloudbuildYaml + "\n")
}

func printCloudbuildCacheYaml(ctx context.Context, spec *v1.BuildSpec) {
	build, err := cloudbuild.GenerateCloudbuildCache(spec)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Error converting to cloudbuild", zap.Error(err))
	}
	cloudbuildYaml, err := GetCloudbuildString(build)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Error marshalling to yaml", zap.Error(err))
	}
	fmt.Printf("cloudbuild.yaml:\n")
	fmt.Printf("----------------\n")
	fmt.Printf(cloudbuildYaml + "\n")
}

func ReadBuildSpec(bytes []byte) (*v1.BuildSpec, error) {
	var spec v1.BuildRun
	if err := protoutils.UnmarshalYaml(bytes, &spec); err != nil {
		return nil, err
	}
	return spec.Spec, nil
}

func GetCloudbuildString(build *cloudbuild.Cloudbuild) (string, error) {
	bytes, err := yaml.Marshal(build)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
