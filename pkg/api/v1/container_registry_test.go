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
		Expect(prefix).To(Equal(dockerRepoUrl))
	})

	It("should handle quay", func() {
		orgName := "myorg"
		baseQuayUrl := "quay.io"
		cr := ContainerRegistry{
			Registry: &ContainerRegistry_Quay{
				Quay: &QuayRegistry{
					Organization: orgName,
					BaseUrl:      baseQuayUrl,
				},
			},
		}
		Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
		Expect(prefix).To(Equal("quay.io/myorg"))
	})

	It("should handle gcr", func() {
		baseGcrUrl := "gcr.io"
		cr := ContainerRegistry{
			Registry: &ContainerRegistry_Gcr{
				Gcr: &GoogleContainerRegistry{
					ProjectName: "myproject",
					BaseUrl:     baseGcrUrl,
				},
			},
		}
		Expect(cr.SetPrefixFromContainerRegistry(&prefix)).NotTo(HaveOccurred())
		Expect(prefix).To(Equal("gcr.io/myproject"))
	})
})
