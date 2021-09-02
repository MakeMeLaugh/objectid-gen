PROJECT_NAME := objectid-gen
.DEFAULT_GOAL := build

REVISION := $(shell git rev-parse --short HEAD || echo "unknown")
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TAG := $(shell git describe --tags --abbrev=0)
LAST_TAG_REVISION := $(shell git rev-parse --short $(BUILD_TAG) || echo "unknown")

ifneq ($(REVISION),$(LAST_TAG_REVISION))
	# we are not building the tag
	BUILD_TAG := $(BRANCH)-$(REVISION)
endif


LD_FLAGS := '-s -w -X "main.applicationName=$(PROJECT_NAME)" \
-X "main.applicationVersion=$(BUILD_TAG)" \
-X "main.buildAt=$(shell date +"%FT%T")" \
-X "main.buildFrom=$(REVISION)"'

BUILD_PLATFORMS ?= darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

.PHONY: clean setup build
setup:
	$Q mkdir -p bin/

build: clean setup
	$Q for platform in $(BUILD_PLATFORMS) ; do \
  echo "BUILD TARGET OS: $${platform%/*} (ARCH: $${platform#*/})" ; \
  GOOS="$${platform%/*}" \
  GOARCH="$${platform#*/}" \
  go build \
  -ldflags=$(LD_FLAGS) \
  -o "bin/$(PROJECT_NAME)-$${platform%/*}-$${platform#*/}-$(BUILD_TAG)$$([ "$${platform%/*}" = "windows" ] && echo ".exe" || echo "")" ; \
done

clean:
	$Q rm -rf bin/*