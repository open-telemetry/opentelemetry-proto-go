module go.opentelemetry.io/proto/internal/slimotlpcompat

go 1.25.0

require (
	go.opentelemetry.io/proto/otlp v1.11.0
	go.opentelemetry.io/proto/otlp/collector/profiles/v1development v0.4.0
	go.opentelemetry.io/proto/otlp/processcontext/v1development v0.4.0
	go.opentelemetry.io/proto/otlp/profiles/v1development v0.4.0
	go.opentelemetry.io/proto/slim/otlp v1.11.0
	go.opentelemetry.io/proto/slim/otlp/collector/profiles/v1development v0.4.0
	go.opentelemetry.io/proto/slim/otlp/processcontext/v1development v0.4.0
	go.opentelemetry.io/proto/slim/otlp/profiles/v1development v0.4.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260831171406-18b4a7587f8a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260831171406-18b4a7587f8a // indirect
	google.golang.org/grpc v1.83.2 // indirect
)

replace go.opentelemetry.io/proto/otlp => ../../otlp

replace go.opentelemetry.io/proto/otlp/collector/profiles/v1development => ../../otlp/collector/profiles/v1development

replace go.opentelemetry.io/proto/otlp/processcontext/v1development => ../../otlp/processcontext/v1development

replace go.opentelemetry.io/proto/otlp/profiles/v1development => ../../otlp/profiles/v1development

replace go.opentelemetry.io/proto/slim/otlp => ../../slim/otlp

replace go.opentelemetry.io/proto/slim/otlp/collector/profiles/v1development => ../../slim/otlp/collector/profiles/v1development

replace go.opentelemetry.io/proto/slim/otlp/processcontext/v1development => ../../slim/otlp/processcontext/v1development

replace go.opentelemetry.io/proto/slim/otlp/profiles/v1development => ../../slim/otlp/profiles/v1development
