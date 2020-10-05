# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/plannings/types.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prophetstor.api.datahub.common import metrics_pb2 as prophetstor_dot_api_dot_datahub_dot_common_dot_metrics__pb2
from prophetstor.api.datahub.common import rawdata_pb2 as prophetstor_dot_api_dot_datahub_dot_common_dot_rawdata__pb2
from prophetstor.api.datahub.common import types_pb2 as prophetstor_dot_api_dot_datahub_dot_common_dot_types__pb2
from prophetstor.api.datahub.schemas import types_pb2 as prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/plannings/types.proto',
  package='prophetstor.api.datahub.plannings',
  syntax='proto3',
  serialized_options=b'Z%prophetstor.com/api/datahub/plannings',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n-prophetstor/api/datahub/plannings/types.proto\x12!prophetstor.api.datahub.plannings\x1a,prophetstor/api/datahub/common/metrics.proto\x1a,prophetstor/api/datahub/common/rawdata.proto\x1a*prophetstor/api/datahub/common/types.proto\x1a+prophetstor/api/datahub/schemas/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe7\x02\n\x16\x43ontrollerPlanningSpec\x12\x18\n\x10\x63urrent_replicas\x18\x01 \x01(\x05\x12\x18\n\x10\x64\x65sired_replicas\x18\x02 \x01(\x05\x12(\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x63reate_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1c\n\x14\x63urrent_cpu_requests\x18\x05 \x01(\x01\x12\x1c\n\x14\x63urrent_mem_requests\x18\x06 \x01(\x01\x12\x1a\n\x12\x63urrent_cpu_limits\x18\x07 \x01(\x01\x12\x1a\n\x12\x63urrent_mem_limits\x18\x08 \x01(\x01\x12\x1a\n\x12\x64\x65sired_cpu_limits\x18\t \x01(\x01\x12\x1a\n\x12\x64\x65sired_mem_limits\x18\n \x01(\x01\x12\x12\n\ntotal_cost\x18\x0b \x01(\x01\"\xaa\x01\n\x19\x43ontrollerPlanningSpecK8s\x12\x18\n\x10\x63urrent_replicas\x18\x01 \x01(\x05\x12\x18\n\x10\x64\x65sired_replicas\x18\x02 \x01(\x05\x12(\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x63reate_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x97\x01\n\x0bRawPlanning\x12@\n\x0bschema_meta\x18\x01 \x01(\x0b\x32+.prophetstor.api.datahub.schemas.SchemaMeta\x12\x46\n\rplanning_data\x18\x02 \x03(\x0b\x32/.prophetstor.api.datahub.plannings.PlanningData\"\xd3\x01\n\x0cPlanningData\x12?\n\x0bmetric_type\x18\x01 \x01(\x0e\x32*.prophetstor.api.datahub.common.MetricType\x12\x45\n\x0eresource_quota\x18\x02 \x01(\x0e\x32-.prophetstor.api.datahub.common.ResourceQuota\x12;\n\tread_data\x18\x03 \x01(\x0b\x32(.prophetstor.api.datahub.common.ReadData*H\n\x0cPlanningType\x12\x10\n\x0cPT_UNDEFINED\x10\x00\x12\x15\n\x11PT_RECOMMENDATION\x10\x01\x12\x0f\n\x0bPT_PLANNING\x10\x02*K\n\x16\x43ontrollerPlanningType\x12\x11\n\rCPT_UNDEFINED\x10\x00\x12\x11\n\rCPT_PRIMITIVE\x10\x01\x12\x0b\n\x07\x43PT_K8S\x10\x02\x42\'Z%prophetstor.com/api/datahub/planningsb\x06proto3'
  ,
  dependencies=[prophetstor_dot_api_dot_datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,prophetstor_dot_api_dot_datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,prophetstor_dot_api_dot_datahub_dot_common_dot_types__pb2.DESCRIPTOR,prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_PLANNINGTYPE = _descriptor.EnumDescriptor(
  name='PlanningType',
  full_name='prophetstor.api.datahub.plannings.PlanningType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PT_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PT_RECOMMENDATION', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PT_PLANNING', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1201,
  serialized_end=1273,
)
_sym_db.RegisterEnumDescriptor(_PLANNINGTYPE)

PlanningType = enum_type_wrapper.EnumTypeWrapper(_PLANNINGTYPE)
_CONTROLLERPLANNINGTYPE = _descriptor.EnumDescriptor(
  name='ControllerPlanningType',
  full_name='prophetstor.api.datahub.plannings.ControllerPlanningType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CPT_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='CPT_PRIMITIVE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='CPT_K8S', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1275,
  serialized_end=1350,
)
_sym_db.RegisterEnumDescriptor(_CONTROLLERPLANNINGTYPE)

ControllerPlanningType = enum_type_wrapper.EnumTypeWrapper(_CONTROLLERPLANNINGTYPE)
PT_UNDEFINED = 0
PT_RECOMMENDATION = 1
PT_PLANNING = 2
CPT_UNDEFINED = 0
CPT_PRIMITIVE = 1
CPT_K8S = 2



_CONTROLLERPLANNINGSPEC = _descriptor.Descriptor(
  name='ControllerPlanningSpec',
  full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='current_replicas', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.current_replicas', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desired_replicas', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.desired_replicas', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='time', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.create_time', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='current_cpu_requests', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.current_cpu_requests', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='current_mem_requests', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.current_mem_requests', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='current_cpu_limits', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.current_cpu_limits', index=6,
      number=7, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='current_mem_limits', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.current_mem_limits', index=7,
      number=8, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desired_cpu_limits', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.desired_cpu_limits', index=8,
      number=9, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desired_mem_limits', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.desired_mem_limits', index=9,
      number=10, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total_cost', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpec.total_cost', index=10,
      number=11, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=299,
  serialized_end=658,
)


_CONTROLLERPLANNINGSPECK8S = _descriptor.Descriptor(
  name='ControllerPlanningSpecK8s',
  full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpecK8s',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='current_replicas', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpecK8s.current_replicas', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desired_replicas', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpecK8s.desired_replicas', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='time', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpecK8s.time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='prophetstor.api.datahub.plannings.ControllerPlanningSpecK8s.create_time', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=661,
  serialized_end=831,
)


_RAWPLANNING = _descriptor.Descriptor(
  name='RawPlanning',
  full_name='prophetstor.api.datahub.plannings.RawPlanning',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.api.datahub.plannings.RawPlanning.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='planning_data', full_name='prophetstor.api.datahub.plannings.RawPlanning.planning_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=834,
  serialized_end=985,
)


_PLANNINGDATA = _descriptor.Descriptor(
  name='PlanningData',
  full_name='prophetstor.api.datahub.plannings.PlanningData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='prophetstor.api.datahub.plannings.PlanningData.metric_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_quota', full_name='prophetstor.api.datahub.plannings.PlanningData.resource_quota', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='read_data', full_name='prophetstor.api.datahub.plannings.PlanningData.read_data', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=988,
  serialized_end=1199,
)

_CONTROLLERPLANNINGSPEC.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTROLLERPLANNINGSPEC.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTROLLERPLANNINGSPECK8S.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTROLLERPLANNINGSPECK8S.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_RAWPLANNING.fields_by_name['schema_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_RAWPLANNING.fields_by_name['planning_data'].message_type = _PLANNINGDATA
_PLANNINGDATA.fields_by_name['metric_type'].enum_type = prophetstor_dot_api_dot_datahub_dot_common_dot_metrics__pb2._METRICTYPE
_PLANNINGDATA.fields_by_name['resource_quota'].enum_type = prophetstor_dot_api_dot_datahub_dot_common_dot_types__pb2._RESOURCEQUOTA
_PLANNINGDATA.fields_by_name['read_data'].message_type = prophetstor_dot_api_dot_datahub_dot_common_dot_rawdata__pb2._READDATA
DESCRIPTOR.message_types_by_name['ControllerPlanningSpec'] = _CONTROLLERPLANNINGSPEC
DESCRIPTOR.message_types_by_name['ControllerPlanningSpecK8s'] = _CONTROLLERPLANNINGSPECK8S
DESCRIPTOR.message_types_by_name['RawPlanning'] = _RAWPLANNING
DESCRIPTOR.message_types_by_name['PlanningData'] = _PLANNINGDATA
DESCRIPTOR.enum_types_by_name['PlanningType'] = _PLANNINGTYPE
DESCRIPTOR.enum_types_by_name['ControllerPlanningType'] = _CONTROLLERPLANNINGTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ControllerPlanningSpec = _reflection.GeneratedProtocolMessageType('ControllerPlanningSpec', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLERPLANNINGSPEC,
  '__module__' : 'prophetstor.api.datahub.plannings.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.plannings.ControllerPlanningSpec)
  })
_sym_db.RegisterMessage(ControllerPlanningSpec)

ControllerPlanningSpecK8s = _reflection.GeneratedProtocolMessageType('ControllerPlanningSpecK8s', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLERPLANNINGSPECK8S,
  '__module__' : 'prophetstor.api.datahub.plannings.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.plannings.ControllerPlanningSpecK8s)
  })
_sym_db.RegisterMessage(ControllerPlanningSpecK8s)

RawPlanning = _reflection.GeneratedProtocolMessageType('RawPlanning', (_message.Message,), {
  'DESCRIPTOR' : _RAWPLANNING,
  '__module__' : 'prophetstor.api.datahub.plannings.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.plannings.RawPlanning)
  })
_sym_db.RegisterMessage(RawPlanning)

PlanningData = _reflection.GeneratedProtocolMessageType('PlanningData', (_message.Message,), {
  'DESCRIPTOR' : _PLANNINGDATA,
  '__module__' : 'prophetstor.api.datahub.plannings.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.plannings.PlanningData)
  })
_sym_db.RegisterMessage(PlanningData)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
