.PHONY: generate
generate:
	protoc -I=. -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --gogo_out=./pkg api/v1/build.proto

.PHONY: call-sample
call-sample:
	TAG_VERSION=v1.2.3 make -f examples/sample.makefile print-release-vals
