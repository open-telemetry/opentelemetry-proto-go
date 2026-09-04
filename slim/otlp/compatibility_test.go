// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlp_test

import (
	"bytes"
	"testing"

	collectlogs "go.opentelemetry.io/proto/slim/otlp/collector/logs/v1"
	collectmetrics "go.opentelemetry.io/proto/slim/otlp/collector/metrics/v1"
	collecttrace "go.opentelemetry.io/proto/slim/otlp/collector/trace/v1"
	common "go.opentelemetry.io/proto/slim/otlp/common/v1"
	logs "go.opentelemetry.io/proto/slim/otlp/logs/v1"
	metrics "go.opentelemetry.io/proto/slim/otlp/metrics/v1"
	resource "go.opentelemetry.io/proto/slim/otlp/resource/v1"
	trace "go.opentelemetry.io/proto/slim/otlp/trace/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// These module-local tests intentionally cover only the stable slim module.
// Cross-module compatibility, including development profiles and process context,
// is covered by ../../internal/slimotlpcompat/compatibility_test.go.

func TestStableFileDescriptorNamespace(t *testing.T) {
	tests := []struct {
		descriptor protoreflect.FileDescriptor
		path       string
		pkg        string
	}{
		{common.File_opentelemetry_proto_common_v1_common_proto, "opentelemetry/proto/slim/common/v1/common.proto", "opentelemetry.proto.slim.common.v1"},
		{resource.File_opentelemetry_proto_resource_v1_resource_proto, "opentelemetry/proto/slim/resource/v1/resource.proto", "opentelemetry.proto.slim.resource.v1"},
		{trace.File_opentelemetry_proto_trace_v1_trace_proto, "opentelemetry/proto/slim/trace/v1/trace.proto", "opentelemetry.proto.slim.trace.v1"},
		{metrics.File_opentelemetry_proto_metrics_v1_metrics_proto, "opentelemetry/proto/slim/metrics/v1/metrics.proto", "opentelemetry.proto.slim.metrics.v1"},
		{logs.File_opentelemetry_proto_logs_v1_logs_proto, "opentelemetry/proto/slim/logs/v1/logs.proto", "opentelemetry.proto.slim.logs.v1"},
		{collecttrace.File_opentelemetry_proto_collector_trace_v1_trace_service_proto, "opentelemetry/proto/slim/collector/trace/v1/trace_service.proto", "opentelemetry.proto.slim.collector.trace.v1"},
		{collectmetrics.File_opentelemetry_proto_collector_metrics_v1_metrics_service_proto, "opentelemetry/proto/slim/collector/metrics/v1/metrics_service.proto", "opentelemetry.proto.slim.collector.metrics.v1"},
		{collectlogs.File_opentelemetry_proto_collector_logs_v1_logs_service_proto, "opentelemetry/proto/slim/collector/logs/v1/logs_service.proto", "opentelemetry.proto.slim.collector.logs.v1"},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			if got := test.descriptor.Path(); got != test.path {
				t.Errorf("descriptor path = %q, want %q", got, test.path)
			}
			if got := string(test.descriptor.Package()); got != test.pkg {
				t.Errorf("protobuf package = %q, want %q", got, test.pkg)
			}
		})
	}
}

func TestStableCollectorRequestEncodings(t *testing.T) {
	const schemaURL = "schema"
	// Field 1 is the repeated resource message in every OTLP export request.
	// Field 3 in each resource message is schema_url. This is the canonical OTLP
	// protobuf encoding for one resource message containing only that value.
	wantBinary := []byte{0x0a, 0x08, 0x1a, 0x06, 's', 'c', 'h', 'e', 'm', 'a'}

	tests := []struct {
		name     string
		message  proto.Message
		wantJSON string
	}{
		{
			"traces",
			&collecttrace.ExportTraceServiceRequest{ResourceSpans: []*trace.ResourceSpans{{SchemaUrl: schemaURL}}},
			`{"resourceSpans":[{"schemaUrl":"schema"}]}`,
		},
		{
			"metrics",
			&collectmetrics.ExportMetricsServiceRequest{ResourceMetrics: []*metrics.ResourceMetrics{{SchemaUrl: schemaURL}}},
			`{"resourceMetrics":[{"schemaUrl":"schema"}]}`,
		},
		{
			"logs",
			&collectlogs.ExportLogsServiceRequest{ResourceLogs: []*logs.ResourceLogs{{SchemaUrl: schemaURL}}},
			`{"resourceLogs":[{"schemaUrl":"schema"}]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotBinary, err := proto.Marshal(test.message)
			if err != nil {
				t.Fatalf("marshal binary protobuf: %v", err)
			}
			if !bytes.Equal(gotBinary, wantBinary) {
				t.Errorf("binary protobuf = %x, want %x", gotBinary, wantBinary)
			}

			gotJSON, err := protojson.Marshal(test.message)
			if err != nil {
				t.Fatalf("marshal JSON protobuf: %v", err)
			}
			if string(gotJSON) != test.wantJSON {
				t.Errorf("JSON protobuf = %s, want %s", gotJSON, test.wantJSON)
			}
		})
	}
}
