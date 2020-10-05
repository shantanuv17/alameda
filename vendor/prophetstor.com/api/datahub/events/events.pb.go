// This file has messages related to system events

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: prophetstor/api/datahub/events/events.proto

package events

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Id        string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId string               `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Source    *EventSource         `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Type      EventType            `protobuf:"varint,5,opt,name=type,proto3,enum=prophetstor.api.datahub.events.EventType" json:"type,omitempty"`
	Version   EventVersion         `protobuf:"varint,6,opt,name=version,proto3,enum=prophetstor.api.datahub.events.EventVersion" json:"version,omitempty"`
	Level     EventLevel           `protobuf:"varint,7,opt,name=level,proto3,enum=prophetstor.api.datahub.events.EventLevel" json:"level,omitempty"`
	Subject   *K8SObjectReference  `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject,omitempty"`
	Message   string               `protobuf:"bytes,9,opt,name=message,proto3" json:"message,omitempty"`
	Data      string               `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prophetstor_api_datahub_events_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_prophetstor_api_datahub_events_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_prophetstor_api_datahub_events_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Event) GetSource() *EventSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_TYPE_UNDEFINED
}

func (x *Event) GetVersion() EventVersion {
	if x != nil {
		return x.Version
	}
	return EventVersion_EVENT_VERSION_UNDEFINED
}

func (x *Event) GetLevel() EventLevel {
	if x != nil {
		return x.Level
	}
	return EventLevel_EVENT_LEVEL_UNDEFINED
}

func (x *Event) GetSubject() *K8SObjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Event) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_prophetstor_api_datahub_events_events_proto protoreflect.FileDescriptor

var file_prophetstor_api_datahub_events_events_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x2a, 0x70,
	0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x4c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x4b, 0x38, 0x53, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x24, 0x5a,
	0x22, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prophetstor_api_datahub_events_events_proto_rawDescOnce sync.Once
	file_prophetstor_api_datahub_events_events_proto_rawDescData = file_prophetstor_api_datahub_events_events_proto_rawDesc
)

func file_prophetstor_api_datahub_events_events_proto_rawDescGZIP() []byte {
	file_prophetstor_api_datahub_events_events_proto_rawDescOnce.Do(func() {
		file_prophetstor_api_datahub_events_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_prophetstor_api_datahub_events_events_proto_rawDescData)
	})
	return file_prophetstor_api_datahub_events_events_proto_rawDescData
}

var file_prophetstor_api_datahub_events_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prophetstor_api_datahub_events_events_proto_goTypes = []interface{}{
	(*Event)(nil),               // 0: prophetstor.api.datahub.events.Event
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*EventSource)(nil),         // 2: prophetstor.api.datahub.events.EventSource
	(EventType)(0),              // 3: prophetstor.api.datahub.events.EventType
	(EventVersion)(0),           // 4: prophetstor.api.datahub.events.EventVersion
	(EventLevel)(0),             // 5: prophetstor.api.datahub.events.EventLevel
	(*K8SObjectReference)(nil),  // 6: prophetstor.api.datahub.events.K8SObjectReference
}
var file_prophetstor_api_datahub_events_events_proto_depIdxs = []int32{
	1, // 0: prophetstor.api.datahub.events.Event.time:type_name -> google.protobuf.Timestamp
	2, // 1: prophetstor.api.datahub.events.Event.source:type_name -> prophetstor.api.datahub.events.EventSource
	3, // 2: prophetstor.api.datahub.events.Event.type:type_name -> prophetstor.api.datahub.events.EventType
	4, // 3: prophetstor.api.datahub.events.Event.version:type_name -> prophetstor.api.datahub.events.EventVersion
	5, // 4: prophetstor.api.datahub.events.Event.level:type_name -> prophetstor.api.datahub.events.EventLevel
	6, // 5: prophetstor.api.datahub.events.Event.subject:type_name -> prophetstor.api.datahub.events.K8SObjectReference
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_prophetstor_api_datahub_events_events_proto_init() }
func file_prophetstor_api_datahub_events_events_proto_init() {
	if File_prophetstor_api_datahub_events_events_proto != nil {
		return
	}
	file_prophetstor_api_datahub_events_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prophetstor_api_datahub_events_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_prophetstor_api_datahub_events_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prophetstor_api_datahub_events_events_proto_goTypes,
		DependencyIndexes: file_prophetstor_api_datahub_events_events_proto_depIdxs,
		MessageInfos:      file_prophetstor_api_datahub_events_events_proto_msgTypes,
	}.Build()
	File_prophetstor_api_datahub_events_events_proto = out.File
	file_prophetstor_api_datahub_events_events_proto_rawDesc = nil
	file_prophetstor_api_datahub_events_events_proto_goTypes = nil
	file_prophetstor_api_datahub_events_events_proto_depIdxs = nil
}
