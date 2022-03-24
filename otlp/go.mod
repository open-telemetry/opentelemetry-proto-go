module go.opentelemetry.io/proto/otlp

go 1.14

replace github.com/grpc-ecosystem/grpc-gateway/v2 => github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
