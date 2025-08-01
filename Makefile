#   -*- mode: makefile; -*-
# Copyright The OpenTelemetry Authors
#
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
#
# This Makefile has rules to generate go code for otlp exporter. It does it by
# copying the proto files from `opentelemetry-proto` (which is a submodule that
# needs to be checked out) into `gen/proto`, changing the go_package option to
# a valid string, generating the go files and finally copying the files into
# the module. The files are not generated in place, because protoc generates a
# too-deep directory structure.
#
# Currently, all the generated code is place in the base directory.
#
# Prereqs: docker, rsync (for copying back the generated files).

PROTOC_VERSION := 3.14.0

TOOLS_MOD_DIR                   := ./internal/tools
PROTOBUF_VERSION                := v1*
OTEL_PROTO_SUBMODULE            := opentelemetry-proto
GEN_TEMP_DIR                    := gen
SUBMODULE_PROTO_FILES           := $(wildcard $(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/*/$(PROTOBUF_VERSION)/*.proto) $(wildcard $(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/*/$(PROTOBUF_VERSION)/*.proto)

ifeq ($(strip $(SUBMODULE_PROTO_FILES)),)
$(error Submodule at $(OTEL_PROTO_SUBMODULE) is not checked out, use "git submodule update --init")
endif

GO                := go
GO_MOD_ROOT       := go.opentelemetry.io/proto
ALL_GO_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | sort)
ALL_GO_SUB_MOD_DIRS := $(shell find . -type f -name 'go.mod' -mindepth 2 -exec dirname {} \; | sort)
OTEL_GO_MOD_DIRS := $(filter-out $(TOOLS_MOD_DIR), $(ALL_GO_SUB_MOD_DIRS))
TIMEOUT = 60

PROTOBUF_GEN_DIR  := opentelemetry-proto-gen
PROTOBUF_TEMP_DIR := $(GEN_TEMP_DIR)/go

PROTO_SOURCE_DIR   := $(GEN_TEMP_DIR)/proto
SOURCE_PROTO_FILES := $(subst $(OTEL_PROTO_SUBMODULE),$(PROTO_SOURCE_DIR),$(SUBMODULE_PROTO_FILES))
OTLP_OUTPUT_DIR    := otlp

PROTOSLIM_SOURCE_DIR   := $(GEN_TEMP_DIR)/slim/proto
SOURCE_PROTOSLIM_FILES := $(subst $(OTEL_PROTO_SUBMODULE),$(PROTOSLIM_SOURCE_DIR),$(SUBMODULE_PROTO_FILES))
OTLPSLIM_OUTPUT_DIR    := slim/otlp

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

DEPENDENCIES_DOCKERFILE=./dependencies.Dockerfile
OTEL_DOCKER_PROTOBUF := $(shell awk '$$4=="build-protobuf" {print $$2}' $(DEPENDENCIES_DOCKERFILE))

PROTOC := docker run --rm -u ${shell id -u} -v${PWD}:${PWD} -w${PWD} ${OTEL_DOCKER_PROTOBUF} --proto_path="$(PROTO_SOURCE_DIR)"
PROTOC_SLIM := docker run --rm -u ${shell id -u} -v${PWD}:${PWD} -w${PWD} ${OTEL_DOCKER_PROTOBUF} --proto_path="$(PROTOSLIM_SOURCE_DIR)"

.DEFAULT_GOAL := protobuf

# Tools

TOOLS = $(CURDIR)/.tools

$(TOOLS):
	@mkdir -p $@
$(TOOLS)/%: | $(TOOLS)
	cd $(TOOLS_MOD_DIR) && \
	$(GO) build -o $@ $(PACKAGE)

MULTIMOD = $(TOOLS)/multimod
$(TOOLS)/multimod: PACKAGE=go.opentelemetry.io/build-tools/multimod

.PHONY: tools
tools: $(MULTIMOD)

.PHONY: protobuf
protobuf: protobuf-source gen-otlp-protobuf copy-otlp-protobuf gen-otlp-protobuf-slim copy-otlp-protobuf-slim

.PHONY: protobuf-source
protobuf-source: $(SOURCE_PROTO_FILES)

# The sed expression for replacing the go_package option in proto
# file with a one that's valid for us.
SED_EXPR := 's,go_package = "github.com/open-telemetry/opentelemetry-proto/gen/go/,go_package = "$(GO_MOD_ROOT)/$(OTLP_OUTPUT_DIR)/,'

# This copies proto files from submodule into $(PROTO_SOURCE_DIR),
# thus satisfying the $(SOURCE_PROTO_FILES) prerequisite. The copies
# have their package name replaced by go.opentelemetry.io/proto.
$(PROTO_SOURCE_DIR)/%.proto: $(OTEL_PROTO_SUBMODULE)/%.proto
	@ \
	mkdir -p $(@D); \
	sed -e $(SED_EXPR) "$<" >"$@.tmp"; \
	mv "$@.tmp" "$@"

# The sed expression for replacing the go_package option in proto
# file with a one that's valid for us.
SED_EXPR_SLIM := 's,go_package = "go.opentelemetry.io/proto/otlp/,go_package = "$(GO_MOD_ROOT)/$(OTLPSLIM_OUTPUT_DIR)/,'

# This copies proto files from submodule into $(PROTO_SOURCE_DIR),
# thus satisfying the $(SOURCE_PROTOSLIM_FILES) prerequisite. The copies
# have their package name replaced by go.opentelemetry.io/proto.
$(PROTOSLIM_SOURCE_DIR)/%.proto: $(OTEL_PROTO_SUBMODULE)/%.proto
	@ \
	mkdir -p $(@D); \
	sed -e $(SED_EXPR_SLIM) "$<" >"$@.tmp"; \
	mv "$@.tmp" "$@"

.PHONY: gen-otlp-protobuf
gen-otlp-protobuf: $(SOURCE_PROTO_FILES)
	rm -rf ./$(PROTOBUF_TEMP_DIR)
	mkdir -p ./$(PROTOBUF_TEMP_DIR)
	$(foreach file,$(SOURCE_PROTO_FILES),$(call exec-command,$(PROTOC) $(PROTO_INCLUDES) --go_out=./$(PROTOBUF_TEMP_DIR) $(file)))
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/trace/v1/trace_service_http.yaml:./$(PROTOBUF_TEMP_DIR) --go_out=./$(PROTOBUF_TEMP_DIR) --go-grpc_out=./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/metrics/v1/metrics_service_http.yaml:./$(PROTOBUF_TEMP_DIR) --go_out=./$(PROTOBUF_TEMP_DIR) --go-grpc_out=./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/logs/v1/logs_service_http.yaml:./$(PROTOBUF_TEMP_DIR) --go_out=./$(PROTOBUF_TEMP_DIR) --go-grpc_out=./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/logs/v1/logs_service.proto
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/profiles/v1development/profiles_service_http.yaml:./$(PROTOBUF_TEMP_DIR) --go_out=./$(PROTOBUF_TEMP_DIR) --go-grpc_out=./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/profiles/v1development/profiles_service.proto


.PHONY: copy-otlp-protobuf
copy-otlp-protobuf:
	rm -rf ./$(OTLP_OUTPUT_DIR)/*/
	@rsync -a $(PROTOBUF_TEMP_DIR)/go.opentelemetry.io/proto/otlp/ ./$(OTLP_OUTPUT_DIR)
	@git restore $$(git ls-files --deleted | grep -E 'go\.(mod|sum)$$')
	cd ./$(OTLP_OUTPUT_DIR)	&& go mod tidy

.PHONY: gen-otlp-protobuf-slim
gen-otlp-protobuf-slim: $(SOURCE_PROTOSLIM_FILES)
	rm -rf ./$(PROTOBUF_TEMP_DIR)
	mkdir -p ./$(PROTOBUF_TEMP_DIR)
	$(foreach file,$(SOURCE_PROTOSLIM_FILES),$(call exec-command,$(PROTOC_SLIM) $(PROTO_INCLUDES) --go_out=./$(PROTOBUF_TEMP_DIR) $(file)))

.PHONY: copy-otlp-protobuf-slim
copy-otlp-protobuf-slim:
	rm -rf $(OTLPSLIM_OUTPUT_DIR)/*/
	@rsync -a $(PROTOBUF_TEMP_DIR)/go.opentelemetry.io/proto/slim/otlp/ ./$(OTLPSLIM_OUTPUT_DIR)
	@git restore $$(git ls-files --deleted | grep -E 'go\.(mod|sum)$$')
	cd ./$(OTLPSLIM_OUTPUT_DIR)	&& go mod tidy

.PHONY: toolchain-check
toolchain-check:
	@toolchainRes=$$(for f in $(ALL_GO_MOD_DIRS); do \
	           awk '/^toolchain/ { found=1; next } END { if (found) print FILENAME }' $$f/go.mod; \
	done); \
	if [ -n "$${toolchainRes}" ]; then \
			echo "toolchain checking failed:"; echo "$${toolchainRes}"; \
			exit 1; \
	fi

.PHONY: clean-gen
clean-gen:
	rm -rf $(GEN_TEMP_DIR)

.PHONY: clean
clean:
	rm -rf $(GEN_TEMP_DIR)
	rm -rf $(OTLP_OUTPUT_DIR)/*/ $(OTLPSLIM_OUTPUT_DIR)/*/

.PHONY: go-mod-tidy
go-mod-tidy: clean-gen $(ALL_GO_MOD_DIRS:%=go-mod-tidy/%)
go-mod-tidy/%: DIR=$*
go-mod-tidy/%:
	@echo "$(GO) mod tidy in $(DIR)" \
		&& cd $(DIR) \
		&& $(GO) mod tidy

test: $(OTEL_GO_MOD_DIRS:%=test/%)
test/%: DIR=$*
test/%:
	@echo "$(GO) test -timeout $(TIMEOUT)s $(ARGS) $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) test -timeout $(TIMEOUT)s $(ARGS) ./...

.PHONY: check-clean-work-tree
check-clean-work-tree:
	@if ! git diff --quiet; then \
	  echo; \
	  echo 'Working tree is not clean, did you forget to run "make protobuf"?'; \
	  echo; \
	  git status; \
	  exit 1; \
	fi

# Releasing

.PHONY: submodule-version
submodule-version:
	@[ "$(VERSION)" ] || ( echo "VERSION unset: set to target opentelemetry-proto release tag"; exit 1 )
	@echo "upgrading opentelemetry-proto submodule to $(VERSION)"
	@cd opentelemetry-proto \
		&& git fetch \
		&& git checkout $(VERSION) \
		&& cd .. \
		&& git config -f .gitmodules submodule.opentelemetry-proto.branch $(VERSION)

.PHONY: sync
sync: submodule-version clean protobuf

.PHONY: verify-versions
verify-versions: | $(MULTIMOD)
	$(MULTIMOD) verify

COMMIT ?= "HEAD"
REMOTE ?= upstream
.PHONY: push-tags
push-tags: | $(MULTIMOD)
	$(MULTIMOD) verify
	for module in stable unstable; do \
		for tag in `$(MULTIMOD) tag -m $$module -c ${COMMIT} --print-tags | grep -v "Using"`; do \
			echo "pushing tag $$tag"; \
			git push ${REMOTE} $$tag; \
		done; \
	done
