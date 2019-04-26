package v1

import (
	"fmt"
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

const dockerRepoUrl = "docker.io"

func (x *ContainerRegistry_DockerHub) setRepoPrefix(prefix *string) error {
	*prefix = dockerRepoUrl
	return nil
}

func (x *ContainerRegistry_Quay) setRepoPrefix(prefix *string) error {
	if x.Quay.BaseUrl == "" {
		return fmt.Errorf("must provide a base url for quay repos")
	}
	if x.Quay.Organization == "" {
		return fmt.Errorf("must provide an organization for quay repos")
	}
	*prefix = fmt.Sprintf("%s/%s", x.Quay.BaseUrl, x.Quay.Organization)
	return nil
}

func (x *ContainerRegistry_Gcr) setRepoPrefix(prefix *string) error {
	if x.Gcr.BaseUrl == "" {
		return fmt.Errorf("must provide a base url for gcr repos")
	}
	if x.Gcr.ProjectId == "" {
		return fmt.Errorf("must provide a project name for gcr repos")
	}
	*prefix = fmt.Sprintf("%s/%s", x.Gcr.BaseUrl, x.Gcr.ProjectId)
	return nil
}
