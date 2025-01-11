package main

import (
	"testing"

	logs "go.opentelemetry.io/proto/otlp/collector/logs/v1"

	slimlogs "go.opentelemetry.io/proto/slim/otlp/collector/logs/v1"
)

func TestImport(t *testing.T) {
	_ = logs.ExportLogsServiceRequest{}
	_ = slimlogs.ExportLogsServiceRequest{}
}
