// This file has messages related general definitions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/common/rawdata.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
// Represents a dataset which will be written to datahub.
type WriteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows    []*Row   `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *WriteData) Reset() {
	*x = WriteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteData) ProtoMessage() {}

func (x *WriteData) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteData.ProtoReflect.Descriptor instead.
func (*WriteData) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_common_rawdata_proto_rawDescGZIP(), []int{0}
}

func (x *WriteData) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *WriteData) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

//*
// Represents a dataset whcih is read from datahub.
type ReadData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ReadData) Reset() {
	*x = ReadData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadData) ProtoMessage() {}

func (x *ReadData) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadData.ProtoReflect.Descriptor instead.
func (*ReadData) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_common_rawdata_proto_rawDescGZIP(), []int{1}
}

func (x *ReadData) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

//*
// Represents a record of data.
type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Values []string             `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_common_rawdata_proto_rawDescGZIP(), []int{2}
}

func (x *Row) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Row) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

//*
// Represents a dataset which are collected that have the same attributes.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows    []*Row   `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_common_rawdata_proto_rawDescGZIP(), []int{3}
}

func (x *Group) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Group) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_prophetstor_api_datahub_common_rawdata_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_common_rawdata_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5e, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22,
	0x49, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x4d, 0x0a, 0x03, 0x52, 0x6f,
	0x77, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x77, 0x52,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x42, 0x24, 0x5a, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_common_rawdata_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_common_rawdata_proto_rawDescData = file_prophetstor_api_datahub_common_rawdata_proto_rawDesc
)

func file_prophetstor_api_datahub_common_rawdata_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_common_rawdata_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_common_rawdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_common_rawdata_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_common_rawdata_proto_rawDescData
}

var file_prophetstor_api_datahub_common_rawdata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prophetstor_api_datahub_common_rawdata_proto_goTypes = []interface{}{
	(*WriteData)(nil),           // 0: prophetstor.api.datahub.common.WriteData
	(*ReadData)(nil),            // 1: prophetstor.api.datahub.common.ReadData
	(*Row)(nil),                 // 2: prophetstor.api.datahub.common.Row
	(*Group)(nil),               // 3: prophetstor.api.datahub.common.Group
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_prophetstor_api_datahub_common_rawdata_proto_depIdxs = []int32{
	2, // 0: prophetstor.api.datahub.common.WriteData.rows:type_name -> prophetstor.api.datahub.common.Row
	3, // 1: prophetstor.api.datahub.common.ReadData.groups:type_name -> prophetstor.api.datahub.common.Group
	4, // 2: prophetstor.api.datahub.common.Row.time:type_name -> google.protobuf.Timestamp
	2, // 3: prophetstor.api.datahub.common.Group.rows:type_name -> prophetstor.api.datahub.common.Row
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_common_rawdata_proto_init() }
func file_prophetstor_api_datahub_common_rawdata_proto_init() {
	if File_prophetstor_api_datahub_common_rawdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteData); i {
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
		file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadData); i {
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
		file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_prophetstor_api_datahub_common_rawdata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_prophetstor_api_datahub_common_rawdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_common_rawdata_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_common_rawdata_proto_depIdxs,
		MessageInfos:      file_prophetstor_api_datahub_common_rawdata_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_common_rawdata_proto = out.File
	file_prophetstor_api_datahub_common_rawdata_proto_rawDesc = nil
	file_prophetstor_api_datahub_common_rawdata_proto_goTypes = nil
	file_prophetstor_api_datahub_common_rawdata_proto_depIdxs = nil
}
