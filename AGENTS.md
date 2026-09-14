# Agent Guide for opentelemetry-proto-go

## Canonical and slim protobuf compatibility

- Keep `internal/slimotlpcompat/compatibility_test.go` exhaustive. Every file-level
  `message` declared under `opentelemetry-proto/opentelemetry/proto/` must have
  canonical-versus-slim compatibility coverage.
- Keep cross-module compatibility coverage, including development profiles and
  process context, in `internal/slimotlpcompat`. Do not add even test-only
  requirements from the stable `go.opentelemetry.io/proto/slim/otlp` module to
  development modules.
- Populate each message with representative non-default data and verify that the
  canonical and slim forms have identical protobuf wire encodings and equivalent
  ProtoJSON encodings.
- In that suite, keep one explicit non-empty case for each signal's collector
  export request (traces, metrics, logs, and profiles) and for the standalone
  process-context payload.
- Cover stable and development schemas. When adding a proto file or a file-level
  message, update the compatibility file pairs and expected message count so a
  schema addition cannot silently bypass compatibility testing.
