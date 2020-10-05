# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/schemas/schemas.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prophetstor.api.datahub.schemas import types_pb2 as prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/schemas/schemas.proto',
  package='prophetstor.api.datahub.schemas',
  syntax='proto3',
  serialized_options=b'Z#prophetstor.com/api/datahub/schemas',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n-prophetstor/api/datahub/schemas/schemas.proto\x12\x1fprophetstor.api.datahub.schemas\x1a+prophetstor/api/datahub/schemas/types.proto\"\x8e\x01\n\x06Schema\x12@\n\x0bschema_meta\x18\x01 \x01(\x0b\x32+.prophetstor.api.datahub.schemas.SchemaMeta\x12\x42\n\x0cmeasurements\x18\x02 \x03(\x0b\x32,.prophetstor.api.datahub.schemas.MeasurementB%Z#prophetstor.com/api/datahub/schemasb\x06proto3'
  ,
  dependencies=[prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,])




_SCHEMA = _descriptor.Descriptor(
  name='Schema',
  full_name='prophetstor.api.datahub.schemas.Schema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.api.datahub.schemas.Schema.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='measurements', full_name='prophetstor.api.datahub.schemas.Schema.measurements', index=1,
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
  serialized_start=128,
  serialized_end=270,
)

_SCHEMA.fields_by_name['schema_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_SCHEMA.fields_by_name['measurements'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._MEASUREMENT
DESCRIPTOR.message_types_by_name['Schema'] = _SCHEMA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Schema = _reflection.GeneratedProtocolMessageType('Schema', (_message.Message,), {
  'DESCRIPTOR' : _SCHEMA,
  '__module__' : 'prophetstor.api.datahub.schemas.schemas_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.schemas.Schema)
  })
_sym_db.RegisterMessage(Schema)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
