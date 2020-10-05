// This file has messages related to license managements

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/licenses/licenses.proto

package licenses

import (
	proto "github.com/golang/protobuf/proto"
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

//*
// Represents the information of a license.
type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_licenses_licenses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_licenses_licenses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_licenses_licenses_proto_rawDescGZIP(), []int{0}
}

func (x *License) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_prophetstor_api_datahub_licenses_licenses_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_licenses_licenses_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x73, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x42, 0x26, 0x5a, 0x24, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_licenses_licenses_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_licenses_licenses_proto_rawDescData = file_prophetstor_api_datahub_licenses_licenses_proto_rawDesc
)

func file_prophetstor_api_datahub_licenses_licenses_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_licenses_licenses_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_licenses_licenses_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_licenses_licenses_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_licenses_licenses_proto_rawDescData
}

var file_prophetstor_api_datahub_licenses_licenses_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prophetstor_api_datahub_licenses_licenses_proto_goTypes = []interface{}{
	(*License)(nil), // 0: prophetstor.api.datahub.licenses.License
}
var file_prophetstor_api_datahub_licenses_licenses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_licenses_licenses_proto_init() }
func file_prophetstor_api_datahub_licenses_licenses_proto_init() {
	if File_prophetstor_api_datahub_licenses_licenses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_licenses_licenses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License); i {
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
			RawDescriptor: file_prophetstor_api_datahub_licenses_licenses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_licenses_licenses_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_licenses_licenses_proto_depIdxs,
		MessageInfos:      file_prophetstor_api_datahub_licenses_licenses_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_licenses_licenses_proto = out.File
	file_prophetstor_api_datahub_licenses_licenses_proto_rawDesc = nil
	file_prophetstor_api_datahub_licenses_licenses_proto_goTypes = nil
	file_prophetstor_api_datahub_licenses_licenses_proto_depIdxs = nil
}
