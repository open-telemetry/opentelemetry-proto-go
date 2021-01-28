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

PROTOBUF_VERSION                := v1
OTEL_PROTO_SUBMODULE            := opentelemetry-proto
GEN_TEMP_DIR                    := gen
SUBMODULE_PROTO_FILES           := $(wildcard $(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/*/$(PROTOBUF_VERSION)/*.proto) $(wildcard $(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/*/$(PROTOBUF_VERSION)/*.proto)

ifeq ($(strip $(SUBMODULE_PROTO_FILES)),)
$(error Submodule at $(OTEL_PROTO_SUBMODULE) is not checked out, use "git submodule update --init")
endif

PROTOBUF_GEN_DIR   := opentelemetry-proto-gen
PROTOBUF_TEMP_DIR  := $(GEN_TEMP_DIR)/go
PROTO_SOURCE_DIR   := $(GEN_TEMP_DIR)/proto
SOURCE_PROTO_FILES := $(subst $(OTEL_PROTO_SUBMODULE),$(PROTO_SOURCE_DIR),$(SUBMODULE_PROTO_FILES))

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

OTEL_DOCKER_PROTOBUF ?= otel/build-protobuf:0.2.1
PROTOC := docker run --rm -u ${shell id -u} -v${PWD}:${PWD} -w${PWD} ${OTEL_DOCKER_PROTOBUF} --proto_path="$(PROTO_SOURCE_DIR)"
PROTO_INCLUDES := -I/usr/include/github.com/gogo/protobuf

.DEFAULT_GOAL := protobuf

# The sed expression for replacing the go_package option in proto
# file with a one that's valid for us.
SED_EXPR := 's,go_package = "github.com/open-telemetry/opentelemetry-proto/gen/go/,go_package = "go.opentelemetry.io/proto/,'

.PHONY: protobuf
protobuf: protobuf-source gen-protobuf copy-protobufs

.PHONY: protobuf-source
protobuf-source: $(SOURCE_PROTO_FILES)

# This copies proto files from submodule into $(PROTO_SOURCE_DIR),
# thus satisfying the $(SOURCE_PROTO_FILES) prerequisite. The copies
# have their package name replaced by go.opentelemetry.io/proto.
$(PROTO_SOURCE_DIR)/%.proto: $(OTEL_PROTO_SUBMODULE)/%.proto
	@ \
	mkdir -p $(@D); \
	sed -e $(SED_EXPR) "$<" >"$@.tmp"; \
	mv "$@.tmp" "$@"

.PHONY: gen-protobuf
gen-protobuf: $(SOURCE_PROTO_FILES)
	rm -rf ./$(PROTOBUF_TEMP_DIR)
	mkdir -p ./$(PROTOBUF_TEMP_DIR)
	$(foreach file,$(SOURCE_PROTO_FILES),$(call exec-command,$(PROTOC) $(PROTO_INCLUDES) --gogo_out=plugins=grpc:./$(PROTOBUF_TEMP_DIR) $(file)))
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/trace/v1/trace_service_http.yaml:./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/metrics/v1/metrics_service_http.yaml:./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$(OTEL_PROTO_SUBMODULE)/opentelemetry/proto/collector/logs/v1/logs_service_http.yaml:./$(PROTOBUF_TEMP_DIR) $(PROTO_SOURCE_DIR)/opentelemetry/proto/collector/logs/v1/logs_service.proto

.PHONY: copy-protobufs
copy-protobufs:
	@rsync -a $(PROTOBUF_TEMP_DIR)/go.opentelemetry.io/proto .

.PHONY: clean
clean:
	rm -rf $(GEN_TEMP_DIR)
