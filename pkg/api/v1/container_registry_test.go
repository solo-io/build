package v1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("container registry", func() {

	var (
		prefix string
	)

	BeforeEach(func() {
		prefix = ""
	})

	Context("docker", func() {
		orgName := "myorg"

		It("should error if oneof is empty", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_DockerHub{},
			}
			err := cr.SetPrefixFromContainerRegistry(&prefix)
			Expect(err).To(BeEquivalentTo(NoDockerOrgSpecifiedError))
		})

		It("should error if orgName not specified", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_DockerHub{
					DockerHub: &DockerHubRegistry{},
				},
			}
			err := cr.SetPrefixFromContainerRegistry(&prefix)
			Expect(err).To(BeEquivalentTo(NoDockerOrgSpecifiedError))
		})

		It("should handle docker without base url", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_DockerHub{
					DockerHub: &DockerHubRegistry{
						Organization: orgName,
					},
				},
			}
			Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
			Expect(prefix).To(Equal("docker.io/myorg"))
		})

		It("should handle docker with specified base url", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_DockerHub{
					DockerHub: &DockerHubRegistry{
						BaseUrl: "other.docker.io",
						Organization: orgName,
					},
				},
			}
			Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
			Expect(prefix).To(Equal("other.docker.io/myorg"))
		})
	})



	Context("quay", func() {

		orgName := "myorg"

		It("should error if oneof is empty", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Quay{},
			}
			err := cr.SetPrefixFromContainerRegistry(&prefix)
			Expect(err).To(BeEquivalentTo(NoQuayOrgSpecifiedError))
		})

		It("should error if orgName not specified", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Quay{
					Quay: &QuayRegistry{
					},
				},
			}
			err := cr.SetPrefixFromContainerRegistry(&prefix)
			Expect(err).To(BeEquivalentTo(NoQuayOrgSpecifiedError))
		})

		It("should handle quay without base url", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Quay{
					Quay: &QuayRegistry{
						Organization: orgName,
					},
				},
			}
			Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
			Expect(prefix).To(Equal("quay.io/myorg"))
		})

		It("should handle quay with specified base url", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Quay{
					Quay: &QuayRegistry{
						Organization: orgName,
						BaseUrl:      "other.quay.io",
					},
				},
			}
			Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
			Expect(prefix).To(Equal("other.quay.io/myorg"))
		})
	})

	Context("gcr", func() {

		projectId := "myproject"

		It("should error if oneof is empty", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Gcr{},
			}
			err := cr.SetPrefixFromContainerRegistry(&prefix)
			Expect(err).To(BeEquivalentTo(NoGcrProjectIdSpecifiedError))
		})

		It("should error if projectId not specified", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Gcr{
					Gcr: &GoogleContainerRegistry{
					},
				},
			}
			err := cr.SetPrefixFromContainerRegistry(&prefix)
			Expect(err).To(BeEquivalentTo(NoGcrProjectIdSpecifiedError))
		})

		It("should handle gcr without base url", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Gcr{
					Gcr: &GoogleContainerRegistry{
						ProjectId: projectId,
					},
				},
			}
			Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
			Expect(prefix).To(Equal("gcr.io/myproject"))
		})

		It("should handle gcr with specified base url", func() {
			cr := ContainerRegistry{
				Registry: &ContainerRegistry_Gcr{
					Gcr: &GoogleContainerRegistry{
						ProjectId: projectId,
						BaseUrl:   "other.gcr.io",
					},
				},
			}
			Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
			Expect(prefix).To(Equal("other.gcr.io/myproject"))
		})
	})




})
