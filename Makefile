include .env

SHELL                := /bin/bash
ROOTDIR              ?= $(CURDIR)
GO                   := $(shell command -v go 2> /dev/null)
GOFMT 				 := $(shell command -v gofmt 2> /dev/null)
PACKAGE              := github.com/zenoss/event-management-service
BIN                  := $(GOPATH)/bin
BASE                 := $(GOPATH)/src/$(PACKAGE)
DOCKER_COMPOSE       := /usr/local/bin/docker-compose
LOCAL_USER_ID        := $(shell id -u)
ZENKIT_BUILD_VERSION := 1.11.0
BUILD_IMG            := zenoss/zenkit-build:$(ZENKIT_BUILD_VERSION)
COVERAGE_DIR         := coverage
DOCKER_PARAMS        := --rm -v $(ROOTDIR):/go/src/$(PACKAGE):rw \
							-v /var/run/docker.sock:/var/run/docker.sock \
							-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
							-w /go/src/$(PACKAGE)
DOCKER_CMD           := docker run -t $(DOCKER_PARAMS) $(BUILD_IMG)

DOCKER_COMPOSE_BASE  := $(DOCKER_COMPOSE)
ifdef PROJECT_NAME
DOCKER_COMPOSE_BASE  += -p $(PROJECT_NAME)
endif

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: default
default: init-containerized

.PHONY: docker-compose
docker-compose: $(DOCKER_COMPOSE)

$(DOCKER_COMPOSE): DOCKER_COMPOSE_VERSION := 1.14.0
$(DOCKER_COMPOSE):
	@if [ ! -w $(@D) ]; then echo 'No docker-compose found. Please run "sudo make docker-compose" to install it.'; exit 2; else true; fi
	@curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-`uname -s`-`uname -m` > $@
	@chmod +x $@

$(BIN):
	@mkdir -p $@
$(BIN)/%: $(BIN) | $(BASE)
	@[ -x $@ ] || echo "$(M)" Installing $* && go get $(REPOSITORY)

GINKGO = $(BIN)/ginkgo
$(BIN)/ginkgo: REPOSITORY=github.com/onsi/ginkgo/v2/ginkgo

GOLINT = $(BIN)/golint
$(BIN)/golint: REPOSITORY=golang.org/x/lint/golint

GOCOV = $(BIN)/gocov
$(BIN)/gocov: REPOSITORY=github.com/axw/gocov/...

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: REPOSITORY=github.com/AlekSi/gocov-xml

$(BASE):

vendor: go.mod go.sum | $(BASE) ; $(info $(M) retrieving dependencies…)
	$Q cd $(BASE) && $(GO) mod vendor
	@touch $@

.PHONY: vendor-update
vendor-update: vendor | $(BASE)
ifeq "$(origin PKG)" "command line"
	$(info $(M) updating $(PKG) dependency…)
	$Q cd $(BASE) && $(GO) get $(PKG) && $(GO) mod vendor
else
	$(info $(M) updating all dependencies…)
	$Q cd $(BASE) && $(GO) get ./... && $(GO) mod vendor
endif
	@touch vendor

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: lint
lint: vendor | $(BASE) $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: build
build: export COMMIT_SHA ?= $(shell git rev-parse HEAD)
build: export GIT_BRANCH ?= $(shell git symbolic-ref HEAD | sed -e "s/^refs\/heads\///")
build: export PULL_REQUEST = ${ghprbPullLink}
build: | $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) build event-management-service

.PHONY: run
run: $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) up --build

.PHONY: down
down: $(DOCKER_COMPOSE)
	$(DOCKER_COMPOSE_BASE) down

.PHONY: check test tests
check test tests: unit-test-containerized

.PHONY: unit-test
unit-test: COVERAGE_PROFILE := coverprofile.out
unit-test: COVERAGE_HTML    := $(COVERAGE_DIR)/index.html
unit-test: COVERAGE_XML     := $(COVERAGE_DIR)/coverage.xml
unit-test: fmt
	@mkdir -p $(COVERAGE_DIR)
	@$(GINKGO) \
		-r \
		-mod vendor \
		-cover \
		-covermode=count \
		--skip-package vendor \
		-tags='integration' \
		--junit-report=junit.xml
	@$(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@$(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: unit-test-containerized
unit-test-containerized: export ZENKIT_BUILD_VERSION ?= $(ZENKIT_BUILD_VERSION)
unit-test-containerized: $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) -f docker-compose.test.yml run $(DOCKER_PARAMS) test make unit-test; \
		$(DOCKER_COMPOSE_BASE) -f docker-compose.test.yml down

.PHONY: clean
clean:: down
	rm -f junit.xml
	rm -f coverprofile.out
	rm -rf $(COVERAGE_DIR)

.PHONY: mrclean
mrclean: clean
	rm -rf vendor

.git: vendor
	@git init
	@git add .; git commit -m "Initial commit"

.PHONY: init
init: .git

.PHONY: init-containerized
init-containerized:
	@$(DOCKER_CMD) make vendor
