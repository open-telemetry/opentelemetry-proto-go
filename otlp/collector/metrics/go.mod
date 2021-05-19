module go.opentelemetry.io/proto/otlp/collector/metrics

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	go.opentelemetry.io/proto/otlp v0.9.1
	go.opentelemetry.io/proto/otlp/metrics v0.9.1
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.26.0
)

replace go.opentelemetry.io/proto/otlp => ../../

replace go.opentelemetry.io/proto/otlp/metrics => ../../metrics
