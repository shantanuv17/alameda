// This file has messages to describe alameda scaler structure

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/applications/types.proto

package applications

import (
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
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
// Represents a dataset of private alameda scalers.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta      *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	ApplicationData []*ApplicationData  `protobuf:"bytes,2,rep,name=application_data,json=applicationData,proto3" json:"application_data,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *Application) GetApplicationData() []*ApplicationData {
	if x != nil {
		return x.ApplicationData
	}
	return nil
}

//*
// Represents a private alameda scaler.
type ApplicationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement string           `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	ReadData    *common.ReadData `protobuf:"bytes,2,opt,name=read_data,json=readData,proto3" json:"read_data,omitempty"`
}

func (x *ApplicationData) Reset() {
	*x = ApplicationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationData) ProtoMessage() {}

func (x *ApplicationData) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationData.ProtoReflect.Descriptor instead.
func (*ApplicationData) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationData) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *ApplicationData) GetReadData() *common.ReadData {
	if x != nil {
		return x.ReadData
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_applications_types_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x31, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5a, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x6e, 0x0a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x88, 0x01, 0x0a,
	0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x53, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescData = file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alameda_api_v1alpha1_datahub_applications_types_proto_goTypes = []interface{}{
	(*Application)(nil),        // 0: containersai.alameda.v1alpha1.datahub.applications.Application
	(*ApplicationData)(nil),    // 1: containersai.alameda.v1alpha1.datahub.applications.ApplicationData
	(*schemas.SchemaMeta)(nil), // 2: containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	(*common.ReadData)(nil),    // 3: containersai.alameda.v1alpha1.datahub.common.ReadData
}
var file_alameda_api_v1alpha1_datahub_applications_types_proto_depIdxs = []int32{
	2, // 0: containersai.alameda.v1alpha1.datahub.applications.Application.schema_meta:type_name -> containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	1, // 1: containersai.alameda.v1alpha1.datahub.applications.Application.application_data:type_name -> containersai.alameda.v1alpha1.datahub.applications.ApplicationData
	3, // 2: containersai.alameda.v1alpha1.datahub.applications.ApplicationData.read_data:type_name -> containersai.alameda.v1alpha1.datahub.common.ReadData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_applications_types_proto_init() }
func file_alameda_api_v1alpha1_datahub_applications_types_proto_init() {
	if File_alameda_api_v1alpha1_datahub_applications_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationData); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_applications_types_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_applications_types_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_applications_types_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_applications_types_proto = out.File
	file_alameda_api_v1alpha1_datahub_applications_types_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_applications_types_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_applications_types_proto_depIdxs = nil
}
