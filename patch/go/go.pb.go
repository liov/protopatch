// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: patch/go.proto

package patch_go

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Options represent Go-specific options for Protobuf messages, fields, oneofs, enums, or enum values.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name option renames the generated Go identifier and related identifiers.
	// For a message, this renames the generated Go struct and nested messages or enums, if any.
	// For a message field, this renames the generated Go struct field and getter method.
	// For a oneof field, this renames the generated Go struct field, getter method, interface type, and wrapper types.
	// For an enum, this renames the generated Go type.
	// For an enum value, this renames the generated Go const.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The getter_name option renames the generated getter method (default: Get<Field>)
	// so a custom getter can be implemented in its place.
	GetterName *string `protobuf:"bytes,10,opt,name=getter_name,json=getterName" json:"getter_name,omitempty"` // TODO: implement this
	// The tags option specifies additional struct tags which are appended a generated Go struct field.
	// This option may be specified on a message field or a oneof field.
	// The value should omit the enclosing backticks.
	Tags *string `protobuf:"bytes,20,opt,name=tags" json:"tags,omitempty"`
	// The stringer_name option renames a generated String() method (if any)
	// so a custom String() method can be implemented in its place.
	StringerName *string `protobuf:"bytes,30,opt,name=stringer_name,json=stringerName" json:"stringer_name,omitempty"` // TODO: implement for messages
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patch_go_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_patch_go_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_patch_go_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Options) GetGetterName() string {
	if x != nil && x.GetterName != nil {
		return *x.GetterName
	}
	return ""
}

func (x *Options) GetTags() string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return ""
}

func (x *Options) GetStringerName() string {
	if x != nil && x.StringerName != nil {
		return *x.StringerName
	}
	return ""
}

var file_patch_go_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         7002,
		Name:          "go.options",
		Tag:           "bytes,7002,opt,name=options",
		Filename:      "patch/go.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// Example: int id = 1 [(go.options) = {name: 'ID'}];
	//
	// optional go.Options options = 7002;
	E_Options = &file_patch_go_proto_extTypes[0]
)

var File_patch_go_proto protoreflect.FileDescriptor

var file_patch_go_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x67, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x3a,
	0x45, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xda, 0x36, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x67, 0x6f,
}

var (
	file_patch_go_proto_rawDescOnce sync.Once
	file_patch_go_proto_rawDescData = file_patch_go_proto_rawDesc
)

func file_patch_go_proto_rawDescGZIP() []byte {
	file_patch_go_proto_rawDescOnce.Do(func() {
		file_patch_go_proto_rawDescData = protoimpl.X.CompressGZIP(file_patch_go_proto_rawDescData)
	})
	return file_patch_go_proto_rawDescData
}

var file_patch_go_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_patch_go_proto_goTypes = []interface{}{
	(*Options)(nil),                 // 0: go.Options
	(*descriptor.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_patch_go_proto_depIdxs = []int32{
	1, // 0: go.options:extendee -> google.protobuf.FieldOptions
	0, // 1: go.options:type_name -> go.Options
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_patch_go_proto_init() }
func file_patch_go_proto_init() {
	if File_patch_go_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_patch_go_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_patch_go_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_patch_go_proto_goTypes,
		DependencyIndexes: file_patch_go_proto_depIdxs,
		MessageInfos:      file_patch_go_proto_msgTypes,
		ExtensionInfos:    file_patch_go_proto_extTypes,
	}.Build()
	File_patch_go_proto = out.File
	file_patch_go_proto_rawDesc = nil
	file_patch_go_proto_goTypes = nil
	file_patch_go_proto_depIdxs = nil
}
