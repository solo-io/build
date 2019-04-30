package ingest

import (
	"fmt"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/envutils"
)

func ValidateOperatingParameters(args []string, cv *v1.ComputedBuildVars) error {
	if len(args) != 5 {
		return errors.Errorf("expected 5 arguments, received %v", len(args))
	}
	errorReport := ""
	expectEqual("release", envutils.StringForBoolToEnv(cv.Release), args[0], &errorReport)
	expectEqual("version", cv.Version, args[1], &errorReport)
	expectEqual("container prefix", cv.ContainerPrefix, args[2], &errorReport)
	expectEqual("image tag", cv.ImageTag, args[3], &errorReport)
	expectEqual("helm repo url", cv.HelmRepository, args[4], &errorReport)
	if errorReport != "" {
		return errors.Errorf(errorReport)
	}
	return nil
}

func expectEqual(name, want, got string, report *string) {
	if want == got {
		return
	}
	*report = fmt.Sprintf("%v%v wants: %v, got: %v\n", *report, name, want, got)
}
