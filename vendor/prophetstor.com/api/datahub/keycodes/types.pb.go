// This file has messages related to keycode managements

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/keycodes/types.proto

package keycodes

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

type Capacity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users int32 `protobuf:"varint,1,opt,name=users,proto3" json:"users,omitempty"` // example: "-1"`
	Hosts int32 `protobuf:"varint,2,opt,name=hosts,proto3" json:"hosts,omitempty"` // example: "20"`
	Disks int32 `protobuf:"varint,3,opt,name=disks,proto3" json:"disks,omitempty"` // example: "200"`
}

func (x *Capacity) Reset() {
	*x = Capacity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capacity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capacity) ProtoMessage() {}

func (x *Capacity) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capacity.ProtoReflect.Descriptor instead.
func (*Capacity) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_keycodes_types_proto_rawDescGZIP(), []int{0}
}

func (x *Capacity) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

func (x *Capacity) GetHosts() int32 {
	if x != nil {
		return x.Hosts
	}
	return 0
}

func (x *Capacity) GetDisks() int32 {
	if x != nil {
		return x.Disks
	}
	return 0
}

type Functionality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiskProphet bool `protobuf:"varint,1,opt,name=disk_prophet,json=diskProphet,proto3" json:"disk_prophet,omitempty"` // example: "true"`
	Workload    bool `protobuf:"varint,2,opt,name=workload,proto3" json:"workload,omitempty"`                          // example: "true"`
}

func (x *Functionality) Reset() {
	*x = Functionality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Functionality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Functionality) ProtoMessage() {}

func (x *Functionality) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Functionality.ProtoReflect.Descriptor instead.
func (*Functionality) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_keycodes_types_proto_rawDescGZIP(), []int{1}
}

func (x *Functionality) GetDiskProphet() bool {
	if x != nil {
		return x.DiskProphet
	}
	return false
}

func (x *Functionality) GetWorkload() bool {
	if x != nil {
		return x.Workload
	}
	return false
}

type Retention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidMonth int32 `protobuf:"varint,1,opt,name=valid_month,json=validMonth,proto3" json:"valid_month,omitempty"` // example: "0"`
	Years      int32 `protobuf:"varint,2,opt,name=years,proto3" json:"years,omitempty"`                             // example: "1"`
}

func (x *Retention) Reset() {
	*x = Retention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retention) ProtoMessage() {}

func (x *Retention) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retention.ProtoReflect.Descriptor instead.
func (*Retention) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_keycodes_types_proto_rawDescGZIP(), []int{2}
}

func (x *Retention) GetValidMonth() int32 {
	if x != nil {
		return x.ValidMonth
	}
	return 0
}

func (x *Retention) GetYears() int32 {
	if x != nil {
		return x.Years
	}
	return 0
}

type ServiceAgreement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceAgreement) Reset() {
	*x = ServiceAgreement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAgreement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAgreement) ProtoMessage() {}

func (x *ServiceAgreement) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAgreement.ProtoReflect.Descriptor instead.
func (*ServiceAgreement) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_keycodes_types_proto_rawDescGZIP(), []int{3}
}

var File_prophetstor_api_datahub_keycodes_types_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_keycodes_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x4c, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x73, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x22, 0x4e,
	0x0a, 0x0d, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x42,
	0x0a, 0x09, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x79, 0x65, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x79, 0x65, 0x61,
	0x72, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x67, 0x72,
	0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6b, 0x65, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_keycodes_types_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_keycodes_types_proto_rawDescData = file_prophetstor_api_datahub_keycodes_types_proto_rawDesc
)

func file_prophetstor_api_datahub_keycodes_types_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_keycodes_types_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_keycodes_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_keycodes_types_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_keycodes_types_proto_rawDescData
}

var file_prophetstor_api_datahub_keycodes_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prophetstor_api_datahub_keycodes_types_proto_goTypes = []interface{}{
	(*Capacity)(nil),         // 0: prophetstor.api.datahub.keycodes.Capacity
	(*Functionality)(nil),    // 1: prophetstor.api.datahub.keycodes.Functionality
	(*Retention)(nil),        // 2: prophetstor.api.datahub.keycodes.Retention
	(*ServiceAgreement)(nil), // 3: prophetstor.api.datahub.keycodes.ServiceAgreement
}
var file_prophetstor_api_datahub_keycodes_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_keycodes_types_proto_init() }
func file_prophetstor_api_datahub_keycodes_types_proto_init() {
	if File_prophetstor_api_datahub_keycodes_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capacity); i {
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
		file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Functionality); i {
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
		file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retention); i {
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
		file_prophetstor_api_datahub_keycodes_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAgreement); i {
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
			RawDescriptor: file_prophetstor_api_datahub_keycodes_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_keycodes_types_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_keycodes_types_proto_depIdxs,
		MessageInfos:      file_prophetstor_api_datahub_keycodes_types_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_keycodes_types_proto = out.File
	file_prophetstor_api_datahub_keycodes_types_proto_rawDesc = nil
	file_prophetstor_api_datahub_keycodes_types_proto_goTypes = nil
	file_prophetstor_api_datahub_keycodes_types_proto_depIdxs = nil
}
