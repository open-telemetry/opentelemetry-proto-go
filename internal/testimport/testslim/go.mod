module testslim.local

go 1.23.4

require go.opentelemetry.io/proto/slim/otlp v1.4.0

require google.golang.org/protobuf v1.36.0 // indirect

replace go.opentelemetry.io/proto/slim/otlp => ../../../slim/otlp
