module go.opentelemetry.io/proto/slim/otlp/collector/profiles

go 1.23.0

require (
	go.opentelemetry.io/proto/slim/otlp/profiles v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.36.6
)

require go.opentelemetry.io/proto/slim/otlp v1.7.0 // indirect

replace go.opentelemetry.io/proto/slim/otlp => ../../

replace go.opentelemetry.io/proto/slim/otlp/profiles => ../../profiles/
