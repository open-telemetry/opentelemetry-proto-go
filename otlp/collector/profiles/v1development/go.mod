module go.opentelemetry.io/proto/otlp/collector/profiles/v1development

go 1.23.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	go.opentelemetry.io/proto/otlp/profiles/v1development v0.0.0-20250721084824-6f76ca90124d
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.10
)

require (
	go.opentelemetry.io/proto/otlp v1.8.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250825161204-c5933d9347a5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250825161204-c5933d9347a5 // indirect
)

replace go.opentelemetry.io/proto/otlp => ../../../

replace go.opentelemetry.io/proto/otlp/profiles/v1development => ../../../profiles/v1development
