module go.opentelemetry.io/proto/light/otlp/collector

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	go.opentelemetry.io/proto/light/otlp v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.61.1
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
)

replace go.opentelemetry.io/proto/light/otlp => ../
