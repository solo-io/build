package ingest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/build/pkg/api/v1"
	"github.com/solo-io/build/pkg/constants"
	"github.com/solo-io/build/pkg/envutils"
)

var _ = Describe("ingest config", func() {

	Context("ValidateOperatingParameters", func() {
		const (
			releaseTrue        = constants.PrintEnvTrue
			releaseFalse       = constants.PrintEnvFalse
			version            = "1.2.3"
			containerRepo      = "gcr.io/abcd"
			releaseImageTag    = "1.2.3"
			nonReleaseImageTag = "4444"
		)
		It("should correctly validate", func() {
			By("with valid args")
			args := []string{releaseTrue, version, containerRepo, releaseImageTag}
			releaseBool, err := envutils.BoolFromEnvString(releaseTrue)
			Expect(err).NotTo(HaveOccurred())
			cv := &v1.ComputedBuildVars{
				Release:         releaseBool,
				ImageTag:        releaseImageTag,
				ContainerPrefix: containerRepo,
				Version:         version,
			}
			Expect(ValidateOperatingParameters(args, cv)).NotTo(HaveOccurred())
			args = []string{releaseFalse, "", containerRepo, nonReleaseImageTag}
			releaseBool, err = envutils.BoolFromEnvString(releaseFalse)
			Expect(err).NotTo(HaveOccurred())
			cv = &v1.ComputedBuildVars{
				Release:         releaseBool,
				ImageTag:        nonReleaseImageTag,
				ContainerPrefix: containerRepo,
				Version:         "",
			}
			Expect(ValidateOperatingParameters(args, cv)).NotTo(HaveOccurred())
			By("with invalid args")
			args = []string{"badString", version, containerRepo, releaseImageTag}
			err = ValidateOperatingParameters(args, cv)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(`release wants: FALSE, got: badString
version wants: , got: 1.2.3
image tag wants: 4444, got: 1.2.3
`))
			args = []string{"badString", "badString2", containerRepo, releaseImageTag}
			err = ValidateOperatingParameters(args, cv)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(`release wants: FALSE, got: badString
version wants: , got: badString2
image tag wants: 4444, got: 1.2.3
`))
			args = []string{}
			err = ValidateOperatingParameters(args, cv)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(`expected 4 arguments, received 0`))
		})
	})
})
