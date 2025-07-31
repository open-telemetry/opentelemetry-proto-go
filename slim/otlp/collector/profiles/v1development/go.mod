module go.opentelemetry.io/proto/slim/otlp/collector/profiles/v1development

go 1.23.0

require (
	go.opentelemetry.io/proto/slim/otlp/profiles/v1development v0.0.0-20250721084824-6f76ca90124d
	google.golang.org/protobuf v1.36.6
)

require go.opentelemetry.io/proto/slim/otlp v1.7.1 // indirect

replace go.opentelemetry.io/proto/slim/otlp => ../../../

replace go.opentelemetry.io/proto/slim/otlp/profiles/v1development => ../../../profiles/v1development/
