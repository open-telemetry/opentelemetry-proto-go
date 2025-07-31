module go.opentelemetry.io/proto/otlp/collector/profiles/v1development

go 1.23.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1
	go.opentelemetry.io/proto/otlp/profiles/v1development v0.0.0-20250721084824-6f76ca90124d
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
)

require (
	go.opentelemetry.io/proto/otlp v1.7.1 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250728155136-f173205681a0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250728155136-f173205681a0 // indirect
)

replace go.opentelemetry.io/proto/otlp => ../../../

replace go.opentelemetry.io/proto/otlp/profiles/v1development => ../../../profiles/v1development
