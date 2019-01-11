# Copyright (C) 2016-2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP = paleta

VERSION=$(shell \
        grep "const Version" pkg/version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.rc]//g" \
	|sed -e "s/ //g")

SHELL = /bin/bash

DIR = $(shell pwd)

GO = go

DOCKER = docker
NAMESPACE = nimbus

GOX = gox

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

MAKE_COLOR=\033[33;01m%-20s\033[0m

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

clean: ## Cleanup
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -fr diablod diabloctl diabloadm

.PHONY: tools
tools:
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck
	@go get -u github.com/mitchellh/gox

init: tools ## Install requirements

.PHONY: build
build: ## Make binary
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GO) build -o paleta github.com/pilotariak/paleta

.PHONY: test
test: ## Launch unit tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@go test -tags '$(BUILD_TAGS)' -ldflags '$(GO_LDFLAGS)' -gcflags '$(GO_GCFLAGS)' \
		-v $$(go list ./... | grep -v /vendor/)

.PHONY: test-verbose
test-verbose: ## Launch unit tests with logs
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests with verbosity$(NO_COLOR)"
	@go test -tags '$(BUILD_TAGS)' -ldflags '$(GO_LDFLAGS)' -gcflags '$(GO_GCFLAGS)' \
		-v $$(go list ./... | grep -v /vendor/) -args --alsologtostderr -v 9

.PHONY: test
test-pkg: ## Launch unit tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests for $(pkg) $(NO_COLOR)"
	@go test -tags '$(BUILD_TAGS)' -ldflags '$(GO_LDFLAGS)' -gcflags '$(GO_GCFLAGS)' $(pkg)

.PHONY: test-verbose
test-pkg-verbose: ## Launch unit tests with logs
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests for $(pkg) with verbosity$(NO_COLOR)"
	@go test -tags '$(BUILD_TAGS)' -ldflags '$(GO_LDFLAGS)' -gcflags '$(GO_GCFLAGS)' \
		-v $(pkg) -args --alsologtostderr -v 9

.PHONY: coverage
coverage: ## Launch code coverage
	@echo -e "$(OK_COLOR)[$(APP)] Code coverage $(NO_COLOR)"
	@go test -cover -tags '$(BUILD_TAGS)' -ldflags '$(GO_LDFLAGS)' -gcflags '$(GO_GCFLAGS)' \
		-v $$(go list ./... | grep -v /vendor/)

.PHONY: integration
integration: ## Launch integration tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch integration tests $(NO_COLOR)"
	@govendor test +local -tags=integration

.PHONY: lint
lint: ## Launch golint
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet: ## Launch go vet
	@$(foreach file,$(SRCS),$(GO) vet $(file) || exit;)

.PHONY: errcheck
errcheck: ## Launch go errcheck
	@echo -e "$(OK_COLOR)[$(APP)] Go Errcheck $(NO_COLOR)"
	@$(foreach pkg,$(PKGS),errcheck $(pkg) $(glide novendor) || exit;)

binaries: changelog ## Make all binaries for a new release
	@echo -e "$(OK_COLOR)[$(APP)] Create binaries $(NO_COLOR)"
	$(GOX) -output=paleta-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/pilotariak/paleta

docker-build: ## Build Docker image
	@echo "$()Docker build $(NAMESPACE)/$(APP):$(VERSION)$(NO_COLOR)"
	@docker build -t $(NAMESPACE)/$(APP):$(VERSION) .

docker-run: ## Run the Docker image
	@echo "$(OK_COLOR)Docker run $(NAMESPACE)/$(APP):$(VERSION)$(NO_COLOR)"
	docker run --rm=true \
		-v `pwd`:/etc/trinquet \
		$(NAMESPACE)/$(APP):$(VERSION)

# for goprojectile
.PHONY: gopath
gopath:
	@echo `pwd`:`pwd`/vendor
