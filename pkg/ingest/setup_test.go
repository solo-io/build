package ingest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/build/pkg/api/v1"
)

var _ = Describe("ingest config", func() {

	var ()

	var _ = BeforeEach(func() {
	})
	var _ = AfterEach(func() {
	})

	Context("unit test isRelease", func() {
		It("should report release correctly", func() {
			ev1 := &v1.BuildEnvVars{
				TagVersion: "",
			}
			Expect(isRelease(ev1)).To(Equal(false))
			ev2 := &v1.BuildEnvVars{
				TagVersion: "v1.2.3",
			}
			Expect(isRelease(ev2)).To(Equal(true))
		})
	})
	Context("unit test setImageTag", func() {
		It("should report tag correctly for valid configs", func() {
			buildId := "12345"
			nonRelease := &v1.BuildEnvVars{
				TagVersion: "",
				BuildId:    buildId,
			}
			const rootTag = "v1.2.3"
			tag := rootTag
			Expect(setImageTag(&tag, nonRelease)).NotTo(HaveOccurred())
			Expect(tag).To(Equal(buildId))
			taggedVersion := rootTag
			version := "1.2.3"
			release := &v1.BuildEnvVars{
				TagVersion: taggedVersion,
				BuildId:    buildId,
			}
			Expect(setImageTag(&tag, release)).NotTo(HaveOccurred())
			Expect(tag).To(Equal(version))
			// we require semver
			release.TagVersion = "vabcdefg"
			Expect(setImageTag(&tag, release)).To(HaveOccurred())
			release.TagVersion = "vv"
			Expect(setImageTag(&tag, release)).To(HaveOccurred())
			release.TagVersion = "v with some space"
			Expect(setImageTag(&tag, release)).To(HaveOccurred())
		})
		It("should error for invalid configs", func() {
			buildId := "12345"
			nonReleaseNoBuildId := &v1.BuildEnvVars{
				TagVersion: "",
				BuildId:    "",
			}
			tag := ""
			Expect(setImageTag(&tag, nonReleaseNoBuildId)).To(HaveOccurred())
			taggedVersion := "1.2.3"
			releaseBadTaggedVersion := &v1.BuildEnvVars{
				TagVersion: taggedVersion,
				BuildId:    buildId,
			}
			Expect(setImageTag(&tag, releaseBadTaggedVersion)).To(HaveOccurred())
			releaseBadTaggedVersion.TagVersion = "v"
			Expect(setImageTag(&tag, releaseBadTaggedVersion)).To(HaveOccurred())
			releaseBadTaggedVersion.TagVersion = "a1.2.3"
			Expect(setImageTag(&tag, releaseBadTaggedVersion)).To(HaveOccurred())
		})
	})
	Context("unit test setContainerPrefix", func() {
		It("should report container spec correctly for valid single container configs", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:     "gcr.io",
							ProjectName: "aproject",
						},
					},
				},
				TestContainerRegistry: nil,
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject/"))
			release = true
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject/"))
		})
		It("should error for invalid single container configs", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: nil,
				TestContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:     "gcr.io",
							ProjectName: "aproject",
						},
					},
				},
			}
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			release = true
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
		})
		It("should report container spec correctly for valid dual container configs", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:     "gcr.io",
							ProjectName: "aproject",
						},
					},
				},
				TestContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Quay{
						Quay: &v1.QuayRegistry{
							BaseUrl:      "quay.io",
							Organization: "an-org",
						},
					},
				},
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("quay.io/an-org/"))
			release = true
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject/"))
		})
		It("should validate gcr container specs correctly", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:     "gcr.io",
							ProjectName: "aproject",
						},
					},
				},
				TestContainerRegistry: nil,
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject/"))
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.BaseUrl = ""
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.BaseUrl = "gcr.io"
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.ProjectName = ""
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.ProjectName = "aproject"
			Expect(setContainerPrefix(&prefix, release, config)).NotTo(HaveOccurred())
		})
		It("should validate quay container specs correctly", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Quay{
						Quay: &v1.QuayRegistry{
							BaseUrl:      "quay.io",
							Organization: "an-org",
						},
					},
				},
				TestContainerRegistry: nil,
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("quay.io/an-org/"))
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Quay).Quay.BaseUrl = ""
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Quay).Quay.BaseUrl = "quay.io"
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Quay).Quay.Organization = ""
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Quay).Quay.Organization = "an-org"
			Expect(setContainerPrefix(&prefix, release, config)).NotTo(HaveOccurred())
		})
		It("should validate docker container specs correctly", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_DockerHub{
						DockerHub: &v1.DockerHubRegistry{},
					},
				},
				TestContainerRegistry: nil,
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal(""))
		})
	})
})
