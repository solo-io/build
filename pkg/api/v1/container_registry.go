package v1

import (
	"fmt"
	"github.com/pkg/errors"
)

func (r *ContainerRegistry) SetPrefixFromContainerRegistry(prefix *string) error {
	switch x := r.Registry.(type) {
	case *ContainerRegistry_DockerHub:
		return x.setRepoPrefix(prefix)
	case *ContainerRegistry_Quay:
		return x.setRepoPrefix(prefix)
	case *ContainerRegistry_Gcr:
		return x.setRepoPrefix(prefix)
	default:
		return fmt.Errorf("unrecognized type: %v", x)
	}
}

var (
	NoDockerOrgSpecifiedError = errors.Errorf("must provide an organization for docker repos")
	NoQuayOrgSpecifiedError = errors.Errorf("must provide an organization for quay repos")
	NoGcrProjectIdSpecifiedError = errors.Errorf("must provide a project id for gcr repos")
)

const (
	DockerBaseUrl = "docker.io"
	QuayBaseUrl = "quay.io"
	GcrBaseUrl = "gcr.io"
)

func (x *ContainerRegistry_DockerHub) setRepoPrefix(prefix *string) error {
	base, org := x.DockerHub.GetBaseUrl(), x.DockerHub.GetOrganization()
	if base == "" {
		base = DockerBaseUrl
	}
	if org == "" {
		return NoDockerOrgSpecifiedError
	}
	*prefix = fmt.Sprintf("%s/%s", base, org)
	return nil
}

func (x *ContainerRegistry_Quay) setRepoPrefix(prefix *string) error {
	base, org := x.Quay.GetBaseUrl(), x.Quay.GetOrganization()
	if base == "" {
		base = QuayBaseUrl
	}
	if org == "" {
		return NoQuayOrgSpecifiedError
	}
	*prefix = fmt.Sprintf("%s/%s", base, org)
	return nil
}

func (x *ContainerRegistry_Gcr) setRepoPrefix(prefix *string) error {
	base, proj := x.Gcr.GetBaseUrl(), x.Gcr.GetProjectId()
	if base == "" {
		base = GcrBaseUrl
	}
	if proj == "" {
		return NoGcrProjectIdSpecifiedError
	}
	*prefix = fmt.Sprintf("%s/%s", base, proj)
	return nil
}
