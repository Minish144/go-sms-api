// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: common.proto

package pb

import (
	validate "github.com/envoyproxy/protoc-gen-validate/validate"
	options "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	annotations "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/annotations.proto.

var E_Http = annotations.E_Http

// Symbols defined in public import of github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2/options/annotations.proto.

var E_Openapiv2Swagger = options.E_Openapiv2Swagger
var E_Openapiv2Operation = options.E_Openapiv2Operation
var E_Openapiv2Schema = options.E_Openapiv2Schema
var E_Openapiv2Tag = options.E_Openapiv2Tag
var E_Openapiv2Field = options.E_Openapiv2Field

// Symbols defined in public import of github.com/envoyproxy/protoc-gen-validate/validate/validate.proto.

type KnownRegex = validate.KnownRegex

const KnownRegex_UNKNOWN = validate.KnownRegex_UNKNOWN
const KnownRegex_HTTP_HEADER_NAME = validate.KnownRegex_HTTP_HEADER_NAME
const KnownRegex_HTTP_HEADER_VALUE = validate.KnownRegex_HTTP_HEADER_VALUE

var KnownRegex_name = validate.KnownRegex_name
var KnownRegex_value = validate.KnownRegex_value

type FieldRules = validate.FieldRules
type FieldRules_Float = validate.FieldRules_Float
type FieldRules_Double = validate.FieldRules_Double
type FieldRules_Int32 = validate.FieldRules_Int32
type FieldRules_Int64 = validate.FieldRules_Int64
type FieldRules_Uint32 = validate.FieldRules_Uint32
type FieldRules_Uint64 = validate.FieldRules_Uint64
type FieldRules_Sint32 = validate.FieldRules_Sint32
type FieldRules_Sint64 = validate.FieldRules_Sint64
type FieldRules_Fixed32 = validate.FieldRules_Fixed32
type FieldRules_Fixed64 = validate.FieldRules_Fixed64
type FieldRules_Sfixed32 = validate.FieldRules_Sfixed32
type FieldRules_Sfixed64 = validate.FieldRules_Sfixed64
type FieldRules_Bool = validate.FieldRules_Bool
type FieldRules_String_ = validate.FieldRules_String_
type FieldRules_Bytes = validate.FieldRules_Bytes
type FieldRules_Enum = validate.FieldRules_Enum
type FieldRules_Repeated = validate.FieldRules_Repeated
type FieldRules_Map = validate.FieldRules_Map
type FieldRules_Any = validate.FieldRules_Any
type FieldRules_Duration = validate.FieldRules_Duration
type FieldRules_Timestamp = validate.FieldRules_Timestamp
type FloatRules = validate.FloatRules
type DoubleRules = validate.DoubleRules
type Int32Rules = validate.Int32Rules
type Int64Rules = validate.Int64Rules
type UInt32Rules = validate.UInt32Rules
type UInt64Rules = validate.UInt64Rules
type SInt32Rules = validate.SInt32Rules
type SInt64Rules = validate.SInt64Rules
type Fixed32Rules = validate.Fixed32Rules
type Fixed64Rules = validate.Fixed64Rules
type SFixed32Rules = validate.SFixed32Rules
type SFixed64Rules = validate.SFixed64Rules
type BoolRules = validate.BoolRules
type StringRules = validate.StringRules

const Default_StringRules_Strict = validate.Default_StringRules_Strict

type StringRules_Email = validate.StringRules_Email
type StringRules_Hostname = validate.StringRules_Hostname
type StringRules_Ip = validate.StringRules_Ip
type StringRules_Ipv4 = validate.StringRules_Ipv4
type StringRules_Ipv6 = validate.StringRules_Ipv6
type StringRules_Uri = validate.StringRules_Uri
type StringRules_UriRef = validate.StringRules_UriRef
type StringRules_Address = validate.StringRules_Address
type StringRules_Uuid = validate.StringRules_Uuid
type StringRules_WellKnownRegex = validate.StringRules_WellKnownRegex
type BytesRules = validate.BytesRules
type BytesRules_Ip = validate.BytesRules_Ip
type BytesRules_Ipv4 = validate.BytesRules_Ipv4
type BytesRules_Ipv6 = validate.BytesRules_Ipv6
type EnumRules = validate.EnumRules
type MessageRules = validate.MessageRules
type RepeatedRules = validate.RepeatedRules
type MapRules = validate.MapRules
type AnyRules = validate.AnyRules
type DurationRules = validate.DurationRules
type TimestampRules = validate.TimestampRules

var E_Disabled = validate.E_Disabled
var E_Ignored = validate.E_Ignored
var E_Required = validate.E_Required
var E_Rules = validate.E_Rules

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x1a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x23, 0xfa, 0x42, 0x20, 0x72, 0x1e, 0x32, 0x1c, 0x5e, 0x5c, 0x2b, 0x28,
	0x3f, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0xe2, 0x97, 0x8f, 0x3f, 0x29, 0x7b, 0x36, 0x2c, 0x31,
	0x34, 0x7d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x62, 0x50, 0x00, 0x50, 0x01, 0x50, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: api.Message
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
