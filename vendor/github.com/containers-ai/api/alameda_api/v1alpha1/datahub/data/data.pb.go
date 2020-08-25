// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/data/data.proto

package data

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//*
// Represents the data which is to be written to datahub.
type WriteData struct {
	Measurement          string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType           common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary     common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota        common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	Columns              []string                `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*common.Row           `protobuf:"bytes,6,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *WriteData) Reset()         { *m = WriteData{} }
func (m *WriteData) String() string { return proto.CompactTextString(m) }
func (*WriteData) ProtoMessage()    {}
func (*WriteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d0076597afc4d7, []int{0}
}

func (m *WriteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteData.Unmarshal(m, b)
}
func (m *WriteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteData.Marshal(b, m, deterministic)
}
func (m *WriteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteData.Merge(m, src)
}
func (m *WriteData) XXX_Size() int {
	return xxx_messageInfo_WriteData.Size(m)
}
func (m *WriteData) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteData.DiscardUnknown(m)
}

var xxx_messageInfo_WriteData proto.InternalMessageInfo

func (m *WriteData) GetMeasurement() string {
	if m != nil {
		return m.Measurement
	}
	return ""
}

func (m *WriteData) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *WriteData) GetResourceBoundary() common.ResourceBoundary {
	if m != nil {
		return m.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (m *WriteData) GetResourceQuota() common.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (m *WriteData) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *WriteData) GetRows() []*common.Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

//*
// Represents the condition of reading data from datahub.
type ReadData struct {
	Measurement          string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType           common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary     common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota        common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	QueryCondition       *common.QueryCondition  `protobuf:"bytes,5,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ReadData) Reset()         { *m = ReadData{} }
func (m *ReadData) String() string { return proto.CompactTextString(m) }
func (*ReadData) ProtoMessage()    {}
func (*ReadData) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d0076597afc4d7, []int{1}
}

func (m *ReadData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadData.Unmarshal(m, b)
}
func (m *ReadData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadData.Marshal(b, m, deterministic)
}
func (m *ReadData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadData.Merge(m, src)
}
func (m *ReadData) XXX_Size() int {
	return xxx_messageInfo_ReadData.Size(m)
}
func (m *ReadData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadData.DiscardUnknown(m)
}

var xxx_messageInfo_ReadData proto.InternalMessageInfo

func (m *ReadData) GetMeasurement() string {
	if m != nil {
		return m.Measurement
	}
	return ""
}

func (m *ReadData) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *ReadData) GetResourceBoundary() common.ResourceBoundary {
	if m != nil {
		return m.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (m *ReadData) GetResourceQuota() common.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (m *ReadData) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

//*
// Represents the condition of deleting data in datahub.
type DeleteData struct {
	Measurement          string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType           common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary     common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota        common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	QueryCondition       *common.QueryCondition  `protobuf:"bytes,5,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DeleteData) Reset()         { *m = DeleteData{} }
func (m *DeleteData) String() string { return proto.CompactTextString(m) }
func (*DeleteData) ProtoMessage()    {}
func (*DeleteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d0076597afc4d7, []int{2}
}

func (m *DeleteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteData.Unmarshal(m, b)
}
func (m *DeleteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteData.Marshal(b, m, deterministic)
}
func (m *DeleteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteData.Merge(m, src)
}
func (m *DeleteData) XXX_Size() int {
	return xxx_messageInfo_DeleteData.Size(m)
}
func (m *DeleteData) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteData.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteData proto.InternalMessageInfo

func (m *DeleteData) GetMeasurement() string {
	if m != nil {
		return m.Measurement
	}
	return ""
}

func (m *DeleteData) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *DeleteData) GetResourceBoundary() common.ResourceBoundary {
	if m != nil {
		return m.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (m *DeleteData) GetResourceQuota() common.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (m *DeleteData) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

//*
// Represents the data(none time-series) which is to be written to datahub.
type WriteMeta struct {
	Measurement          string                  `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	MetricType           common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary     common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota        common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	Condition            *common.Condition       `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
	Columns              []string                `protobuf:"bytes,6,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*common.Row           `protobuf:"bytes,7,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *WriteMeta) Reset()         { *m = WriteMeta{} }
func (m *WriteMeta) String() string { return proto.CompactTextString(m) }
func (*WriteMeta) ProtoMessage()    {}
func (*WriteMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d0076597afc4d7, []int{3}
}

func (m *WriteMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteMeta.Unmarshal(m, b)
}
func (m *WriteMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteMeta.Marshal(b, m, deterministic)
}
func (m *WriteMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteMeta.Merge(m, src)
}
func (m *WriteMeta) XXX_Size() int {
	return xxx_messageInfo_WriteMeta.Size(m)
}
func (m *WriteMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteMeta.DiscardUnknown(m)
}

var xxx_messageInfo_WriteMeta proto.InternalMessageInfo

func (m *WriteMeta) GetMeasurement() string {
	if m != nil {
		return m.Measurement
	}
	return ""
}

func (m *WriteMeta) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *WriteMeta) GetResourceBoundary() common.ResourceBoundary {
	if m != nil {
		return m.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (m *WriteMeta) GetResourceQuota() common.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (m *WriteMeta) GetCondition() *common.Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *WriteMeta) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *WriteMeta) GetRows() []*common.Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteData)(nil), "containersai.alameda.v1alpha1.datahub.data.WriteData")
	proto.RegisterType((*ReadData)(nil), "containersai.alameda.v1alpha1.datahub.data.ReadData")
	proto.RegisterType((*DeleteData)(nil), "containersai.alameda.v1alpha1.datahub.data.DeleteData")
	proto.RegisterType((*WriteMeta)(nil), "containersai.alameda.v1alpha1.datahub.data.WriteMeta")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/data/data.proto", fileDescriptor_27d0076597afc4d7)
}

var fileDescriptor_27d0076597afc4d7 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0x41, 0x8f, 0x94, 0x30,
	0x14, 0xc7, 0x33, 0x32, 0x3b, 0x2b, 0x9d, 0x38, 0x2a, 0xa7, 0x66, 0x4f, 0x64, 0x4e, 0xc4, 0x28,
	0x84, 0xf1, 0xa0, 0x89, 0x66, 0x63, 0xd6, 0xf5, 0xb8, 0x87, 0x6d, 0x34, 0x46, 0x2f, 0xe4, 0x01,
	0x2f, 0x6e, 0x23, 0x6d, 0x99, 0x52, 0x9c, 0xf0, 0x31, 0xfc, 0x3a, 0x5e, 0x8c, 0xdf, 0xcc, 0xd0,
	0x85, 0x65, 0xdc, 0x18, 0x1d, 0x74, 0x2f, 0x9b, 0xcc, 0x85, 0xf2, 0x0a, 0xff, 0xdf, 0x3f, 0xe5,
	0x9f, 0xbe, 0x42, 0x1e, 0x43, 0x01, 0x02, 0x73, 0x48, 0xa0, 0xe4, 0xd1, 0x97, 0x18, 0x8a, 0xf2,
	0x02, 0xe2, 0x28, 0x07, 0x03, 0x17, 0x75, 0x6a, 0x47, 0x7b, 0x09, 0x4b, 0xad, 0x8c, 0xf2, 0x1e,
	0x65, 0x4a, 0x1a, 0xe0, 0x12, 0x75, 0x05, 0x3c, 0xec, 0xa4, 0x61, 0x2f, 0x0b, 0x3b, 0x99, 0x1d,
	0x8f, 0xe2, 0x3f, 0x92, 0x33, 0x25, 0x84, 0x92, 0x91, 0x40, 0xa3, 0x79, 0x56, 0x5d, 0xe2, 0x77,
	0x93, 0xac, 0x6b, 0xd4, 0x1c, 0x47, 0x49, 0x34, 0x6c, 0x86, 0x45, 0x1c, 0x45, 0xbb, 0x48, 0x4c,
	0x53, 0xf6, 0x1e, 0xcb, 0x6f, 0x0e, 0x71, 0xdf, 0x6b, 0x6e, 0xf0, 0x14, 0x0c, 0x78, 0x3e, 0x99,
	0x0b, 0x84, 0xaa, 0xd6, 0x28, 0x50, 0x1a, 0x3a, 0xf1, 0x27, 0x81, 0xcb, 0xb6, 0xa7, 0xbc, 0x0f,
	0xed, 0x1b, 0xed, 0xba, 0x92, 0x96, 0x42, 0xef, 0xf8, 0x93, 0x60, 0xb1, 0x7a, 0x1e, 0xee, 0xf6,
	0xed, 0x2e, 0xfd, 0xc3, 0x33, 0x0b, 0x78, 0xdb, 0x94, 0xc8, 0x88, 0xb8, 0xba, 0xf7, 0x3e, 0x93,
	0x87, 0x1a, 0x2b, 0x55, 0xeb, 0x0c, 0x93, 0x54, 0xd5, 0x32, 0x07, 0xdd, 0x50, 0xc7, 0x1a, 0x1c,
	0x8f, 0x33, 0x60, 0x1d, 0xe6, 0xa4, 0xa3, 0xb0, 0x07, 0xfa, 0xda, 0x8c, 0x97, 0x92, 0xc5, 0x95,
	0xd9, 0xba, 0x56, 0x06, 0xe8, 0xd4, 0x3a, 0xbd, 0xf8, 0x37, 0xa7, 0xf3, 0x16, 0xc1, 0xee, 0xe9,
	0xed, 0xd2, 0xa3, 0xe4, 0x30, 0x53, 0x45, 0x2d, 0x64, 0x45, 0x0f, 0x7c, 0x27, 0x70, 0x59, 0x5f,
	0x7a, 0x6f, 0xc8, 0x54, 0xab, 0x4d, 0x45, 0x67, 0xbe, 0x13, 0xcc, 0x57, 0xf1, 0x48, 0x4f, 0xb5,
	0x61, 0x56, 0xbe, 0xfc, 0xee, 0x90, 0xbb, 0x0c, 0x21, 0xdf, 0x67, 0x77, 0xe3, 0xd9, 0x21, 0xb9,
	0xdf, 0x6e, 0xc6, 0x26, 0xc9, 0x94, 0xcc, 0xb9, 0xe1, 0x4a, 0xd2, 0x03, 0x7f, 0x12, 0xcc, 0x57,
	0x2f, 0xc7, 0x99, 0x9c, 0xb7, 0x90, 0xd7, 0x3d, 0x83, 0x2d, 0xd6, 0xbf, 0xd4, 0xcb, 0x1f, 0x0e,
	0x21, 0xa7, 0x58, 0xe0, 0x7e, 0xff, 0xdd, 0xde, 0x0c, 0xbf, 0x4e, 0xbb, 0x16, 0x7a, 0x86, 0xfb,
	0x08, 0x6f, 0x38, 0xc2, 0x77, 0xc4, 0xbd, 0x1e, 0xde, 0xb3, 0x71, 0xf8, 0x21, 0xb7, 0x81, 0xb4,
	0xdd, 0x99, 0x67, 0xbf, 0xef, 0xcc, 0x87, 0xff, 0xd5, 0x99, 0x4f, 0x5e, 0x7d, 0x3c, 0xfe, 0xc4,
	0x4d, 0xf7, 0x28, 0x1a, 0x20, 0x4f, 0x80, 0x47, 0xed, 0xd1, 0xfc, 0xd7, 0x3f, 0x93, 0x74, 0x66,
	0xcf, 0xe7, 0xa7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53, 0x09, 0x68, 0x94, 0xc5, 0x08, 0x00,
	0x00,
}
