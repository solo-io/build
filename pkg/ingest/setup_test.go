package ingest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	v1 "github.com/solo-io/build/pkg/api/v1"
)

var _ = Describe("ingest config", func() {

	const relativePathToSampleConfig = "../../examples/sample-solo-project.yaml"
	var buildId = "some-build-id"
	const nonReleaseTagValue = ""
	var explicitBuildEnvVars = &v1.BuildEnvVars{
		BuildId:       buildId,
		TaggedVersion: nonReleaseTagValue,
	}
	var sampleConfigIngested = v1.BuildRun{
		Spec: &v1.BuildSpec{
			Config: &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Quay{
						Quay: &v1.QuayRegistry{
							BaseUrl:      "quay.io",
							Organization: "solo-io",
						},
					},
				},
				TestContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:   "gcr.io",
							ProjectId: "solo-public-1010",
						},
					},
				},
			},
		},
		Config: &v1.BuildRunConfig{
			BuildEnvVars: &v1.BuildEnvVars{
				BuildId:       buildId,
				TaggedVersion: nonReleaseTagValue,
			},
			ComputedBuildVars: &v1.ComputedBuildVars{
				Release:         false,
				ImageTag:        buildId,
				ContainerPrefix: "gcr.io/solo-public-1010",
				Version:         buildId,
			},
		},
	}

	Context("InitializeBuildRun", func() {
		It("should fallback to default filename when no file specified and env var not set", func() {
			br, err := InitializeBuildRun("", &v1.BuildEnvVars{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("could not parse build spec: open solo-project.yaml: no such file or directory"))
			Expect(br).To(Equal(v1.BuildRun{}))
		})
		It("should read from file when provided", func() {
			br, err := InitializeBuildRun(relativePathToSampleConfig, explicitBuildEnvVars)
			Expect(err).ToNot(HaveOccurred())

			// aside on matching large structs
			// Gomega has MatchAllFields, but it requires you to fill in everything, including the XXXabcd fields made by protoc
			// An alternative is to pass IgnoreExtras with MatchFields
			// Either way, you have to write out a lot of stuff by hand
			// This is worthwhile in some cases because you can be explicit about how fields should match and you get
			// great, scoped feedback on errors
			Expect(br).To(MatchFields(IgnoreExtras, Fields{
				"Spec": PointTo(MatchFields(IgnoreExtras, Fields{
					"Config": PointTo(MatchFields(IgnoreExtras, Fields{
						"ReleaseContainerRegistry": PointTo(MatchFields(IgnoreExtras, Fields{
							"Registry": PointTo(MatchFields(IgnoreExtras, Fields{
								"Quay": PointTo(MatchFields(IgnoreExtras, Fields{
									"BaseUrl":      Equal("quay.io"),
									"Organization": Equal("solo-io"),
								})),
							})),
						})),
					})),
				})),
			}))
			// It is much easier to do this - stringify the struct
			// One shortcoming is that you only get +/- 5 chars of context
			// per: https://github.com/onsi/gomega/blob/master/format/format.go#L146
			// TODO(mitchdraft) gomega pr to modify charactersAroundMismatchToInclude (if not merged will make a util)
			Expect(br.String()).To(Equal(sampleConfigIngested.String()))
			// This was my first thought, but the error message is not human friendly, you have to compare large blocks of text
			Expect(br).To(Equal(sampleConfigIngested))
		})
	})
	Context("unit test isRelease", func() {
		It("should report release correctly", func() {
			ev1 := &v1.BuildEnvVars{
				TaggedVersion: "",
			}
			Expect(isRelease(ev1)).To(Equal(false))
			ev2 := &v1.BuildEnvVars{
				TaggedVersion: "v1.2.3",
			}
			Expect(isRelease(ev2)).To(Equal(true))
		})
	})
	Context("unit test setImageTag", func() {
		It("should report tag correctly for valid configs", func() {
			buildId := "12345"
			nonRelease := &v1.BuildEnvVars{
				TaggedVersion: "",
				BuildId:       buildId,
			}
			const rootTag = "v1.2.3"
			tag := rootTag
			Expect(setImageTag(&tag, nonRelease)).NotTo(HaveOccurred())
			Expect(tag).To(Equal(buildId))
			taggedVersion := rootTag
			version := "1.2.3"
			release := &v1.BuildEnvVars{
				TaggedVersion: taggedVersion,
				BuildId:       buildId,
			}
			Expect(setImageTag(&tag, release)).NotTo(HaveOccurred())
			Expect(tag).To(Equal(version))
			// we require semver
			release.TaggedVersion = "vabcdefg"
			Expect(setImageTag(&tag, release)).To(HaveOccurred())
			release.TaggedVersion = "vv"
			Expect(setImageTag(&tag, release)).To(HaveOccurred())
			release.TaggedVersion = "v with some space"
			Expect(setImageTag(&tag, release)).To(HaveOccurred())
		})
		It("should error for invalid configs", func() {
			buildId := "12345"
			nonReleaseNoBuildId := &v1.BuildEnvVars{
				TaggedVersion: "",
				BuildId:       "",
			}
			tag := ""
			Expect(setImageTag(&tag, nonReleaseNoBuildId)).To(HaveOccurred())
			taggedVersion := "1.2.3"
			releaseBadTaggedVersion := &v1.BuildEnvVars{
				TaggedVersion: taggedVersion,
				BuildId:       buildId,
			}
			Expect(setImageTag(&tag, releaseBadTaggedVersion)).To(HaveOccurred())
			releaseBadTaggedVersion.TaggedVersion = "v"
			Expect(setImageTag(&tag, releaseBadTaggedVersion)).To(HaveOccurred())
			releaseBadTaggedVersion.TaggedVersion = "a1.2.3"
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
							BaseUrl:   "gcr.io",
							ProjectId: "aproject",
						},
					},
				},
				TestContainerRegistry: nil,
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject"))
			release = true
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject"))
		})
		It("should error for invalid single container configs", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: nil,
				TestContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:   "gcr.io",
							ProjectId: "aproject",
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
							BaseUrl:   "gcr.io",
							ProjectId: "aproject",
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
			Expect(prefix).To(Equal("quay.io/an-org"))
			release = true
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject"))
		})
		It("should validate gcr container specs correctly", func() {
			prefix := ""
			release := false
			config := &v1.BuildConfig{
				ReleaseContainerRegistry: &v1.ContainerRegistry{
					Registry: &v1.ContainerRegistry_Gcr{
						Gcr: &v1.GoogleContainerRegistry{
							BaseUrl:   "gcr.io",
							ProjectId: "aproject",
						},
					},
				},
				TestContainerRegistry: nil,
			}
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/aproject"))
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.BaseUrl = ""
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.BaseUrl = "gcr.io"
			Expect(setContainerPrefix(&prefix, release, config)).ToNot(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.ProjectId = ""
			Expect(setContainerPrefix(&prefix, release, config)).To(HaveOccurred())
			config.ReleaseContainerRegistry.Registry.(*v1.ContainerRegistry_Gcr).Gcr.ProjectId = "aproject"
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
			Expect(prefix).To(Equal("quay.io/an-org"))
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
			Expect(prefix).To(Equal("docker.io"))
		})
	})
})
