# Build

Tools for standardizing Solo.io's build system across repos.

This is a work-in-progress, likely to change.

[Original Proposal](https://github.com/solo-io/proposals/pull/2)

## Intent

Create a standard, easy way for repositories to define their build spec in a yaml file. Create a go library in `go-utils` for parsing the spec in a go struct. 

## Motivation

An example of a script that requires an artifact specification is here: `https://github.com/solo-io/go-utils/tree/master/githubutils`. This utility has been reused across several projects for uploading release assets, and demonstrates the value of distributing go scripts that solve common build and release tasks. This script could read `artifacts.yaml` to find the artifact spec, rather than requiring it to be constructed by the clients in a custom go struct. 

There are several other ways we could leverage the build specification:

* Automate the creation of PRs into the solo-io homebrew and gofish repositories, which require understanding the artifacts to construct the formulas. This is tricky today because artifact naming is not consistent across repos, and not all artifacts should be published to the formulas. 
* Autogenerate make targets and cloudbuild files. 


