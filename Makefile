# vi: set ft=make:

PACKAGE = github.com/johndistasio/cauldron

GIT_TAG    = $(shell git describe --tags --always 2>/dev/null)
GO_LDFLAGS = $(addprefix -X main.,version=$(GIT_TAG))

.PHONY: default
default: clean fmt lint test build smoketest

.PHONY: build
build:
	@go mod download
	@CGO_ENABLED=0 go build -ldflags '$(GO_LDFLAGS)' -o cauldron $(PACKAGE)

.PHONY: test
test:
	@go mod download
	@go test -v ./...

.PHONY: smoketest
smoketest:
	@bash ./smoketest.sh

.PHONY: fmt
fmt:
	@files=$$(go fmt $(PACKAGE)); \
	if [ -n "$$files" ]; then \
	  echo "Incorrect formatting on:"; \
	  echo $$files; \
	  exit 1; \
	fi

.PHONY: lint
lint:
	@golint -set_exit_status

.PHONY: clean
clean:
	@rm cauldron 2> /dev/null || :
