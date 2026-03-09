module go.opentelemetry.io/proto/slim/otlp/collector/profiles/v1development

go 1.24.0

require (
	go.opentelemetry.io/proto/slim/otlp/profiles/v1development v0.3.0
	google.golang.org/protobuf v1.36.11
)

require go.opentelemetry.io/proto/slim/otlp v1.10.0 // indirect

replace go.opentelemetry.io/proto/slim/otlp => ../../../

replace go.opentelemetry.io/proto/slim/otlp/profiles/v1development => ../../../profiles/v1development/
