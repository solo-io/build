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

	It("should handle docker", func() {
		cr := ContainerRegistry{
			Registry: &ContainerRegistry_DockerHub{},
		}
		Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
		Expect(prefix).To(Equal(DockerBaseUrl))
	})

	It("should handle quay without base url", func() {
		orgName := "myorg"
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
		orgName := "myorg"
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

	It("should handle gcr without base url", func() {
		cr := ContainerRegistry{
			Registry: &ContainerRegistry_Gcr{
				Gcr: &GoogleContainerRegistry{
					ProjectId: "myproject",
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
					ProjectId: "myproject",
					BaseUrl:   "other.gcr.io",
				},
			},
		}
		Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
		Expect(prefix).To(Equal("other.gcr.io/myproject"))
	})
})
