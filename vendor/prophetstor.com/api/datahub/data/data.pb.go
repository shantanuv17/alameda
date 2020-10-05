// This file has messages related to read & write data

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/data/data.proto

package data

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

//*
// Represents the data which is to be written to datahub.
type WriteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement      string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType       common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=prophetstor.api.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=prophetstor.api.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota    common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=prophetstor.api.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	Columns          []string                `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows             []*common.Row           `protobuf:"bytes,6,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *WriteData) Reset() {
	*x = WriteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteData) ProtoMessage() {}

func (x *WriteData) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[0]
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
	return file_prophetstor_api_datahub_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *WriteData) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *WriteData) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *WriteData) GetResourceBoundary() common.ResourceBoundary {
	if x != nil {
		return x.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (x *WriteData) GetResourceQuota() common.ResourceQuota {
	if x != nil {
		return x.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (x *WriteData) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *WriteData) GetRows() []*common.Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

//*
// Represents the condition of reading data from datahub.
type ReadData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement      string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType       common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=prophetstor.api.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=prophetstor.api.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota    common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=prophetstor.api.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	QueryCondition   *common.QueryCondition  `protobuf:"bytes,5,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
}

func (x *ReadData) Reset() {
	*x = ReadData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadData) ProtoMessage() {}

func (x *ReadData) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[1]
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
	return file_prophetstor_api_datahub_data_data_proto_rawDescGZIP(), []int{1}
}

func (x *ReadData) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *ReadData) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *ReadData) GetResourceBoundary() common.ResourceBoundary {
	if x != nil {
		return x.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (x *ReadData) GetResourceQuota() common.ResourceQuota {
	if x != nil {
		return x.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (x *ReadData) GetQueryCondition() *common.QueryCondition {
	if x != nil {
		return x.QueryCondition
	}
	return nil
}

//*
// Represents the condition of deleting data in datahub.
type DeleteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement      string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType       common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=prophetstor.api.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=prophetstor.api.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota    common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=prophetstor.api.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	QueryCondition   *common.QueryCondition  `protobuf:"bytes,5,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
}

func (x *DeleteData) Reset() {
	*x = DeleteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteData) ProtoMessage() {}

func (x *DeleteData) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteData.ProtoReflect.Descriptor instead.
func (*DeleteData) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_data_data_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteData) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *DeleteData) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *DeleteData) GetResourceBoundary() common.ResourceBoundary {
	if x != nil {
		return x.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (x *DeleteData) GetResourceQuota() common.ResourceQuota {
	if x != nil {
		return x.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (x *DeleteData) GetQueryCondition() *common.QueryCondition {
	if x != nil {
		return x.QueryCondition
	}
	return nil
}

//*
// Represents the data(none time-series) which is to be written to datahub.
type WriteMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement      string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType       common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=prophetstor.api.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=prophetstor.api.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota    common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=prophetstor.api.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	Condition        *common.Condition       `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
	Columns          []string                `protobuf:"bytes,6,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows             []*common.Row           `protobuf:"bytes,7,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *WriteMeta) Reset() {
	*x = WriteMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteMeta) ProtoMessage() {}

func (x *WriteMeta) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_data_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteMeta.ProtoReflect.Descriptor instead.
func (*WriteMeta) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_data_data_proto_rawDescGZIP(), []int{3}
}

func (x *WriteMeta) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *WriteMeta) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *WriteMeta) GetResourceBoundary() common.ResourceBoundary {
	if x != nil {
		return x.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (x *WriteMeta) GetResourceQuota() common.ResourceQuota {
	if x != nil {
		return x.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (x *WriteMeta) GetCondition() *common.Condition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *WriteMeta) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *WriteMeta) GetRows() []*common.Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_prophetstor_api_datahub_data_data_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_data_data_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03,
	0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x11, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x72, 0x6f, 0x77,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x22, 0x87, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5d,
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x12, 0x54, 0x0a,
	0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x03, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x11, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12,
	0x57, 0x0a, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcb, 0x03, 0x0a, 0x09, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
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
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x37, 0x0a,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x77,
	0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_data_data_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_data_data_proto_rawDescData = file_prophetstor_api_datahub_data_data_proto_rawDesc
)

func file_prophetstor_api_datahub_data_data_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_data_data_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_data_data_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_data_data_proto_rawDescData
}

var file_prophetstor_api_datahub_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prophetstor_api_datahub_data_data_proto_goTypes = []interface{}{
	(*WriteData)(nil),             // 0: prophetstor.api.datahub.data.WriteData
	(*ReadData)(nil),              // 1: prophetstor.api.datahub.data.ReadData
	(*DeleteData)(nil),            // 2: prophetstor.api.datahub.data.DeleteData
	(*WriteMeta)(nil),             // 3: prophetstor.api.datahub.data.WriteMeta
	(common.MetricType)(0),        // 4: prophetstor.api.datahub.common.MetricType
	(common.ResourceBoundary)(0),  // 5: prophetstor.api.datahub.common.ResourceBoundary
	(common.ResourceQuota)(0),     // 6: prophetstor.api.datahub.common.ResourceQuota
	(*common.Row)(nil),            // 7: prophetstor.api.datahub.common.Row
	(*common.QueryCondition)(nil), // 8: prophetstor.api.datahub.common.QueryCondition
	(*common.Condition)(nil),      // 9: prophetstor.api.datahub.common.Condition
}
var file_prophetstor_api_datahub_data_data_proto_depIdxs = []int32{
	4,  // 0: prophetstor.api.datahub.data.WriteData.metric_type:type_name -> prophetstor.api.datahub.common.MetricType
	5,  // 1: prophetstor.api.datahub.data.WriteData.resource_boundary:type_name -> prophetstor.api.datahub.common.ResourceBoundary
	6,  // 2: prophetstor.api.datahub.data.WriteData.resource_quota:type_name -> prophetstor.api.datahub.common.ResourceQuota
	7,  // 3: prophetstor.api.datahub.data.WriteData.rows:type_name -> prophetstor.api.datahub.common.Row
	4,  // 4: prophetstor.api.datahub.data.ReadData.metric_type:type_name -> prophetstor.api.datahub.common.MetricType
	5,  // 5: prophetstor.api.datahub.data.ReadData.resource_boundary:type_name -> prophetstor.api.datahub.common.ResourceBoundary
	6,  // 6: prophetstor.api.datahub.data.ReadData.resource_quota:type_name -> prophetstor.api.datahub.common.ResourceQuota
	8,  // 7: prophetstor.api.datahub.data.ReadData.query_condition:type_name -> prophetstor.api.datahub.common.QueryCondition
	4,  // 8: prophetstor.api.datahub.data.DeleteData.metric_type:type_name -> prophetstor.api.datahub.common.MetricType
	5,  // 9: prophetstor.api.datahub.data.DeleteData.resource_boundary:type_name -> prophetstor.api.datahub.common.ResourceBoundary
	6,  // 10: prophetstor.api.datahub.data.DeleteData.resource_quota:type_name -> prophetstor.api.datahub.common.ResourceQuota
	8,  // 11: prophetstor.api.datahub.data.DeleteData.query_condition:type_name -> prophetstor.api.datahub.common.QueryCondition
	4,  // 12: prophetstor.api.datahub.data.WriteMeta.metric_type:type_name -> prophetstor.api.datahub.common.MetricType
	5,  // 13: prophetstor.api.datahub.data.WriteMeta.resource_boundary:type_name -> prophetstor.api.datahub.common.ResourceBoundary
	6,  // 14: prophetstor.api.datahub.data.WriteMeta.resource_quota:type_name -> prophetstor.api.datahub.common.ResourceQuota
	9,  // 15: prophetstor.api.datahub.data.WriteMeta.condition:type_name -> prophetstor.api.datahub.common.Condition
	7,  // 16: prophetstor.api.datahub.data.WriteMeta.rows:type_name -> prophetstor.api.datahub.common.Row
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_data_data_proto_init() }
func file_prophetstor_api_datahub_data_data_proto_init() {
	if File_prophetstor_api_datahub_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_prophetstor_api_datahub_data_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_prophetstor_api_datahub_data_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteData); i {
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
		file_prophetstor_api_datahub_data_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteMeta); i {
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
			RawDescriptor: file_prophetstor_api_datahub_data_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_data_data_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_data_data_proto_depIdxs,
		MessageInfos:      file_prophetstor_api_datahub_data_data_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_data_data_proto = out.File
	file_prophetstor_api_datahub_data_data_proto_rawDesc = nil
	file_prophetstor_api_datahub_data_data_proto_goTypes = nil
	file_prophetstor_api_datahub_data_data_proto_depIdxs = nil
}
