module go.opentelemetry.io/proto/otlp/collector/profiles/v1development

go 1.23.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	go.opentelemetry.io/proto/otlp/profiles/v1development v0.0.0-20250721084824-6f76ca90124d
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.9
)

require (
	go.opentelemetry.io/proto/otlp v1.8.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/api 9219d122eba9 // indirect
	google.golang.org/genproto/googleapis/rpc 9219d122eba9 // indirect
)

replace go.opentelemetry.io/proto/otlp => ../../../

replace go.opentelemetry.io/proto/otlp/profiles/v1development => ../../../profiles/v1development
