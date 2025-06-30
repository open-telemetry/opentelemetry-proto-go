# opentelemetry-proto-go

[![Go Reference](https://pkg.go.dev/badge/go.opentelemetry.io/proto/otlp.svg)](https://pkg.go.dev/go.opentelemetry.io/proto/otlp)

Generated Go code for the OpenTelemetry protobuf data model.

## Versioning Policy

The auto-generated Go code follows the stability guarantees as defined in
[maturity
level](https://github.com/open-telemetry/opentelemetry-proto?tab=readme-ov-file#maturity-level).

Versioning of modules in this project will be idiomatic of a Go project using [Go modules](https://github.com/golang/go/wiki/Modules).
They will use [semantic import versioning](https://github.com/golang/go/wiki/Modules#semantic-import-versioning).
Meaning modules will comply with [semver 2.0](https://semver.org/spec/v2.0.0.html) with the following exception:

- Packages with a `development` suffix do not comply with [semver 2.0](https://semver.org/spec/v2.0.0.html).
  - Backwards incompatible changes may be introduced to these packages between minor versions.
  - These packages are intended to be temporary.
    They will be deprecated and removed when the protobuf definition stabilizes or is removed.
    If the protobuf definition stabilizes, the package will be replaced with a stable "non-development" package.
    If the protobuf definition is removed, the package will be removed without a replacement.

## Getting Started

Install the latest version in your project.

```sh
go get go.opentelemetry.io/proto/otlp@latest
```

Import the generated code directly in your project.

```go
import (
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)
```

### Compatibility

OpenTelemetry Proto Go ensures compatibility with the current supported
versions of
the [Go language](https://golang.org/doc/devel/release#policy):

> Each major Go release is supported until there are two newer major releases.
> For example, Go 1.5 was supported until the Go 1.7 release, and Go 1.6 was supported until the Go 1.8 release.

For versions of Go that are no longer supported upstream, opentelemetry-proto-go will
stop ensuring compatibility with these versions in the following manner:

- A minor release of opentelemetry-proto-go will be made to add support for the new
  supported release of Go.
- The following minor release of opentelemetry-proto-go will remove compatibility
  testing for the oldest (now archived upstream) version of Go. This, and
  future, releases of opentelemetry-proto-=go may include features only supported by
  the currently supported versions of Go.

This project is tested on the following systems.

| OS       | Go Version |
| -------- | ---------- |
| Ubuntu   | 1.24       |
| Ubuntu   | 1.23       |

While this project should work for other systems, no compatibility guarantees
are made for those systems currently.

## Maintainers

- [Mike Goldsmith](https://github.com/MikeGoldsmith), Honeycomb
- [OpenTelemetry Go Maintainers](https://github.com/open-telemetry/opentelemetry-go/blob/main/CONTRIBUTING.md#maintainers)

For more information about the maintainer role, see the [community repository](https://github.com/open-telemetry/community/blob/main/guides/contributor/membership.md#maintainer).

## Approvers

- [OpenTelemetry Go Approvers](https://github.com/open-telemetry/opentelemetry-go/blob/main/CONTRIBUTING.md#approvers)

For more information about the approver role, see the [community repository](https://github.com/open-telemetry/community/blob/main/guides/contributor/membership.md#approver).
