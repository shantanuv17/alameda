// This file has messages related to recommendations of containers, pods, and nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/schemas/types.proto

package schemas

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "prophetstor.com/api/datahub/common"
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

type Scope int32

const (
	Scope_SCOPE_UNDEFINED      Scope = 0
	Scope_SCOPE_APPLICATION    Scope = 1
	Scope_SCOPE_CONFIG         Scope = 2
	Scope_SCOPE_FEDEMETER      Scope = 3
	Scope_SCOPE_METERING       Scope = 4
	Scope_SCOPE_METRIC         Scope = 5
	Scope_SCOPE_PLANNING       Scope = 6
	Scope_SCOPE_PREDICTION     Scope = 7
	Scope_SCOPE_RECOMMENDATION Scope = 8
	Scope_SCOPE_RESOURCE       Scope = 9
	Scope_SCOPE_TARGET         Scope = 10
)

// Enum value maps for Scope.
var (
	Scope_name = map[int32]string{
		0:  "SCOPE_UNDEFINED",
		1:  "SCOPE_APPLICATION",
		2:  "SCOPE_CONFIG",
		3:  "SCOPE_FEDEMETER",
		4:  "SCOPE_METERING",
		5:  "SCOPE_METRIC",
		6:  "SCOPE_PLANNING",
		7:  "SCOPE_PREDICTION",
		8:  "SCOPE_RECOMMENDATION",
		9:  "SCOPE_RESOURCE",
		10: "SCOPE_TARGET",
	}
	Scope_value = map[string]int32{
		"SCOPE_UNDEFINED":      0,
		"SCOPE_APPLICATION":    1,
		"SCOPE_CONFIG":         2,
		"SCOPE_FEDEMETER":      3,
		"SCOPE_METERING":       4,
		"SCOPE_METRIC":         5,
		"SCOPE_PLANNING":       6,
		"SCOPE_PREDICTION":     7,
		"SCOPE_RECOMMENDATION": 8,
		"SCOPE_RESOURCE":       9,
		"SCOPE_TARGET":         10,
	}
)

func (x Scope) Enum() *Scope {
	p := new(Scope)
	*p = x
	return p
}

func (x Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_prophetstor_api_datahub_schemas_types_proto_enumTypes[0].Descriptor()
}

func (Scope) Type() protoreflect.EnumType {
	return &file_prophetstor_api_datahub_schemas_types_proto_enumTypes[0]
}

func (x Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scope.Descriptor instead.
func (Scope) EnumDescriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_schemas_types_proto_rawDescGZIP(), []int{0}
}

//*
// Represents the private metadata of datahub schema.
type SchemaMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope    Scope  `protobuf:"varint,1,opt,name=scope,proto3,enum=prophetstor.api.datahub.schemas.Scope" json:"scope,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SchemaMeta) Reset() {
	*x = SchemaMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_schemas_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaMeta) ProtoMessage() {}

func (x *SchemaMeta) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_schemas_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaMeta.ProtoReflect.Descriptor instead.
func (*SchemaMeta) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_schemas_types_proto_rawDescGZIP(), []int{0}
}

func (x *SchemaMeta) GetScope() Scope {
	if x != nil {
		return x.Scope
	}
	return Scope_SCOPE_UNDEFINED
}

func (x *SchemaMeta) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *SchemaMeta) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

//*
// Represents the information of measurment which datahub will write data in InfluxDB.
type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricType       common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=prophetstor.api.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=prophetstor.api.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota    common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=prophetstor.api.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	IsTs             bool                    `protobuf:"varint,5,opt,name=is_ts,json=isTs,proto3" json:"is_ts,omitempty"`
	Columns          []*Column               `protobuf:"bytes,6,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_schemas_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_schemas_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_schemas_types_proto_rawDescGZIP(), []int{1}
}

func (x *Measurement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Measurement) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *Measurement) GetResourceBoundary() common.ResourceBoundary {
	if x != nil {
		return x.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (x *Measurement) GetResourceQuota() common.ResourceQuota {
	if x != nil {
		return x.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (x *Measurement) GetIsTs() bool {
	if x != nil {
		return x.IsTs
	}
	return false
}

func (x *Measurement) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

//*
// Represents a data record.
type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required   bool              `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	ColumnType common.ColumnType `protobuf:"varint,3,opt,name=column_type,json=columnType,proto3,enum=prophetstor.api.datahub.common.ColumnType" json:"column_type,omitempty"`
	DataType   common.DataType   `protobuf:"varint,4,opt,name=data_type,json=dataType,proto3,enum=prophetstor.api.datahub.common.DataType" json:"data_type,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_schemas_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_schemas_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_schemas_types_proto_rawDescGZIP(), []int{2}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Column) GetColumnType() common.ColumnType {
	if x != nil {
		return x.ColumnType
	}
	return common.ColumnType_COLUMNTYPE_UDEFINED
}

func (x *Column) GetDataType() common.DataType {
	if x != nil {
		return x.DataType
	}
	return common.DataType_DATATYPE_UNDEFINED
}

var File_prophetstor_api_datahub_schemas_types_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_schemas_types_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x1a, 0x2c,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xfb, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73,
	0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x54, 0x73, 0x12,
	0x41, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x4b, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x2a, 0xea, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x4f, 0x50, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43, 0x4f,
	0x50, 0x45, 0x5f, 0x46, 0x45, 0x44, 0x45, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x50, 0x4c,
	0x41, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x4f, 0x50,
	0x45, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x43, 0x4f, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x0a, 0x42, 0x25,
	0x5a, 0x23, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_schemas_types_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_schemas_types_proto_rawDescData = file_prophetstor_api_datahub_schemas_types_proto_rawDesc
)

func file_prophetstor_api_datahub_schemas_types_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_schemas_types_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_schemas_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_schemas_types_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_schemas_types_proto_rawDescData
}

var file_prophetstor_api_datahub_schemas_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prophetstor_api_datahub_schemas_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prophetstor_api_datahub_schemas_types_proto_goTypes = []interface{}{
	(Scope)(0),                   // 0: prophetstor.api.datahub.schemas.Scope
	(*SchemaMeta)(nil),           // 1: prophetstor.api.datahub.schemas.SchemaMeta
	(*Measurement)(nil),          // 2: prophetstor.api.datahub.schemas.Measurement
	(*Column)(nil),               // 3: prophetstor.api.datahub.schemas.Column
	(common.MetricType)(0),       // 4: prophetstor.api.datahub.common.MetricType
	(common.ResourceBoundary)(0), // 5: prophetstor.api.datahub.common.ResourceBoundary
	(common.ResourceQuota)(0),    // 6: prophetstor.api.datahub.common.ResourceQuota
	(common.ColumnType)(0),       // 7: prophetstor.api.datahub.common.ColumnType
	(common.DataType)(0),         // 8: prophetstor.api.datahub.common.DataType
}
var file_prophetstor_api_datahub_schemas_types_proto_depIdxs = []int32{
	0, // 0: prophetstor.api.datahub.schemas.SchemaMeta.scope:type_name -> prophetstor.api.datahub.schemas.Scope
	4, // 1: prophetstor.api.datahub.schemas.Measurement.metric_type:type_name -> prophetstor.api.datahub.common.MetricType
	5, // 2: prophetstor.api.datahub.schemas.Measurement.resource_boundary:type_name -> prophetstor.api.datahub.common.ResourceBoundary
	6, // 3: prophetstor.api.datahub.schemas.Measurement.resource_quota:type_name -> prophetstor.api.datahub.common.ResourceQuota
	3, // 4: prophetstor.api.datahub.schemas.Measurement.columns:type_name -> prophetstor.api.datahub.schemas.Column
	7, // 5: prophetstor.api.datahub.schemas.Column.column_type:type_name -> prophetstor.api.datahub.common.ColumnType
	8, // 6: prophetstor.api.datahub.schemas.Column.data_type:type_name -> prophetstor.api.datahub.common.DataType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_schemas_types_proto_init() }
func file_prophetstor_api_datahub_schemas_types_proto_init() {
	if File_prophetstor_api_datahub_schemas_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_schemas_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaMeta); i {
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
		file_prophetstor_api_datahub_schemas_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
		file_prophetstor_api_datahub_schemas_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
			RawDescriptor: file_prophetstor_api_datahub_schemas_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_schemas_types_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_schemas_types_proto_depIdxs,
		EnumInfos:         file_prophetstor_api_datahub_schemas_types_proto_enumTypes,
		MessageInfos:      file_prophetstor_api_datahub_schemas_types_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_schemas_types_proto = out.File
	file_prophetstor_api_datahub_schemas_types_proto_rawDesc = nil
	file_prophetstor_api_datahub_schemas_types_proto_goTypes = nil
	file_prophetstor_api_datahub_schemas_types_proto_depIdxs = nil
}
