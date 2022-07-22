ZENKIT_BUILD_VERSION := 1.12.0
PACKAGE              := github.com/zenoss/zenkit
ROOTDIR              ?= $(CURDIR)
NETRC                ?= ~/.netrc
GO                   := $(shell command -v go 2> /dev/null)
GOFMT 				 := $(shell command -v gofmt 2> /dev/null)
PACKAGE              := github.com/zenoss/zenkit
LOCAL_USER_ID        := $(shell id -u)
BUILD_IMG            := zenoss/zenkit-build:$(ZENKIT_BUILD_VERSION)
COVERAGE_DIR         := coverage
DOCKER_COMPOSE       := /usr/local/bin/docker-compose

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

export GOPRIVATE = github.com/zenoss

# Native

.PHONY: all
all: tools dependencies check test

# Containerized

.PHONY: docker-compose
docker-compose: $(DOCKER_COMPOSE)

$(DOCKER_COMPOSE): DOCKER_COMPOSE_VERSION := 1.14.0
$(DOCKER_COMPOSE):
	@if [ ! -w $(@D) ]; then echo 'No docker-compose found. Please run "sudo make docker-compose" to install it.'; exit 2; else true; fi
	@curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-`uname -s`-`uname -m` > $@
	@chmod +x $@

.PHONY: all-containerized
all-containerized: export ZENKIT_BUILD_VERSION ?= $(ZENKIT_BUILD_VERSION)
all-containerized: $(DOCKER_COMPOSE)
	$Q $(DOCKER_COMPOSE) -f docker-compose.test.yml run \
		-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
		-v $(ROOTDIR):/go/src/$(PACKAGE):rw \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(NETRC):/home/user/.netrc \
		-w /go/src/$(PACKAGE) \
		--rm \
		zenkit-build \
		make all

# Tools

ifeq ($(GO),)
GOBIN := go/bin
else
GOBIN := $(shell $(GO) env GOPATH)/bin
endif

$(GOBIN):
	@mkdir -p $@

GOLINT = $(GOBIN)/golint
$(GOBIN)/golint: PACKAGE=golang.org/x/lint/golint

GINKGO = $(GOBIN)/ginkgo
$(GOBIN)/ginkgo: PACKAGE=github.com/onsi/ginkgo/v2/ginkgo

GOCOVMERGE = $(GOBIN)/gocovmerge
$(GOBIN)/gocovmerge: PACKAGE=github.com/wadey/gocovmerge

GOCOV = $(GOBIN)/gocov
$(GOBIN)/gocov: PACKAGE=github.com/axw/gocov/...

GOCOVXML = $(GOBIN)/gocov-xml
$(GOBIN)/gocov-xml: PACKAGE=github.com/AlekSi/gocov-xml

.PHONY: tools
tools: $(GOLINT) $(GINKGO) $(GOCOVMERGE) $(GOCOV) $(GOCOVXML) ; $(info $(M) installed tools) @ ## Install tools.

# Avoid making tools dependencies.
$(GOBIN)/%: | $(GOBIN) ; $(info $(M) installing $(PACKAGE)…)
	tmp=$$(mktemp -d); \
		env GO111MODULE=off GOPATH=$$tmp GOBIN=$(GOBIN) $(GO) get $(PACKAGE) || ret=$$?; \
		rm -rf $$tmp ; exit $$ret

# Dependencies

.PHONY: dependencies
dependencies: ; $(info $(M) downloading dependencies…) @ ## Install dependencies with go mod download.
	$Q go mod download

# Static Checks

# TODO: Add lint to check dependencies. Currently it complains about missing comments.

.PHONY: check
check: fmt-check vet ; $(info $(M) static checks complete) @ ## Run all static checks.

.PHONY: fmt-check
fmt-check: ; $(info $(M) running gofmt…) @ ## Check all files with gofmt.
	$Q [ -z "$$($(GOFMT) -l -s .)" ] || ($(GOFMT) -d -e -s . ; exit 1)

.PHONY: vet
vet: ; $(info $(M) running go vet…) @ ## Run go vet on all packages.
	$Q $(GO) vet $$($(GO) list ./...)

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint…) @ ## Check packages with golint.
	$Q $(GOLINT) -set_exit_status $$($(GO) list ./...)

# Tests

.PHONY: test
test: COVERAGE_PROFILE := $(COVERAGE_DIR)/profile.out
test: COVERAGE_HTML    := $(COVERAGE_DIR)/index.html
test: COVERAGE_XML     := $(COVERAGE_DIR)/coverage.xml
test: | $(GINKGO) $(GOCOV) $(GOCOVXML) ; $(info $(M) running tests...) @ ## Execute tests with ginkgo.
	$Q mkdir -p $(COVERAGE_DIR)
	$Q $(GINKGO) -r -cover -coverprofile=profile.out -covermode=count -output-dir=$(COVERAGE_DIR) --junit-report=junit.xml
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

# Fixes

.PHONY: fix
fix: mod-tidy fmt

.PHONY: mod-tidy
mod-tidy: ; $(info $(M) running go mod tidy…) @ ## Update dependencies with go mod tidy.
	$Q $(GO) mod tidy

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all files.
	$Q $(GOFMT) -l -w .

# Miscellaneous

.PHONY: clean
clean: ; $(info $(M) cleaning…) @ ## Cleanup everything.
	$Q rm -rf $(COVERAGE_DIR)

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
