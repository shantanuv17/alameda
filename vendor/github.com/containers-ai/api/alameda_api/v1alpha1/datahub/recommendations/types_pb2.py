# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/recommendations/types.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import metrics_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2
from alameda_api.v1alpha1.datahub.common import rawdata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2
from alameda_api.v1alpha1.datahub.common import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2
from alameda_api.v1alpha1.datahub.schemas import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/recommendations/types.proto',
  package='containersai.alameda.v1alpha1.datahub.recommendations',
  syntax='proto3',
  serialized_options=_b('ZIgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations'),
  serialized_pb=_b('\n8alameda_api/v1alpha1/datahub/recommendations/types.proto\x12\x35\x63ontainersai.alameda.v1alpha1.datahub.recommendations\x1a\x31\x61lameda_api/v1alpha1/datahub/common/metrics.proto\x1a\x31\x61lameda_api/v1alpha1/datahub/common/rawdata.proto\x1a/alameda_api/v1alpha1/datahub/common/types.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/schemas/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xea\x02\n\x19\x43ontrollerRecommendedSpec\x12\x18\n\x10\x63urrent_replicas\x18\x01 \x01(\x05\x12\x18\n\x10\x64\x65sired_replicas\x18\x02 \x01(\x05\x12(\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x63reate_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1c\n\x14\x63urrent_cpu_requests\x18\x05 \x01(\x01\x12\x1c\n\x14\x63urrent_mem_requests\x18\x06 \x01(\x01\x12\x1a\n\x12\x63urrent_cpu_limits\x18\x07 \x01(\x01\x12\x1a\n\x12\x63urrent_mem_limits\x18\x08 \x01(\x01\x12\x1a\n\x12\x64\x65sired_cpu_limits\x18\t \x01(\x01\x12\x1a\n\x12\x64\x65sired_mem_limits\x18\n \x01(\x01\x12\x12\n\ntotal_cost\x18\x0b \x01(\x01\"\xad\x01\n\x1c\x43ontrollerRecommendedSpecK8s\x12\x18\n\x10\x63urrent_replicas\x18\x01 \x01(\x05\x12\x18\n\x10\x64\x65sired_replicas\x18\x02 \x01(\x05\x12(\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x63reate_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xc8\x01\n\x0eRecommendation\x12N\n\x0bschema_meta\x18\x01 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta\x12\x66\n\x13recommendation_data\x18\x02 \x03(\x0b\x32I.containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData\"\x83\x02\n\x12RecommendationData\x12M\n\x0bmetric_type\x18\x01 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.MetricType\x12S\n\x0eresource_quota\x18\x02 \x01(\x0e\x32;.containersai.alameda.v1alpha1.datahub.common.ResourceQuota\x12I\n\tread_data\x18\x03 \x01(\x0b\x32\x36.containersai.alameda.v1alpha1.datahub.common.ReadData*F\n\x19\x43ontrollerRecommendedType\x12\x11\n\rCRT_UNDEFINED\x10\x00\x12\r\n\tPRIMITIVE\x10\x01\x12\x07\n\x03K8S\x10\x02\x42KZIgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendationsb\x06proto3')
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_CONTROLLERRECOMMENDEDTYPE = _descriptor.EnumDescriptor(
  name='ControllerRecommendedType',
  full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CRT_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIMITIVE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='K8S', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1355,
  serialized_end=1425,
)
_sym_db.RegisterEnumDescriptor(_CONTROLLERRECOMMENDEDTYPE)

ControllerRecommendedType = enum_type_wrapper.EnumTypeWrapper(_CONTROLLERRECOMMENDEDTYPE)
CRT_UNDEFINED = 0
PRIMITIVE = 1
K8S = 2



_CONTROLLERRECOMMENDEDSPEC = _descriptor.Descriptor(
  name='ControllerRecommendedSpec',
  full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='current_replicas', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.current_replicas', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desired_replicas', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.desired_replicas', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='time', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.create_time', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='current_cpu_requests', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.current_cpu_requests', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='current_mem_requests', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.current_mem_requests', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='current_cpu_limits', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.current_cpu_limits', index=6,
      number=7, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='current_mem_limits', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.current_mem_limits', index=7,
      number=8, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desired_cpu_limits', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.desired_cpu_limits', index=8,
      number=9, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desired_mem_limits', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.desired_mem_limits', index=9,
      number=10, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='total_cost', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec.total_cost', index=10,
      number=11, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=350,
  serialized_end=712,
)


_CONTROLLERRECOMMENDEDSPECK8S = _descriptor.Descriptor(
  name='ControllerRecommendedSpecK8s',
  full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='current_replicas', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s.current_replicas', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desired_replicas', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s.desired_replicas', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='time', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s.time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s.create_time', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=715,
  serialized_end=888,
)


_RECOMMENDATION = _descriptor.Descriptor(
  name='Recommendation',
  full_name='containersai.alameda.v1alpha1.datahub.recommendations.Recommendation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='containersai.alameda.v1alpha1.datahub.recommendations.Recommendation.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='recommendation_data', full_name='containersai.alameda.v1alpha1.datahub.recommendations.Recommendation.recommendation_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=891,
  serialized_end=1091,
)


_RECOMMENDATIONDATA = _descriptor.Descriptor(
  name='RecommendationData',
  full_name='containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData.metric_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_quota', full_name='containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData.resource_quota', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='read_data', full_name='containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData.read_data', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1094,
  serialized_end=1353,
)

_CONTROLLERRECOMMENDEDSPEC.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTROLLERRECOMMENDEDSPEC.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTROLLERRECOMMENDEDSPECK8S.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTROLLERRECOMMENDEDSPECK8S.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_RECOMMENDATION.fields_by_name['schema_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_RECOMMENDATION.fields_by_name['recommendation_data'].message_type = _RECOMMENDATIONDATA
_RECOMMENDATIONDATA.fields_by_name['metric_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2._METRICTYPE
_RECOMMENDATIONDATA.fields_by_name['resource_quota'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._RESOURCEQUOTA
_RECOMMENDATIONDATA.fields_by_name['read_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2._READDATA
DESCRIPTOR.message_types_by_name['ControllerRecommendedSpec'] = _CONTROLLERRECOMMENDEDSPEC
DESCRIPTOR.message_types_by_name['ControllerRecommendedSpecK8s'] = _CONTROLLERRECOMMENDEDSPECK8S
DESCRIPTOR.message_types_by_name['Recommendation'] = _RECOMMENDATION
DESCRIPTOR.message_types_by_name['RecommendationData'] = _RECOMMENDATIONDATA
DESCRIPTOR.enum_types_by_name['ControllerRecommendedType'] = _CONTROLLERRECOMMENDEDTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ControllerRecommendedSpec = _reflection.GeneratedProtocolMessageType('ControllerRecommendedSpec', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLERRECOMMENDEDSPEC,
  '__module__' : 'alameda_api.v1alpha1.datahub.recommendations.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec)
  })
_sym_db.RegisterMessage(ControllerRecommendedSpec)

ControllerRecommendedSpecK8s = _reflection.GeneratedProtocolMessageType('ControllerRecommendedSpecK8s', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLERRECOMMENDEDSPECK8S,
  '__module__' : 'alameda_api.v1alpha1.datahub.recommendations.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s)
  })
_sym_db.RegisterMessage(ControllerRecommendedSpecK8s)

Recommendation = _reflection.GeneratedProtocolMessageType('Recommendation', (_message.Message,), {
  'DESCRIPTOR' : _RECOMMENDATION,
  '__module__' : 'alameda_api.v1alpha1.datahub.recommendations.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.recommendations.Recommendation)
  })
_sym_db.RegisterMessage(Recommendation)

RecommendationData = _reflection.GeneratedProtocolMessageType('RecommendationData', (_message.Message,), {
  'DESCRIPTOR' : _RECOMMENDATIONDATA,
  '__module__' : 'alameda_api.v1alpha1.datahub.recommendations.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData)
  })
_sym_db.RegisterMessage(RecommendationData)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
