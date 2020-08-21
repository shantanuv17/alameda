# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/common/queries.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/common/queries.proto',
  package='containersai.alameda.v1alpha1.datahub.common',
  syntax='proto3',
  serialized_options=b'Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n1alameda_api/v1alpha1/datahub/common/queries.proto\x12,containersai.alameda.v1alpha1.datahub.common\x1a/alameda_api/v1alpha1/datahub/common/types.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd9\x02\n\tTimeRange\x12.\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\x04step\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x64\n\x11\x61ggregateFunction\x18\x04 \x01(\x0e\x32I.containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction\x12.\n\napply_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"/\n\x11\x41ggregateFunction\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03MAX\x10\x01\x12\x07\n\x03\x41VG\x10\x02\"\x83\x01\n\tCondition\x12\x0c\n\x04keys\x18\x01 \x03(\t\x12\x0e\n\x06values\x18\x02 \x03(\t\x12\x11\n\toperators\x18\x03 \x03(\t\x12\x45\n\x05types\x18\x04 \x03(\x0e\x32\x36.containersai.alameda.v1alpha1.datahub.common.DataType\"\xbc\x01\n\x08\x46unction\x12H\n\x04type\x18\x01 \x01(\x0e\x32:.containersai.alameda.v1alpha1.datahub.common.FunctionType\x12\x0e\n\x06\x66ields\x18\x02 \x03(\t\x12\x0c\n\x04tags\x18\x03 \x03(\t\x12\x0e\n\x06target\x18\x04 \x01(\t\x12\x1a\n\x12regular_expression\x18\x05 \x01(\t\x12\x0c\n\x04unit\x18\x06 \x01(\t\x12\x0e\n\x06number\x18\x07 \x01(\x03\"\x89\x01\n\x04Into\x12\x10\n\x08\x64\x61tabase\x18\x01 \x01(\t\x12\x18\n\x10retention_policy\x18\x02 \x01(\t\x12\x13\n\x0bmeasurement\x18\x03 \x01(\t\x12#\n\x1bis_default_retention_policy\x18\x04 \x01(\x08\x12\x1b\n\x13is_all_measurements\x18\x05 \x01(\x08\"\xfa\x03\n\x0eQueryCondition\x12K\n\ntime_range\x18\x01 \x01(\x0b\x32\x37.containersai.alameda.v1alpha1.datahub.common.TimeRange\x12Q\n\x05order\x18\x02 \x01(\x0e\x32\x42.containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order\x12H\n\x08\x66unction\x18\x03 \x01(\x0b\x32\x36.containersai.alameda.v1alpha1.datahub.common.Function\x12@\n\x04into\x18\x04 \x01(\x0b\x32\x32.containersai.alameda.v1alpha1.datahub.common.Into\x12\x14\n\x0cwhere_clause\x18\x05 \x01(\t\x12P\n\x0fwhere_condition\x18\x06 \x03(\x0b\x32\x37.containersai.alameda.v1alpha1.datahub.common.Condition\x12\x0f\n\x07selects\x18\x07 \x03(\t\x12\x0e\n\x06groups\x18\x08 \x03(\t\x12\r\n\x05limit\x18\t \x01(\x04\"$\n\x05Order\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41SC\x10\x01\x12\x08\n\x04\x44\x45SC\x10\x02\x42\x42Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/commonb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])



_TIMERANGE_AGGREGATEFUNCTION = _descriptor.EnumDescriptor(
  name='AggregateFunction',
  full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='MAX', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='AVG', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=512,
  serialized_end=559,
)
_sym_db.RegisterEnumDescriptor(_TIMERANGE_AGGREGATEFUNCTION)

_QUERYCONDITION_ORDER = _descriptor.EnumDescriptor(
  name='Order',
  full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ASC', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DESC', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1497,
  serialized_end=1533,
)
_sym_db.RegisterEnumDescriptor(_QUERYCONDITION_ORDER)


_TIMERANGE = _descriptor.Descriptor(
  name='TimeRange',
  full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.start_time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.end_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='step', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.step', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='aggregateFunction', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.aggregateFunction', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='apply_time', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.apply_time', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TIMERANGE_AGGREGATEFUNCTION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=214,
  serialized_end=559,
)


_CONDITION = _descriptor.Descriptor(
  name='Condition',
  full_name='containersai.alameda.v1alpha1.datahub.common.Condition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='keys', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.keys', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='values', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.values', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='operators', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.operators', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='types', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.types', index=3,
      number=4, type=14, cpp_type=8, label=3,
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
  serialized_start=562,
  serialized_end=693,
)


_FUNCTION = _descriptor.Descriptor(
  name='Function',
  full_name='containersai.alameda.v1alpha1.datahub.common.Function',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='containersai.alameda.v1alpha1.datahub.common.Function.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='fields', full_name='containersai.alameda.v1alpha1.datahub.common.Function.fields', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tags', full_name='containersai.alameda.v1alpha1.datahub.common.Function.tags', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='target', full_name='containersai.alameda.v1alpha1.datahub.common.Function.target', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='regular_expression', full_name='containersai.alameda.v1alpha1.datahub.common.Function.regular_expression', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='unit', full_name='containersai.alameda.v1alpha1.datahub.common.Function.unit', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='number', full_name='containersai.alameda.v1alpha1.datahub.common.Function.number', index=6,
      number=7, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=696,
  serialized_end=884,
)


_INTO = _descriptor.Descriptor(
  name='Into',
  full_name='containersai.alameda.v1alpha1.datahub.common.Into',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='database', full_name='containersai.alameda.v1alpha1.datahub.common.Into.database', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='retention_policy', full_name='containersai.alameda.v1alpha1.datahub.common.Into.retention_policy', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='measurement', full_name='containersai.alameda.v1alpha1.datahub.common.Into.measurement', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_default_retention_policy', full_name='containersai.alameda.v1alpha1.datahub.common.Into.is_default_retention_policy', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_all_measurements', full_name='containersai.alameda.v1alpha1.datahub.common.Into.is_all_measurements', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=887,
  serialized_end=1024,
)


_QUERYCONDITION = _descriptor.Descriptor(
  name='QueryCondition',
  full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='time_range', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.time_range', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='order', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.order', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='function', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.function', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='into', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.into', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='where_clause', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.where_clause', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='where_condition', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.where_condition', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='selects', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.selects', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='groups', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.groups', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limit', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.limit', index=8,
      number=9, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _QUERYCONDITION_ORDER,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1027,
  serialized_end=1533,
)

_TIMERANGE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['step'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_TIMERANGE.fields_by_name['aggregateFunction'].enum_type = _TIMERANGE_AGGREGATEFUNCTION
_TIMERANGE.fields_by_name['apply_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE_AGGREGATEFUNCTION.containing_type = _TIMERANGE
_CONDITION.fields_by_name['types'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._DATATYPE
_FUNCTION.fields_by_name['type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._FUNCTIONTYPE
_QUERYCONDITION.fields_by_name['time_range'].message_type = _TIMERANGE
_QUERYCONDITION.fields_by_name['order'].enum_type = _QUERYCONDITION_ORDER
_QUERYCONDITION.fields_by_name['function'].message_type = _FUNCTION
_QUERYCONDITION.fields_by_name['into'].message_type = _INTO
_QUERYCONDITION.fields_by_name['where_condition'].message_type = _CONDITION
_QUERYCONDITION_ORDER.containing_type = _QUERYCONDITION
DESCRIPTOR.message_types_by_name['TimeRange'] = _TIMERANGE
DESCRIPTOR.message_types_by_name['Condition'] = _CONDITION
DESCRIPTOR.message_types_by_name['Function'] = _FUNCTION
DESCRIPTOR.message_types_by_name['Into'] = _INTO
DESCRIPTOR.message_types_by_name['QueryCondition'] = _QUERYCONDITION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TimeRange = _reflection.GeneratedProtocolMessageType('TimeRange', (_message.Message,), {
  'DESCRIPTOR' : _TIMERANGE,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.TimeRange)
  })
_sym_db.RegisterMessage(TimeRange)

Condition = _reflection.GeneratedProtocolMessageType('Condition', (_message.Message,), {
  'DESCRIPTOR' : _CONDITION,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.Condition)
  })
_sym_db.RegisterMessage(Condition)

Function = _reflection.GeneratedProtocolMessageType('Function', (_message.Message,), {
  'DESCRIPTOR' : _FUNCTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.Function)
  })
_sym_db.RegisterMessage(Function)

Into = _reflection.GeneratedProtocolMessageType('Into', (_message.Message,), {
  'DESCRIPTOR' : _INTO,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.Into)
  })
_sym_db.RegisterMessage(Into)

QueryCondition = _reflection.GeneratedProtocolMessageType('QueryCondition', (_message.Message,), {
  'DESCRIPTOR' : _QUERYCONDITION,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.QueryCondition)
  })
_sym_db.RegisterMessage(QueryCondition)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
