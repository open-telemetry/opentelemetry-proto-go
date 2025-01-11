package main

import (
	"testing"

	logs "go.opentelemetry.io/proto/otlp/collector/logs/v1"
)

func TestImport(t *testing.T) {
	_ = logs.ExportLogsServiceRequest{}
}
