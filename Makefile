.PHONY: generate
generate:
	protoc -I=. -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --gogo_out=./pkg api/v1/build.proto
