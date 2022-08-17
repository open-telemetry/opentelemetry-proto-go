package otlp

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	otlpmetrics "go.opentelemetry.io/proto/otlp/metrics/v1"
)

func TestDecodeAggregationTemporality(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    otlpmetrics.AggregationTemporality
	}{
		{
			name: "string",
			jsonStr: fmt.Sprintf("{\"aggregation_temporality\":\"%s\"}",
				otlpmetrics.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE.String()),
			want: otlpmetrics.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE,
		},
		{
			name:    "unknown string",
			jsonStr: fmt.Sprintf("{\"aggregation_temporality\":\"%s\"}", "foo"),
			want:    otlpmetrics.AggregationTemporality_AGGREGATION_TEMPORALITY_UNSPECIFIED,
		},
		{
			name: "int",
			jsonStr: fmt.Sprintf("{\"aggregation_temporality\": %d}",
				otlpmetrics.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE),
			want: otlpmetrics.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE,
		},
		{
			name:    "unknown int",
			jsonStr: fmt.Sprintf("{\"aggregation_temporality\": %d}", 5),
			want:    otlpmetrics.AggregationTemporality(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := &otlpmetrics.Sum{}
			if err := protojson.Unmarshal([]byte(tt.jsonStr), sum); err != nil {
				t.Errorf("no error expected, got: %v", err)
			}

			if got := sum.AggregationTemporality; got != tt.want {
				t.Errorf("readSpanKind() = %v, want %v", got, tt.want)
			}
		})
	}
}
