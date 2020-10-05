# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/rawdata/types.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prophetstor.api.datahub.common import queries_pb2 as prophetstor_dot_api_dot_datahub_dot_common_dot_queries__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/rawdata/types.proto',
  package='prophetstor.api.datahub.rawdata',
  syntax='proto3',
  serialized_options=b'Z#prophetstor.com/api/datahub/rawdata',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n+prophetstor/api/datahub/rawdata/types.proto\x12\x1fprophetstor.api.datahub.rawdata\x1a,prophetstor/api/datahub/common/queries.proto\"\x7f\n\x05Query\x12\x10\n\x08\x64\x61tabase\x18\x01 \x01(\t\x12\r\n\x05table\x18\x02 \x01(\t\x12\x12\n\nexpression\x18\x03 \x01(\t\x12\x41\n\tcondition\x18\x04 \x01(\x0b\x32..prophetstor.api.datahub.common.QueryConditionB%Z#prophetstor.com/api/datahub/rawdatab\x06proto3'
  ,
  dependencies=[prophetstor_dot_api_dot_datahub_dot_common_dot_queries__pb2.DESCRIPTOR,])




_QUERY = _descriptor.Descriptor(
  name='Query',
  full_name='prophetstor.api.datahub.rawdata.Query',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='database', full_name='prophetstor.api.datahub.rawdata.Query.database', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='table', full_name='prophetstor.api.datahub.rawdata.Query.table', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='expression', full_name='prophetstor.api.datahub.rawdata.Query.expression', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='condition', full_name='prophetstor.api.datahub.rawdata.Query.condition', index=3,
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
  serialized_start=126,
  serialized_end=253,
)

_QUERY.fields_by_name['condition'].message_type = prophetstor_dot_api_dot_datahub_dot_common_dot_queries__pb2._QUERYCONDITION
DESCRIPTOR.message_types_by_name['Query'] = _QUERY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Query = _reflection.GeneratedProtocolMessageType('Query', (_message.Message,), {
  'DESCRIPTOR' : _QUERY,
  '__module__' : 'prophetstor.api.datahub.rawdata.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.rawdata.Query)
  })
_sym_db.RegisterMessage(Query)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
