# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/configs/service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prophetstor.api.datahub.resources import metadata_pb2 as prophetstor_dot_api_dot_datahub_dot_resources_dot_metadata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/configs/service.proto',
  package='prophetstor.api.datahub.configs',
  syntax='proto3',
  serialized_options=b'Z#prophetstor.com/api/datahub/configs',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n-prophetstor/api/datahub/configs/service.proto\x12\x1fprophetstor.api.datahub.configs\x1a\x30prophetstor/api/datahub/resources/metadata.proto\"6\n\x07Keycode\x12\x13\n\x0b\x63ode_number\x18\x01 \x01(\t\x12\x16\n\x0esignature_data\x18\x02 \x01(\t\"\x88\x01\n\x07Service\x12\x42\n\x0bobject_meta\x18\x01 \x01(\x0b\x32-.prophetstor.api.datahub.resources.ObjectMeta\x12\x39\n\x07keycode\x18\x02 \x01(\x0b\x32(.prophetstor.api.datahub.configs.KeycodeB%Z#prophetstor.com/api/datahub/configsb\x06proto3'
  ,
  dependencies=[prophetstor_dot_api_dot_datahub_dot_resources_dot_metadata__pb2.DESCRIPTOR,])




_KEYCODE = _descriptor.Descriptor(
  name='Keycode',
  full_name='prophetstor.api.datahub.configs.Keycode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code_number', full_name='prophetstor.api.datahub.configs.Keycode.code_number', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='signature_data', full_name='prophetstor.api.datahub.configs.Keycode.signature_data', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=132,
  serialized_end=186,
)


_SERVICE = _descriptor.Descriptor(
  name='Service',
  full_name='prophetstor.api.datahub.configs.Service',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='prophetstor.api.datahub.configs.Service.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='keycode', full_name='prophetstor.api.datahub.configs.Service.keycode', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=189,
  serialized_end=325,
)

_SERVICE.fields_by_name['object_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_SERVICE.fields_by_name['keycode'].message_type = _KEYCODE
DESCRIPTOR.message_types_by_name['Keycode'] = _KEYCODE
DESCRIPTOR.message_types_by_name['Service'] = _SERVICE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Keycode = _reflection.GeneratedProtocolMessageType('Keycode', (_message.Message,), {
  'DESCRIPTOR' : _KEYCODE,
  '__module__' : 'prophetstor.api.datahub.configs.service_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.configs.Keycode)
  })
_sym_db.RegisterMessage(Keycode)

Service = _reflection.GeneratedProtocolMessageType('Service', (_message.Message,), {
  'DESCRIPTOR' : _SERVICE,
  '__module__' : 'prophetstor.api.datahub.configs.service_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.configs.Service)
  })
_sym_db.RegisterMessage(Service)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
