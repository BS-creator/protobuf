// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/golang/protobuf/ptypes/struct/struct.proto

package structpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

// Symbols defined in public import of google/protobuf/struct.proto

type NullValue = structpb.NullValue

const NullValue_NULL_VALUE = structpb.NullValue_NULL_VALUE

var NullValue_name = structpb.NullValue_name
var NullValue_value = structpb.NullValue_value

type Struct = structpb.Struct
type Value = structpb.Value
type Value_NullValue = structpb.Value_NullValue
type Value_NumberValue = structpb.Value_NumberValue
type Value_StringValue = structpb.Value_StringValue
type Value_BoolValue = structpb.Value_BoolValue
type Value_StructValue = structpb.Value_StructValue
type Value_ListValue = structpb.Value_ListValue
type ListValue = structpb.ListValue

var File_github_com_golang_protobuf_ptypes_struct_struct_proto protoreflect.FileDescriptor

var file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x3b, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x70, 0x62, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescOnce sync.Once
	file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescData = file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc
)

func file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescGZIP() []byte {
	file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescOnce.Do(func() {
		file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescData)
	})
	return file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescData
}

var file_github_com_golang_protobuf_ptypes_struct_struct_proto_goTypes = []interface{}{}
var file_github_com_golang_protobuf_ptypes_struct_struct_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_github_com_golang_protobuf_ptypes_struct_struct_proto_init() }
func file_github_com_golang_protobuf_ptypes_struct_struct_proto_init() {
	if File_github_com_golang_protobuf_ptypes_struct_struct_proto != nil {
		return
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_golang_protobuf_ptypes_struct_struct_proto_goTypes,
		DependencyIndexes: file_github_com_golang_protobuf_ptypes_struct_struct_proto_depIdxs,
	}.Build()
	File_github_com_golang_protobuf_ptypes_struct_struct_proto = out.File
	file_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc = nil
	file_github_com_golang_protobuf_ptypes_struct_struct_proto_goTypes = nil
	file_github_com_golang_protobuf_ptypes_struct_struct_proto_depIdxs = nil
}
