module go.opentelemetry.io/proto/otlp/metrics

go 1.21

require (
	go.opentelemetry.io/proto/otlp/common v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/proto/otlp/resource v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.34.2
)

replace go.opentelemetry.io/proto/otlp/common => ../common

replace go.opentelemetry.io/proto/otlp/resource => ../resource
