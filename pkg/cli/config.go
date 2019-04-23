package cli

import "github.com/solo-io/build/pkg/version"

var (
	Version             = version.Version
	FileLogPathElements = []string{".solobuild", "log"}
	OutputModeEnvVar    = "SOLOBUILD_OUTPUT_MODE"
	RootErrorMessage    = "error while running cli build"
)
