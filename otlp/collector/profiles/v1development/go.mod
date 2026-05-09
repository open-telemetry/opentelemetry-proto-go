module go.opentelemetry.io/proto/otlp/collector/profiles/v1development

go 1.25.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.29.0
	go.opentelemetry.io/proto/otlp/profiles/v1development v0.3.0
	google.golang.org/grpc v1.81.0
	google.golang.org/protobuf v1.36.11
)

require (
	go.opentelemetry.io/proto/otlp v1.10.0 // indirect
	golang.org/x/net v0.54.0 // indirect
	golang.org/x/sys v0.44.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260504160031-60b97b32f348 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260504160031-60b97b32f348 // indirect
)

replace go.opentelemetry.io/proto/otlp => ../../../

replace go.opentelemetry.io/proto/otlp/profiles/v1development => ../../../profiles/v1development
