// Code generated by protoc-gen-goshim. DO NOT EDIT.

package v1

import slim "go.opentelemetry.io/proto/slim/otlp/logs/v1"

type LogsData = slim.LogsData
type ResourceLogs = slim.ResourceLogs
type ScopeLogs = slim.ScopeLogs
type LogRecord = slim.LogRecord
type SeverityNumber = slim.SeverityNumber

const (
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED = slim.SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
	SeverityNumber_SEVERITY_NUMBER_TRACE       = slim.SeverityNumber_SEVERITY_NUMBER_TRACE
	SeverityNumber_SEVERITY_NUMBER_TRACE2      = slim.SeverityNumber_SEVERITY_NUMBER_TRACE2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      = slim.SeverityNumber_SEVERITY_NUMBER_TRACE3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      = slim.SeverityNumber_SEVERITY_NUMBER_TRACE4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       = slim.SeverityNumber_SEVERITY_NUMBER_DEBUG
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      = slim.SeverityNumber_SEVERITY_NUMBER_DEBUG2
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      = slim.SeverityNumber_SEVERITY_NUMBER_DEBUG3
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      = slim.SeverityNumber_SEVERITY_NUMBER_DEBUG4
	SeverityNumber_SEVERITY_NUMBER_INFO        = slim.SeverityNumber_SEVERITY_NUMBER_INFO
	SeverityNumber_SEVERITY_NUMBER_INFO2       = slim.SeverityNumber_SEVERITY_NUMBER_INFO2
	SeverityNumber_SEVERITY_NUMBER_INFO3       = slim.SeverityNumber_SEVERITY_NUMBER_INFO3
	SeverityNumber_SEVERITY_NUMBER_INFO4       = slim.SeverityNumber_SEVERITY_NUMBER_INFO4
	SeverityNumber_SEVERITY_NUMBER_WARN        = slim.SeverityNumber_SEVERITY_NUMBER_WARN
	SeverityNumber_SEVERITY_NUMBER_WARN2       = slim.SeverityNumber_SEVERITY_NUMBER_WARN2
	SeverityNumber_SEVERITY_NUMBER_WARN3       = slim.SeverityNumber_SEVERITY_NUMBER_WARN3
	SeverityNumber_SEVERITY_NUMBER_WARN4       = slim.SeverityNumber_SEVERITY_NUMBER_WARN4
	SeverityNumber_SEVERITY_NUMBER_ERROR       = slim.SeverityNumber_SEVERITY_NUMBER_ERROR
	SeverityNumber_SEVERITY_NUMBER_ERROR2      = slim.SeverityNumber_SEVERITY_NUMBER_ERROR2
	SeverityNumber_SEVERITY_NUMBER_ERROR3      = slim.SeverityNumber_SEVERITY_NUMBER_ERROR3
	SeverityNumber_SEVERITY_NUMBER_ERROR4      = slim.SeverityNumber_SEVERITY_NUMBER_ERROR4
	SeverityNumber_SEVERITY_NUMBER_FATAL       = slim.SeverityNumber_SEVERITY_NUMBER_FATAL
	SeverityNumber_SEVERITY_NUMBER_FATAL2      = slim.SeverityNumber_SEVERITY_NUMBER_FATAL2
	SeverityNumber_SEVERITY_NUMBER_FATAL3      = slim.SeverityNumber_SEVERITY_NUMBER_FATAL3
	SeverityNumber_SEVERITY_NUMBER_FATAL4      = slim.SeverityNumber_SEVERITY_NUMBER_FATAL4
)

type LogRecordFlags = slim.LogRecordFlags

const (
	LogRecordFlags_LOG_RECORD_FLAGS_DO_NOT_USE       = slim.LogRecordFlags_LOG_RECORD_FLAGS_DO_NOT_USE
	LogRecordFlags_LOG_RECORD_FLAGS_TRACE_FLAGS_MASK = slim.LogRecordFlags_LOG_RECORD_FLAGS_TRACE_FLAGS_MASK
)
