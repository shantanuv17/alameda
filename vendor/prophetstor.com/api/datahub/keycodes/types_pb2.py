# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/keycodes/types.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/keycodes/types.proto',
  package='prophetstor.api.datahub.keycodes',
  syntax='proto3',
  serialized_options=b'Z$prophetstor.com/api/datahub/keycodes',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n,prophetstor/api/datahub/keycodes/types.proto\x12 prophetstor.api.datahub.keycodes\"7\n\x08\x43\x61pacity\x12\r\n\x05users\x18\x01 \x01(\x05\x12\r\n\x05hosts\x18\x02 \x01(\x05\x12\r\n\x05\x64isks\x18\x03 \x01(\x05\"7\n\rFunctionality\x12\x14\n\x0c\x64isk_prophet\x18\x01 \x01(\x08\x12\x10\n\x08workload\x18\x02 \x01(\x08\"/\n\tRetention\x12\x13\n\x0bvalid_month\x18\x01 \x01(\x05\x12\r\n\x05years\x18\x02 \x01(\x05\"\x12\n\x10ServiceAgreementB&Z$prophetstor.com/api/datahub/keycodesb\x06proto3'
)




_CAPACITY = _descriptor.Descriptor(
  name='Capacity',
  full_name='prophetstor.api.datahub.keycodes.Capacity',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='users', full_name='prophetstor.api.datahub.keycodes.Capacity.users', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='hosts', full_name='prophetstor.api.datahub.keycodes.Capacity.hosts', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='disks', full_name='prophetstor.api.datahub.keycodes.Capacity.disks', index=2,
      number=3, type=5, cpp_type=1, label=1,
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
  serialized_start=82,
  serialized_end=137,
)


_FUNCTIONALITY = _descriptor.Descriptor(
  name='Functionality',
  full_name='prophetstor.api.datahub.keycodes.Functionality',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='disk_prophet', full_name='prophetstor.api.datahub.keycodes.Functionality.disk_prophet', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='workload', full_name='prophetstor.api.datahub.keycodes.Functionality.workload', index=1,
      number=2, type=8, cpp_type=7, label=1,
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
  serialized_start=139,
  serialized_end=194,
)


_RETENTION = _descriptor.Descriptor(
  name='Retention',
  full_name='prophetstor.api.datahub.keycodes.Retention',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='valid_month', full_name='prophetstor.api.datahub.keycodes.Retention.valid_month', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='years', full_name='prophetstor.api.datahub.keycodes.Retention.years', index=1,
      number=2, type=5, cpp_type=1, label=1,
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
  serialized_start=196,
  serialized_end=243,
)


_SERVICEAGREEMENT = _descriptor.Descriptor(
  name='ServiceAgreement',
  full_name='prophetstor.api.datahub.keycodes.ServiceAgreement',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=245,
  serialized_end=263,
)

DESCRIPTOR.message_types_by_name['Capacity'] = _CAPACITY
DESCRIPTOR.message_types_by_name['Functionality'] = _FUNCTIONALITY
DESCRIPTOR.message_types_by_name['Retention'] = _RETENTION
DESCRIPTOR.message_types_by_name['ServiceAgreement'] = _SERVICEAGREEMENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Capacity = _reflection.GeneratedProtocolMessageType('Capacity', (_message.Message,), {
  'DESCRIPTOR' : _CAPACITY,
  '__module__' : 'prophetstor.api.datahub.keycodes.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.Capacity)
  })
_sym_db.RegisterMessage(Capacity)

Functionality = _reflection.GeneratedProtocolMessageType('Functionality', (_message.Message,), {
  'DESCRIPTOR' : _FUNCTIONALITY,
  '__module__' : 'prophetstor.api.datahub.keycodes.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.Functionality)
  })
_sym_db.RegisterMessage(Functionality)

Retention = _reflection.GeneratedProtocolMessageType('Retention', (_message.Message,), {
  'DESCRIPTOR' : _RETENTION,
  '__module__' : 'prophetstor.api.datahub.keycodes.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.Retention)
  })
_sym_db.RegisterMessage(Retention)

ServiceAgreement = _reflection.GeneratedProtocolMessageType('ServiceAgreement', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEAGREEMENT,
  '__module__' : 'prophetstor.api.datahub.keycodes.types_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.ServiceAgreement)
  })
_sym_db.RegisterMessage(ServiceAgreement)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
