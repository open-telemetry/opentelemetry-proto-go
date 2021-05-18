# opentelemetry-proto-go

Generated code for OpenTelemetry protobuf data model

## Usage

You can import the generated code directly in your project

```go
import tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
```

## Generating new versions

When
[opentelemetry-proto](https://github.com/open-telemetry/opentelemetry-proto)
release a new version we can update submodule, and then regenerate the
protobufs.

```sh
git config -f .gitmodules submodule.opentelemetry-proto.branch v0.7.0
git submodule update --init
make clean gen-otlp-protobuf copy-otlp-protobu
```
