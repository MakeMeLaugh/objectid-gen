REVISION := $(shell git rev-parse --short HEAD || echo "unknown")
BRANCH := $(shell git rev-parse --abbrev-ref HEAD || echo "unknown")
BUILD_TAG := $(shell git describe --tags --abbrev=0 || echo "unknown")
LAST_TAG_REVISION := $(shell git rev-parse --short $(BUILD_TAG) || echo "unknown")

ifneq ($(REVISION),$(LAST_TAG_REVISION))
	# we are not building the tag
	BUILD_TAG := $(BRANCH)-$(REVISION)
endif

LD_FLAGS := '-s -w -X "main.applicationName=$(PROJECT_NAME)" \
-X "main.applicationVersion=$(BUILD_TAG)" \
-X "main.buildAt=$(shell date +"%FT%T %Z")" \
-X "main.buildFrom=$(REVISION)"'

BUILD_PLATFORMS ?= darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

.PHONY: clean setup build
setup:
	$Q mkdir -p bin/

build: clean setup
	$Q set -e ; for platform in $(BUILD_PLATFORMS) ; do \
  printf "Building towards: $${platform%/*} (architecture: $${platform#*/})" ; \
  GOOS="$${platform%/*}" \
  GOARCH="$${platform#*/}" \
  go build \
  -ldflags=$(LD_FLAGS) \
  -o "bin/$(PROJECT_NAME)-$${platform%/*}-$${platform#*/}-$(BUILD_TAG)$$([ "$${platform%/*}" = "windows" ] && printf ".exe" || true)" && \
  printf "...OK\n" || printf "...FAIL\n" ; \
done ; set +e

clean:
	$Q rm -rf bin/*
