# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/schemas/types.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import metrics_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2
from alameda_api.v1alpha1.datahub.common import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/schemas/types.proto',
  package='containersai.alameda.v1alpha1.datahub.schemas',
  syntax='proto3',
  serialized_options=b'ZAgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas',
  serialized_pb=b'\n0alameda_api/v1alpha1/datahub/schemas/types.proto\x12-containersai.alameda.v1alpha1.datahub.schemas\x1a\x31\x61lameda_api/v1alpha1/datahub/common/metrics.proto\x1a/alameda_api/v1alpha1/datahub/common/types.proto\"q\n\nSchemaMeta\x12\x43\n\x05scope\x18\x01 \x01(\x0e\x32\x34.containersai.alameda.v1alpha1.datahub.schemas.Scope\x12\x10\n\x08\x63\x61tegory\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\"\xf1\x02\n\x0bMeasurement\x12\x0c\n\x04name\x18\x01 \x01(\t\x12M\n\x0bmetric_type\x18\x02 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.MetricType\x12Y\n\x11resource_boundary\x18\x03 \x01(\x0e\x32>.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary\x12S\n\x0eresource_quota\x18\x04 \x01(\x0e\x32;.containersai.alameda.v1alpha1.datahub.common.ResourceQuota\x12\r\n\x05is_ts\x18\x05 \x01(\x08\x12\x46\n\x07\x63olumns\x18\x06 \x03(\x0b\x32\x35.containersai.alameda.v1alpha1.datahub.schemas.Column\"\xc2\x01\n\x06\x43olumn\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08required\x18\x02 \x01(\x08\x12M\n\x0b\x63olumn_type\x18\x03 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.ColumnType\x12I\n\tdata_type\x18\x04 \x01(\x0e\x32\x36.containersai.alameda.v1alpha1.datahub.common.DataType*\xea\x01\n\x05Scope\x12\x13\n\x0fSCOPE_UNDEFINED\x10\x00\x12\x15\n\x11SCOPE_APPLICATION\x10\x01\x12\x10\n\x0cSCOPE_CONFIG\x10\x02\x12\x13\n\x0fSCOPE_FEDEMETER\x10\x03\x12\x12\n\x0eSCOPE_METERING\x10\x04\x12\x10\n\x0cSCOPE_METRIC\x10\x05\x12\x12\n\x0eSCOPE_PLANNING\x10\x06\x12\x14\n\x10SCOPE_PREDICTION\x10\x07\x12\x18\n\x14SCOPE_RECOMMENDATION\x10\x08\x12\x12\n\x0eSCOPE_RESOURCE\x10\t\x12\x10\n\x0cSCOPE_TARGET\x10\nBCZAgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemasb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2.DESCRIPTOR,])

_SCOPE = _descriptor.EnumDescriptor(
  name='Scope',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.Scope',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SCOPE_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_APPLICATION', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_CONFIG', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_FEDEMETER', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_METERING', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_METRIC', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_PLANNING', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_PREDICTION', index=7, number=7,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_RECOMMENDATION', index=8, number=8,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_RESOURCE', index=9, number=9,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCOPE_TARGET', index=10, number=10,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=884,
  serialized_end=1118,
)
_sym_db.RegisterEnumDescriptor(_SCOPE)

Scope = enum_type_wrapper.EnumTypeWrapper(_SCOPE)
SCOPE_UNDEFINED = 0
SCOPE_APPLICATION = 1
SCOPE_CONFIG = 2
SCOPE_FEDEMETER = 3
SCOPE_METERING = 4
SCOPE_METRIC = 5
SCOPE_PLANNING = 6
SCOPE_PREDICTION = 7
SCOPE_RECOMMENDATION = 8
SCOPE_RESOURCE = 9
SCOPE_TARGET = 10



_SCHEMAMETA = _descriptor.Descriptor(
  name='SchemaMeta',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='scope', full_name='containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.scope', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='category', full_name='containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.category', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.type', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=199,
  serialized_end=312,
)


_MEASUREMENT = _descriptor.Descriptor(
  name='Measurement',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement.metric_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_boundary', full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement.resource_boundary', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_quota', full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement.resource_quota', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_ts', full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement.is_ts', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='columns', full_name='containersai.alameda.v1alpha1.datahub.schemas.Measurement.columns', index=5,
      number=6, type=11, cpp_type=10, label=3,
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
  serialized_start=315,
  serialized_end=684,
)


_COLUMN = _descriptor.Descriptor(
  name='Column',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.Column',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.alameda.v1alpha1.datahub.schemas.Column.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='required', full_name='containersai.alameda.v1alpha1.datahub.schemas.Column.required', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='column_type', full_name='containersai.alameda.v1alpha1.datahub.schemas.Column.column_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_type', full_name='containersai.alameda.v1alpha1.datahub.schemas.Column.data_type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=687,
  serialized_end=881,
)

_SCHEMAMETA.fields_by_name['scope'].enum_type = _SCOPE
_MEASUREMENT.fields_by_name['metric_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2._METRICTYPE
_MEASUREMENT.fields_by_name['resource_boundary'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._RESOURCEBOUNDARY
_MEASUREMENT.fields_by_name['resource_quota'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._RESOURCEQUOTA
_MEASUREMENT.fields_by_name['columns'].message_type = _COLUMN
_COLUMN.fields_by_name['column_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._COLUMNTYPE
_COLUMN.fields_by_name['data_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._DATATYPE
DESCRIPTOR.message_types_by_name['SchemaMeta'] = _SCHEMAMETA
DESCRIPTOR.message_types_by_name['Measurement'] = _MEASUREMENT
DESCRIPTOR.message_types_by_name['Column'] = _COLUMN
DESCRIPTOR.enum_types_by_name['Scope'] = _SCOPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SchemaMeta = _reflection.GeneratedProtocolMessageType('SchemaMeta', (_message.Message,), {
  'DESCRIPTOR' : _SCHEMAMETA,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta)
  })
_sym_db.RegisterMessage(SchemaMeta)

Measurement = _reflection.GeneratedProtocolMessageType('Measurement', (_message.Message,), {
  'DESCRIPTOR' : _MEASUREMENT,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.Measurement)
  })
_sym_db.RegisterMessage(Measurement)

Column = _reflection.GeneratedProtocolMessageType('Column', (_message.Message,), {
  'DESCRIPTOR' : _COLUMN,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.Column)
  })
_sym_db.RegisterMessage(Column)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
