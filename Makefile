# vi: set ft=make:

VERSION = 0.11.0
PACKAGE = github.com/johndistasio/cauldron

GIT_REVISION = $(shell git rev-parse --short HEAD 2>/dev/null)
GIT_TAG      = $(shell git describe --tags --always 2>/dev/null)

GO_ARCH    = $(shell go env GOARCH)
GO_OS      = $(shell go env GOOS)
GO_VERSION = $(shell go version | awk '{print $$3}' | tr -d 'go')
GO_LDFLAGS = $(addprefix -X $(PACKAGE)/version.,version=$(VERSION) revision=$(GIT_REVISION) tag=$(GIT_TAG) goarch=$(GO_ARCH) goos=$(GO_OS) goversion=$(GO_VERSION))

TARBALL_EXCLUDE = $(addprefix --exclude=,build rpmbuild .git .idea .vagrant)

.PHONY: test build smoketest

default: clean build

archive:
	@mkdir -p build/
	tar $(TARBALL_EXCLUDE) -czvf build/cauldron-$(VERSION).tar.gz .

build:
	@mkdir -p build/
	CGO_ENABLED=0 go build -ldflags '$(GO_LDFLAGS)' -a -o build/cauldron $(PACKAGE)

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

smoketest:
	bash ./smoketest.sh

fmt:
	@files=$$(go fmt $(PACKAGE)); \
	if [ -n "$$files" ]; then \
	  echo "Incorrect formatting on:"; \
	  echo $$files; \
	  exit 1; \
	fi

lint:
	golint -set_exit_status

clean:
	@rm -rf build/
