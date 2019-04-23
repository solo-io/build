
# These are each evaluated a single time when the makefile is loaded
RELEASE := $(shell go run cmd/read_env/main.go parse-env release)
IMAGE_TAG := $(shell go run cmd/read_env/main.go parse-env image-tag)
CONTAINER_PREFIX := $(shell go run cmd/read_env/main.go parse-env container-prefix)

.PHONY: print-release-vals
print-release-vals:
	echo $(RELEASE)
	echo $(IMAGE_TAG)
	echo $(CONTAINER_PREFIX)
