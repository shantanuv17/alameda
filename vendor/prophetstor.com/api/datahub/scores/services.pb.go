// This file has messages and services related to Containers.ai

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/scores/services.proto

package scores

import (
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Represents a request for listing system scores of pod scheduled on node
type ListSimulatedSchedulingScoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryCondition *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
}

func (x *ListSimulatedSchedulingScoresRequest) Reset() {
	*x = ListSimulatedSchedulingScoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_scores_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSimulatedSchedulingScoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSimulatedSchedulingScoresRequest) ProtoMessage() {}

func (x *ListSimulatedSchedulingScoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_scores_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSimulatedSchedulingScoresRequest.ProtoReflect.Descriptor instead.
func (*ListSimulatedSchedulingScoresRequest) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_scores_services_proto_rawDescGZIP(), []int{0}
}

func (x *ListSimulatedSchedulingScoresRequest) GetQueryCondition() *common.QueryCondition {
	if x != nil {
		return x.QueryCondition
	}
	return nil
}

// Represents a response for listing system scores request
type ListSimulatedSchedulingScoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status              `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Scores []*SimulatedSchedulingScore `protobuf:"bytes,2,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *ListSimulatedSchedulingScoresResponse) Reset() {
	*x = ListSimulatedSchedulingScoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_scores_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSimulatedSchedulingScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSimulatedSchedulingScoresResponse) ProtoMessage() {}

func (x *ListSimulatedSchedulingScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_scores_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSimulatedSchedulingScoresResponse.ProtoReflect.Descriptor instead.
func (*ListSimulatedSchedulingScoresResponse) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_scores_services_proto_rawDescGZIP(), []int{1}
}

func (x *ListSimulatedSchedulingScoresResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListSimulatedSchedulingScoresResponse) GetScores() []*SimulatedSchedulingScore {
	if x != nil {
		return x.Scores
	}
	return nil
}

// Represents a request for adding scheduling scores produced by ai engine
type CreateSimulatedSchedulingScoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*SimulatedSchedulingScore `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *CreateSimulatedSchedulingScoresRequest) Reset() {
	*x = CreateSimulatedSchedulingScoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_scores_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSimulatedSchedulingScoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSimulatedSchedulingScoresRequest) ProtoMessage() {}

func (x *CreateSimulatedSchedulingScoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_scores_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSimulatedSchedulingScoresRequest.ProtoReflect.Descriptor instead.
func (*CreateSimulatedSchedulingScoresRequest) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_scores_services_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSimulatedSchedulingScoresRequest) GetScores() []*SimulatedSchedulingScore {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_prophetstor_api_datahub_scores_services_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_scores_services_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa5, 0x01, 0x0a, 0x25, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x26,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x24, 0x5a, 0x22, 0x70, 0x72, 0x6f, 0x70,
	0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_scores_services_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_scores_services_proto_rawDescData = file_prophetstor_api_datahub_scores_services_proto_rawDesc
)

func file_prophetstor_api_datahub_scores_services_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_scores_services_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_scores_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_scores_services_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_scores_services_proto_rawDescData
}

var file_prophetstor_api_datahub_scores_services_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prophetstor_api_datahub_scores_services_proto_goTypes = []interface{}{
	(*ListSimulatedSchedulingScoresRequest)(nil),   // 0: prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresRequest
	(*ListSimulatedSchedulingScoresResponse)(nil),  // 1: prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresResponse
	(*CreateSimulatedSchedulingScoresRequest)(nil), // 2: prophetstor.api.datahub.scores.CreateSimulatedSchedulingScoresRequest
	(*common.QueryCondition)(nil),                  // 3: prophetstor.api.datahub.common.QueryCondition
	(*status.Status)(nil),                          // 4: google.rpc.Status
	(*SimulatedSchedulingScore)(nil),               // 5: prophetstor.api.datahub.scores.SimulatedSchedulingScore
}
var file_prophetstor_api_datahub_scores_services_proto_depIdxs = []int32{
	3, // 0: prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresRequest.query_condition:type_name -> prophetstor.api.datahub.common.QueryCondition
	4, // 1: prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresResponse.status:type_name -> google.rpc.Status
	5, // 2: prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresResponse.scores:type_name -> prophetstor.api.datahub.scores.SimulatedSchedulingScore
	5, // 3: prophetstor.api.datahub.scores.CreateSimulatedSchedulingScoresRequest.scores:type_name -> prophetstor.api.datahub.scores.SimulatedSchedulingScore
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_scores_services_proto_init() }
func file_prophetstor_api_datahub_scores_services_proto_init() {
	if File_prophetstor_api_datahub_scores_services_proto != nil {
		return
	}
	file_prophetstor_api_datahub_scores_scores_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_scores_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSimulatedSchedulingScoresRequest); i {
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
		file_prophetstor_api_datahub_scores_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSimulatedSchedulingScoresResponse); i {
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
		file_prophetstor_api_datahub_scores_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSimulatedSchedulingScoresRequest); i {
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
			RawDescriptor: file_prophetstor_api_datahub_scores_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_scores_services_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_scores_services_proto_depIdxs,
		MessageInfos:      file_prophetstor_api_datahub_scores_services_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_scores_services_proto = out.File
	file_prophetstor_api_datahub_scores_services_proto_rawDesc = nil
	file_prophetstor_api_datahub_scores_services_proto_goTypes = nil
	file_prophetstor_api_datahub_scores_services_proto_depIdxs = nil
}
