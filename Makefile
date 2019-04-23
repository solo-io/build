.PHONY: generate
generate:
	protoc -I=. -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --gogo_out=./pkg api/v1/build.proto

.PHONY: call-sample-release
call-sample-release:
	SOLOBUILD_CONFIG_FILE=./examples/sample-solo-project.yaml TAG_VERSION=v1.2.3 BUILD_ID=12345 make -f examples/sample.makefile print-release-vals

.PHONY: call-sample-nonrelease
call-sample-nonrelease:
	SOLOBUILD_CONFIG_FILE=./examples/sample-solo-project.yaml BUILD_ID=12345 make -f examples/sample.makefile print-release-vals
