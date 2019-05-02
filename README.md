# Build

Tools for standardizing Solo.io's build system across repos.

## Intent

Create a standard, easy way for repositories to define their build spec in a yaml file. Create a go library in `go-utils` for parsing the spec in a go struct. 

## Motivation

An example of a script that requires an artifact specification is here: `https://github.com/solo-io/go-utils/tree/master/githubutils`. This utility has been reused across several projects for uploading release assets, and demonstrates the value of distributing go scripts that solve common build and release tasks. This script could read `artifacts.yaml` to find the artifact spec, rather than requiring it to be constructed by the clients in a custom go struct. 

There are several other ways we could leverage the build specification:

* Automate the creation of PRs into the solo-io homebrew and gofish repositories, which require understanding the artifacts to construct the formulas. This is tricky today because artifact naming is not consistent across repos, and not all artifacts should be published to the formulas. 
* Autogenerate make targets and cloudbuild files. 


# Usage in a Makefile

One of the goals of this tool is to move release logic from makefiles into golang libraries and scripts. Golang has more robust language and version control features than make, but make provides a more flexible and convenient user experience in some cases.
It is expected that a major use of this tool will be augmenting makefile logic. It should be called as a script (`go run cmd/build-my-project/main.go <args>`) so that you can lock the library into a particular version. 

- Run from script
  - Create a `build.go` script somewhere in your project
  - Lock the `github.com/solo-io/build` import to a specific released version

*go script*:
```go
import "github.com/solo-io/build/pkg/cli"

func main(){
	cli.Run()
}
```

*usage in makefile*
```make
# These are each evaluated a single time when the makefile is loaded
RELEASE := $(shell go run cmd/read_env/main.go parse-env release)
VERSION := $(shell go run cmd/read_env/main.go parse-env version)
IMAGE_TAG := $(shell go run cmd/read_env/main.go parse-env image-tag)
CONTAINER_PREFIX := $(shell go run cmd/read_env/main.go parse-env container-prefix)

.PHONY: verify-args
verify-args:
    go run buildcmd/main.go validate-operating-parameters \
    	$(RELEASE) \
    	$(VERSION) \
    	$(CONTAINER_REPO_ORG) \
    	$(IMAGE_TAG)
    	
every-other-make-target: verify-args <other-dependencies>
```
- note that all your make targets should depend on the `verify-args` target so that the build can exit if there was a configuration error

## Compute release
`RELEASE = build parse-env release`

## Compute image tag
`IMAGE_TAG = build parse-env image-tag`

