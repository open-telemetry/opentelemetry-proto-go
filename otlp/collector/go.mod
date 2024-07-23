module go.opentelemetry.io/proto/otlp/collector

go 1.21.12

replace go.opentelemetry.io/proto/otlp/common => ../common

replace go.opentelemetry.io/proto/otlp/logs => ../logs

replace go.opentelemetry.io/proto/otlp/metrics => ../metrics

replace go.opentelemetry.io/proto/otlp/profiles => ../profiles

replace go.opentelemetry.io/proto/otlp/resource => ../resource

replace go.opentelemetry.io/proto/otlp/trace => ../trace

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	go.opentelemetry.io/proto/otlp/logs v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/proto/otlp/metrics v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/proto/otlp/profiles v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/proto/otlp/trace v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	go.opentelemetry.io/proto/otlp/common v0.0.0-00010101000000-000000000000 // indirect
	go.opentelemetry.io/proto/otlp/resource v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)
