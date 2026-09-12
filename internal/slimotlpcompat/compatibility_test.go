// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package slimotlpcompat

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	otlpcollectlogs "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	otlpcollectmetrics "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	otlpcollectprofiles "go.opentelemetry.io/proto/otlp/collector/profiles/v1development"
	otlpcollecttrace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	otlpcommon "go.opentelemetry.io/proto/otlp/common/v1"
	otlplogs "go.opentelemetry.io/proto/otlp/logs/v1"
	otlpmetrics "go.opentelemetry.io/proto/otlp/metrics/v1"
	otlpprocesscontext "go.opentelemetry.io/proto/otlp/processcontext/v1development"
	otlpprofiles "go.opentelemetry.io/proto/otlp/profiles/v1development"
	otlpresource "go.opentelemetry.io/proto/otlp/resource/v1"
	otlptrace "go.opentelemetry.io/proto/otlp/trace/v1"
	slimcollectlogs "go.opentelemetry.io/proto/slim/otlp/collector/logs/v1"
	slimcollectmetrics "go.opentelemetry.io/proto/slim/otlp/collector/metrics/v1"
	slimcollectprofiles "go.opentelemetry.io/proto/slim/otlp/collector/profiles/v1development"
	slimcollecttrace "go.opentelemetry.io/proto/slim/otlp/collector/trace/v1"
	slimcommon "go.opentelemetry.io/proto/slim/otlp/common/v1"
	slimlogs "go.opentelemetry.io/proto/slim/otlp/logs/v1"
	slimmetrics "go.opentelemetry.io/proto/slim/otlp/metrics/v1"
	slimprocesscontext "go.opentelemetry.io/proto/slim/otlp/processcontext/v1development"
	slimprofiles "go.opentelemetry.io/proto/slim/otlp/profiles/v1development"
	slimresource "go.opentelemetry.io/proto/slim/otlp/resource/v1"
	slimtrace "go.opentelemetry.io/proto/slim/otlp/trace/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	canonicalPathPrefix = "opentelemetry/proto/"
	slimPathPrefix      = "opentelemetry/proto/slim/"
	canonicalNamePrefix = "opentelemetry.proto."
	slimNamePrefix      = "opentelemetry.proto.slim."
	canonicalGoPrefix   = "go.opentelemetry.io/proto/otlp"
	slimGoPrefix        = "go.opentelemetry.io/proto/slim/otlp"
)

type filePair struct {
	canonical protoreflect.FileDescriptor
	slim      protoreflect.FileDescriptor
}

var allOTLPFiles = []filePair{
	{otlpcommon.File_opentelemetry_proto_common_v1_common_proto, slimcommon.File_opentelemetry_proto_common_v1_common_proto},
	{otlpresource.File_opentelemetry_proto_resource_v1_resource_proto, slimresource.File_opentelemetry_proto_resource_v1_resource_proto},
	{otlptrace.File_opentelemetry_proto_trace_v1_trace_proto, slimtrace.File_opentelemetry_proto_trace_v1_trace_proto},
	{otlpmetrics.File_opentelemetry_proto_metrics_v1_metrics_proto, slimmetrics.File_opentelemetry_proto_metrics_v1_metrics_proto},
	{otlplogs.File_opentelemetry_proto_logs_v1_logs_proto, slimlogs.File_opentelemetry_proto_logs_v1_logs_proto},
	{otlpcollecttrace.File_opentelemetry_proto_collector_trace_v1_trace_service_proto, slimcollecttrace.File_opentelemetry_proto_collector_trace_v1_trace_service_proto},
	{otlpcollectmetrics.File_opentelemetry_proto_collector_metrics_v1_metrics_service_proto, slimcollectmetrics.File_opentelemetry_proto_collector_metrics_v1_metrics_service_proto},
	{otlpcollectlogs.File_opentelemetry_proto_collector_logs_v1_logs_service_proto, slimcollectlogs.File_opentelemetry_proto_collector_logs_v1_logs_service_proto},
	{otlpprocesscontext.File_opentelemetry_proto_processcontext_v1development_process_context_proto, slimprocesscontext.File_opentelemetry_proto_processcontext_v1development_process_context_proto},
	{otlpprofiles.File_opentelemetry_proto_profiles_v1development_profiles_proto, slimprofiles.File_opentelemetry_proto_profiles_v1development_profiles_proto},
	{otlpcollectprofiles.File_opentelemetry_proto_collector_profiles_v1development_profiles_service_proto, slimcollectprofiles.File_opentelemetry_proto_collector_profiles_v1development_profiles_service_proto},
}

func TestCanonicalAndSlimDescriptorsCoexist(t *testing.T) {
	if got, want := len(allOTLPFiles), 11; got != want {
		t.Fatalf("tested file descriptors = %d, want %d", got, want)
	}

	for _, files := range allOTLPFiles {
		canonicalPath := files.canonical.Path()
		t.Run(canonicalPath, func(t *testing.T) {
			wantSlimPath := strings.Replace(canonicalPath, canonicalPathPrefix, slimPathPrefix, 1)
			if got := files.slim.Path(); got != wantSlimPath {
				t.Errorf("slim descriptor path = %q, want %q", got, wantSlimPath)
			}

			canonicalPackage := string(files.canonical.Package())
			wantSlimPackage := strings.Replace(canonicalPackage, canonicalNamePrefix, slimNamePrefix, 1)
			if got := string(files.slim.Package()); got != wantSlimPackage {
				t.Errorf("slim protobuf package = %q, want %q", got, wantSlimPackage)
			}

			assertRegisteredFile(t, files.canonical)
			assertRegisteredFile(t, files.slim)
			assertSlimFileNamespace(t, files.slim)
			assertCompatibleFileShape(t, files.canonical, files.slim)
		})
	}
}

func TestAllProtoFilesHaveCompatibilityPairs(t *testing.T) {
	covered := make(map[string]struct{}, len(allOTLPFiles))
	for _, files := range allOTLPFiles {
		path := files.canonical.Path()
		if _, exists := covered[path]; exists {
			t.Errorf("duplicate compatibility file pair for %q", path)
		}
		covered[path] = struct{}{}
	}

	const sourceRoot = "../../opentelemetry-proto/opentelemetry/proto"
	sources := make(map[string]struct{})
	err := filepath.WalkDir(sourceRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(path) != ".proto" {
			return nil
		}
		relative, err := filepath.Rel(sourceRoot, path)
		if err != nil {
			return err
		}
		descriptorPath := filepath.ToSlash(filepath.Join(canonicalPathPrefix, relative))
		sources[descriptorPath] = struct{}{}
		return nil
	})
	if err != nil {
		t.Fatalf("discover source proto files: %v", err)
	}

	for path := range sources {
		if _, exists := covered[path]; !exists {
			t.Errorf("source proto %q has no compatibility file pair", path)
		}
	}
	for path := range covered {
		if _, exists := sources[path]; !exists {
			t.Errorf("compatibility file pair %q has no source proto", path)
		}
	}
}

func TestRepresentativeSignalAndProcessContextMessagesHaveCompatibleEncoding(t *testing.T) {
	const schemaURL = "https://opentelemetry.io/schemas/1.40.0"
	tests := []struct {
		name      string
		canonical proto.Message
		slim      proto.Message
	}{
		{
			"traces",
			&otlpcollecttrace.ExportTraceServiceRequest{ResourceSpans: []*otlptrace.ResourceSpans{{SchemaUrl: schemaURL}}},
			&slimcollecttrace.ExportTraceServiceRequest{ResourceSpans: []*slimtrace.ResourceSpans{{SchemaUrl: schemaURL}}},
		},
		{
			"metrics",
			&otlpcollectmetrics.ExportMetricsServiceRequest{ResourceMetrics: []*otlpmetrics.ResourceMetrics{{SchemaUrl: schemaURL}}},
			&slimcollectmetrics.ExportMetricsServiceRequest{ResourceMetrics: []*slimmetrics.ResourceMetrics{{SchemaUrl: schemaURL}}},
		},
		{
			"logs",
			&otlpcollectlogs.ExportLogsServiceRequest{ResourceLogs: []*otlplogs.ResourceLogs{{SchemaUrl: schemaURL}}},
			&slimcollectlogs.ExportLogsServiceRequest{ResourceLogs: []*slimlogs.ResourceLogs{{SchemaUrl: schemaURL}}},
		},
		{
			"profiles",
			&otlpcollectprofiles.ExportProfilesServiceRequest{
				ResourceProfiles: []*otlpprofiles.ResourceProfiles{{SchemaUrl: schemaURL}},
				Dictionary:       &otlpprofiles.ProfilesDictionary{StringTable: []string{"", "cpu"}},
			},
			&slimcollectprofiles.ExportProfilesServiceRequest{
				ResourceProfiles: []*slimprofiles.ResourceProfiles{{SchemaUrl: schemaURL}},
				Dictionary:       &slimprofiles.ProfilesDictionary{StringTable: []string{"", "cpu"}},
			},
		},
		{
			"process context",
			&otlpprocesscontext.ProcessContext{
				Resource: &otlpresource.Resource{DroppedAttributesCount: 7},
				Attributes: []*otlpcommon.KeyValue{{
					Key: "key",
					Value: &otlpcommon.AnyValue{
						Value: &otlpcommon.AnyValue_StringValue{StringValue: "value"},
					},
				}},
			},
			&slimprocesscontext.ProcessContext{
				Resource: &slimresource.Resource{DroppedAttributesCount: 7},
				Attributes: []*slimcommon.KeyValue{{
					Key: "key",
					Value: &slimcommon.AnyValue{
						Value: &slimcommon.AnyValue_StringValue{StringValue: "value"},
					},
				}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertCompatibleEncoding(t, test.canonical, test.slim)
		})
	}
}

func TestEveryTopLevelMessageHasCompatibleEncoding(t *testing.T) {
	const topLevelMessageCount = 57
	tested := 0

	// Generated fixtures exercise protobuf encoding, not signal-specific semantic
	// constraints. New special message types may need tailored fixture values.
	for _, files := range allOTLPFiles {
		canonicalMessages := files.canonical.Messages()
		slimMessages := files.slim.Messages()
		if got, want := slimMessages.Len(), canonicalMessages.Len(); got != want {
			t.Errorf("top-level message count for %s = %d, want %d", files.slim.Path(), got, want)
		}

		for i := 0; i < canonicalMessages.Len(); i++ {
			canonicalDescriptor := canonicalMessages.Get(i)
			slimDescriptor := slimMessages.ByName(canonicalDescriptor.Name())
			if slimDescriptor == nil {
				t.Errorf("slim descriptor for %s not found", canonicalDescriptor.FullName())
				continue
			}

			tested++
			t.Run(string(canonicalDescriptor.FullName()), func(t *testing.T) {
				canonical := newRepresentativeMessage(canonicalDescriptor)
				slim := newRepresentativeMessage(slimDescriptor)
				assertCompatibleEncoding(t, canonical, slim)
			})
		}
	}

	if tested != topLevelMessageCount {
		t.Errorf("tested top-level messages = %d, want %d", tested, topLevelMessageCount)
	}
}

func assertCompatibleFileShape(t *testing.T, canonical, slim protoreflect.FileDescriptor) {
	t.Helper()

	canonicalProto := protodesc.ToFileDescriptorProto(canonical)
	slimProto := protodesc.ToFileDescriptorProto(slim)
	wantSlimGoPackage := strings.Replace(
		canonicalProto.GetOptions().GetGoPackage(),
		canonicalGoPrefix,
		slimGoPrefix,
		1,
	)
	if got := slimProto.GetOptions().GetGoPackage(); got != wantSlimGoPackage {
		t.Errorf("slim Go package = %q, want %q", got, wantSlimGoPackage)
	}

	canonicalizeSlimFileDescriptor(slimProto)
	if !proto.Equal(slimProto, canonicalProto) {
		t.Errorf(
			"file descriptor mismatch for %s:\ncanonical:\n%s\nslim:\n%s",
			canonical.Path(),
			prototext.Format(canonicalProto),
			prototext.Format(slimProto),
		)
	}
}

func canonicalizeSlimFileDescriptor(file *descriptorpb.FileDescriptorProto) {
	name := strings.Replace(file.GetName(), slimPathPrefix, canonicalPathPrefix, 1)
	file.Name = &name
	packageName := strings.Replace(file.GetPackage(), slimNamePrefix, canonicalNamePrefix, 1)
	file.Package = &packageName
	for i, dependency := range file.Dependency {
		file.Dependency[i] = strings.Replace(dependency, slimPathPrefix, canonicalPathPrefix, 1)
	}
	if options := file.Options; options != nil && options.GoPackage != nil {
		goPackage := strings.Replace(options.GetGoPackage(), slimGoPrefix, canonicalGoPrefix, 1)
		options.GoPackage = &goPackage
	}
	for _, message := range file.MessageType {
		canonicalizeSlimTypeNames(message)
	}
	for _, extension := range file.Extension {
		canonicalizeSlimField(extension)
	}
	for _, service := range file.Service {
		for _, method := range service.Method {
			inputType := canonicalizeSlimTypeName(method.GetInputType())
			method.InputType = &inputType
			outputType := canonicalizeSlimTypeName(method.GetOutputType())
			method.OutputType = &outputType
		}
	}
}

func canonicalizeSlimTypeNames(message *descriptorpb.DescriptorProto) {
	for _, field := range message.Field {
		canonicalizeSlimField(field)
	}
	for _, field := range message.Extension {
		canonicalizeSlimField(field)
	}
	for _, nested := range message.NestedType {
		canonicalizeSlimTypeNames(nested)
	}
}

func canonicalizeSlimField(field *descriptorpb.FieldDescriptorProto) {
	if field.TypeName != nil {
		name := canonicalizeSlimTypeName(field.GetTypeName())
		field.TypeName = &name
	}
	if field.Extendee != nil {
		name := canonicalizeSlimTypeName(field.GetExtendee())
		field.Extendee = &name
	}
}

func canonicalizeSlimTypeName(name string) string {
	return strings.Replace(name, "."+slimNamePrefix, "."+canonicalNamePrefix, 1)
}

func newRepresentativeMessage(descriptor protoreflect.MessageDescriptor) proto.Message {
	message := dynamicpb.NewMessage(descriptor)
	populateRepresentativeFields(message, 0)
	return message
}

func populateRepresentativeFields(message protoreflect.Message, depth int) {
	const maxDepth = 4
	if depth >= maxDepth {
		return
	}

	populatedOneofs := make(map[protoreflect.FullName]struct{})
	fields := message.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if oneof := field.ContainingOneof(); oneof != nil {
			if _, exists := populatedOneofs[oneof.FullName()]; exists {
				continue
			}
			populatedOneofs[oneof.FullName()] = struct{}{}
		}

		switch {
		case field.IsMap():
			values := message.Mutable(field).Map()
			key := representativeValue(field.MapKey(), depth+1).MapKey()
			values.Set(key, representativeValue(field.MapValue(), depth+1))
		case field.IsList():
			values := message.Mutable(field).List()
			values.Append(representativeValue(field, depth+1))
		default:
			message.Set(field, representativeValue(field, depth+1))
		}
	}
}

func representativeValue(field protoreflect.FieldDescriptor, depth int) protoreflect.Value {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(true)
	case protoreflect.EnumKind:
		values := field.Enum().Values()
		index := min(1, values.Len()-1)
		return protoreflect.ValueOfEnum(values.Get(index).Number())
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(-42)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(-42)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(42)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(42)
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(1.5)
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(1.5)
	case protoreflect.StringKind:
		return protoreflect.ValueOfString("value")
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes([]byte{1, 2, 3})
	case protoreflect.MessageKind, protoreflect.GroupKind:
		message := dynamicpb.NewMessage(field.Message())
		populateRepresentativeFields(message, depth)
		return protoreflect.ValueOfMessage(message)
	default:
		panic("unsupported protobuf field kind: " + field.Kind().String())
	}
}

func assertCompatibleEncoding(t *testing.T, canonical, slim proto.Message) {
	t.Helper()

	canonicalBinary, err := proto.MarshalOptions{Deterministic: true}.Marshal(canonical)
	if err != nil {
		t.Fatalf("marshal canonical binary protobuf: %v", err)
	}
	slimBinary, err := proto.MarshalOptions{Deterministic: true}.Marshal(slim)
	if err != nil {
		t.Fatalf("marshal slim binary protobuf: %v", err)
	}
	if !bytes.Equal(slimBinary, canonicalBinary) {
		t.Errorf("slim binary protobuf = %x, canonical = %x", slimBinary, canonicalBinary)
	}

	canonicalJSON, err := protojson.Marshal(canonical)
	if err != nil {
		t.Fatalf("marshal canonical JSON protobuf: %v", err)
	}
	slimJSON, err := protojson.Marshal(slim)
	if err != nil {
		t.Fatalf("marshal slim JSON protobuf: %v", err)
	}
	var canonicalValue, slimValue any
	if err := json.Unmarshal(canonicalJSON, &canonicalValue); err != nil {
		t.Fatalf("unmarshal canonical JSON protobuf: %v", err)
	}
	if err := json.Unmarshal(slimJSON, &slimValue); err != nil {
		t.Fatalf("unmarshal slim JSON protobuf: %v", err)
	}
	if !reflect.DeepEqual(slimValue, canonicalValue) {
		t.Errorf("slim JSON protobuf = %s, canonical = %s", slimJSON, canonicalJSON)
	}
}

func TestAnyTypeURLsUseDistinctNamespaces(t *testing.T) {
	canonicalAny, err := anypb.New(&otlpcommon.AnyValue{})
	if err != nil {
		t.Fatalf("pack canonical message: %v", err)
	}
	slimAny, err := anypb.New(&slimcommon.AnyValue{})
	if err != nil {
		t.Fatalf("pack slim message: %v", err)
	}

	if got, want := canonicalAny.TypeUrl, "type.googleapis.com/opentelemetry.proto.common.v1.AnyValue"; got != want {
		t.Errorf("canonical type URL = %q, want %q", got, want)
	}
	if got, want := slimAny.TypeUrl, "type.googleapis.com/opentelemetry.proto.slim.common.v1.AnyValue"; got != want {
		t.Errorf("slim type URL = %q, want %q", got, want)
	}

	resolvedCanonical, err := anypb.UnmarshalNew(canonicalAny, proto.UnmarshalOptions{})
	if err != nil {
		t.Fatalf("resolve canonical type URL: %v", err)
	}
	if _, ok := resolvedCanonical.(*otlpcommon.AnyValue); !ok {
		t.Errorf("canonical type URL resolved to %T", resolvedCanonical)
	}
	resolvedSlim, err := anypb.UnmarshalNew(slimAny, proto.UnmarshalOptions{})
	if err != nil {
		t.Fatalf("resolve slim type URL: %v", err)
	}
	if _, ok := resolvedSlim.(*slimcommon.AnyValue); !ok {
		t.Errorf("slim type URL resolved to %T", resolvedSlim)
	}
}

func TestCoreSchemasDoNotEmbedGoogleProtobufAny(t *testing.T) {
	for _, files := range allOTLPFiles {
		visitMessages(files.slim.Messages(), func(message protoreflect.MessageDescriptor) {
			fields := message.Fields()
			for i := 0; i < fields.Len(); i++ {
				field := fields.Get(i)
				if field.Message() != nil && field.Message().FullName() == "google.protobuf.Any" {
					t.Errorf("%s embeds google.protobuf.Any", field.FullName())
				}
			}
		})
	}
}

func assertRegisteredFile(t *testing.T, descriptor protoreflect.FileDescriptor) {
	t.Helper()
	got, err := protoregistry.GlobalFiles.FindFileByPath(descriptor.Path())
	if err != nil {
		t.Errorf("descriptor path %q is not registered: %v", descriptor.Path(), err)
		return
	}
	if got != descriptor {
		t.Errorf("descriptor path %q resolved to a different file", descriptor.Path())
	}
}

func assertSlimFileNamespace(t *testing.T, file protoreflect.FileDescriptor) {
	t.Helper()

	imports := file.Imports()
	for i := 0; i < imports.Len(); i++ {
		path := imports.Get(i).Path()
		if strings.HasPrefix(path, canonicalPathPrefix) && !strings.HasPrefix(path, slimPathPrefix) {
			t.Errorf("%s imports canonical OTLP descriptor %q", file.Path(), path)
		}
	}

	checkDescriptor := func(descriptor protoreflect.Descriptor) {
		t.Helper()
		name := string(descriptor.FullName())
		if strings.HasPrefix(name, canonicalNamePrefix) && !strings.HasPrefix(name, slimNamePrefix) {
			t.Errorf("descriptor %q does not use the slim namespace", name)
		}
		if _, err := protoregistry.GlobalFiles.FindDescriptorByName(descriptor.FullName()); err != nil {
			t.Errorf("descriptor %q is not registered: %v", name, err)
		}
	}

	visitEnums(file.Enums(), checkDescriptor)
	visitMessages(file.Messages(), func(message protoreflect.MessageDescriptor) {
		checkDescriptor(message)
		visitEnums(message.Enums(), checkDescriptor)

		fields := message.Fields()
		for i := 0; i < fields.Len(); i++ {
			field := fields.Get(i)
			checkDescriptor(field)
			assertSlimReference(t, field.Message())
			assertSlimReference(t, field.Enum())
		}

		oneofs := message.Oneofs()
		for i := 0; i < oneofs.Len(); i++ {
			checkDescriptor(oneofs.Get(i))
		}
	})

	extensions := file.Extensions()
	for i := 0; i < extensions.Len(); i++ {
		extension := extensions.Get(i)
		checkDescriptor(extension)
		assertSlimReference(t, extension.ContainingMessage())
		assertSlimReference(t, extension.Message())
		assertSlimReference(t, extension.Enum())
	}

	services := file.Services()
	for i := 0; i < services.Len(); i++ {
		service := services.Get(i)
		checkDescriptor(service)
		methods := service.Methods()
		for j := 0; j < methods.Len(); j++ {
			method := methods.Get(j)
			checkDescriptor(method)
			assertSlimReference(t, method.Input())
			assertSlimReference(t, method.Output())
		}
	}
}

func assertSlimReference(t *testing.T, descriptor protoreflect.Descriptor) {
	t.Helper()
	if descriptor == nil {
		return
	}
	name := string(descriptor.FullName())
	if strings.HasPrefix(name, canonicalNamePrefix) && !strings.HasPrefix(name, slimNamePrefix) {
		t.Errorf("internal reference %q does not use the slim namespace", name)
	}
}

func visitMessages(messages protoreflect.MessageDescriptors, visit func(protoreflect.MessageDescriptor)) {
	for i := 0; i < messages.Len(); i++ {
		message := messages.Get(i)
		visit(message)
		visitMessages(message.Messages(), visit)
	}
}

func visitEnums(enums protoreflect.EnumDescriptors, visit func(protoreflect.Descriptor)) {
	for i := 0; i < enums.Len(); i++ {
		enum := enums.Get(i)
		visit(enum)
		values := enum.Values()
		for j := 0; j < values.Len(); j++ {
			visit(values.Get(j))
		}
	}
}
