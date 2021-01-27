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
# Prereqs: wget (for downloading the zip file with protoc binary),
# unzip (for unpacking the archive), rsync (for copying back the
# generated files).

PROTOC_VERSION := 3.14.0

TOOLS_DIR                       := $(abspath ./.tools)
TOOLS_MOD_DIR                   := ./internal/tools
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

.DEFAULT_GOAL := protobuf

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Linux)

PROTOC_OS := linux
PROTOC_ARCH := $(UNAME_M)

else ifeq ($(UNAME_S),Darwin)

PROTOC_OS := osx
PROTOC_ARCH := x86_64

endif

PROTOC_ZIP_URL := https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip

$(TOOLS_DIR)/PROTOC_$(PROTOC_VERSION):
	@rm -f "$(TOOLS_DIR)"/PROTOC_* && \
	touch "$@"

# Depend on a versioned file (like PROTOC_3.14.0), so when version
# gets bumped, we will depend on a nonexistent file and thus download
# a newer version.
$(TOOLS_DIR)/protoc/bin/protoc: $(TOOLS_DIR)/PROTOC_$(PROTOC_VERSION)
	echo "Fetching protoc $(PROTOC_VERSION)" && \
	rm -rf $(TOOLS_DIR)/protoc && \
	wget -O $(TOOLS_DIR)/protoc.zip $(PROTOC_ZIP_URL) && \
	unzip $(TOOLS_DIR)/protoc.zip -d $(TOOLS_DIR)/protoc-tmp && \
	rm $(TOOLS_DIR)/protoc.zip && \
	touch $(TOOLS_DIR)/protoc-tmp/bin/protoc && \
	mv $(TOOLS_DIR)/protoc-tmp $(TOOLS_DIR)/protoc

$(TOOLS_DIR)/protoc-gen-gogofast: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/protoc-gen-gogofast github.com/gogo/protobuf/protoc-gen-gogofast && \
	go mod tidy

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
gen-protobuf: $(SOURCE_PROTO_FILES) $(TOOLS_DIR)/protoc-gen-gogofast $(TOOLS_DIR)/protoc/bin/protoc
	@ \
	mkdir -p "$(PROTOBUF_TEMP_DIR)"; \
	set -e; for f in $^; do \
	  if [[ "$${f}" == $(TOOLS_DIR)/* ]]; then continue; fi; \
	  echo "protoc $${f#"$(PROTO_SOURCE_DIR)/"}"; \
	  PATH="$(TOOLS_DIR):$${PATH}" $(TOOLS_DIR)/protoc/bin/protoc --proto_path="$(PROTO_SOURCE_DIR)" --gogofast_out="plugins=grpc:$(PROTOBUF_TEMP_DIR)" "$${f}"; \
	done

.PHONY: copy-protobufs
copy-protobufs:
	@rsync -a $(PROTOBUF_TEMP_DIR)/go.opentelemetry.io/proto .

.PHONY: clean
clean:
	rm -rf $(GEN_TEMP_DIR)
