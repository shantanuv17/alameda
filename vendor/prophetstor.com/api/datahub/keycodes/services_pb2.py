# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/keycodes/services.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prophetstor.api.datahub.keycodes import keycodes_pb2 as prophetstor_dot_api_dot_datahub_dot_keycodes_dot_keycodes__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/keycodes/services.proto',
  package='prophetstor.api.datahub.keycodes',
  syntax='proto3',
  serialized_options=b'Z$prophetstor.com/api/datahub/keycodes',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n/prophetstor/api/datahub/keycodes/services.proto\x12 prophetstor.api.datahub.keycodes\x1a/prophetstor/api/datahub/keycodes/keycodes.proto\x1a\x17google/rpc/status.proto\"$\n\x11\x41\x64\x64KeycodeRequest\x12\x0f\n\x07keycode\x18\x01 \x01(\t\"t\n\x12\x41\x64\x64KeycodeResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12:\n\x07keycode\x18\x02 \x01(\x0b\x32).prophetstor.api.datahub.keycodes.Keycode\"\'\n\x13ListKeycodesRequest\x12\x10\n\x08keycodes\x18\x01 \x03(\t\"\xb3\x01\n\x14ListKeycodesResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12;\n\x08keycodes\x18\x02 \x03(\x0b\x32).prophetstor.api.datahub.keycodes.Keycode\x12:\n\x07summary\x18\x03 \x01(\x0b\x32).prophetstor.api.datahub.keycodes.Keycode\"\'\n\x14\x44\x65leteKeycodeRequest\x12\x0f\n\x07keycode\x18\x01 \x01(\t\"T\n GenerateRegistrationDataResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\t\"/\n\x1f\x41\x63tivateRegistrationDataRequest\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\tB&Z$prophetstor.com/api/datahub/keycodesb\x06proto3'
  ,
  dependencies=[prophetstor_dot_api_dot_datahub_dot_keycodes_dot_keycodes__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_ADDKEYCODEREQUEST = _descriptor.Descriptor(
  name='AddKeycodeRequest',
  full_name='prophetstor.api.datahub.keycodes.AddKeycodeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='keycode', full_name='prophetstor.api.datahub.keycodes.AddKeycodeRequest.keycode', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=159,
  serialized_end=195,
)


_ADDKEYCODERESPONSE = _descriptor.Descriptor(
  name='AddKeycodeResponse',
  full_name='prophetstor.api.datahub.keycodes.AddKeycodeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.api.datahub.keycodes.AddKeycodeResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='keycode', full_name='prophetstor.api.datahub.keycodes.AddKeycodeResponse.keycode', index=1,
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
  serialized_start=197,
  serialized_end=313,
)


_LISTKEYCODESREQUEST = _descriptor.Descriptor(
  name='ListKeycodesRequest',
  full_name='prophetstor.api.datahub.keycodes.ListKeycodesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='keycodes', full_name='prophetstor.api.datahub.keycodes.ListKeycodesRequest.keycodes', index=0,
      number=1, type=9, cpp_type=9, label=3,
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
  serialized_start=315,
  serialized_end=354,
)


_LISTKEYCODESRESPONSE = _descriptor.Descriptor(
  name='ListKeycodesResponse',
  full_name='prophetstor.api.datahub.keycodes.ListKeycodesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.api.datahub.keycodes.ListKeycodesResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='keycodes', full_name='prophetstor.api.datahub.keycodes.ListKeycodesResponse.keycodes', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='summary', full_name='prophetstor.api.datahub.keycodes.ListKeycodesResponse.summary', index=2,
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
  serialized_start=357,
  serialized_end=536,
)


_DELETEKEYCODEREQUEST = _descriptor.Descriptor(
  name='DeleteKeycodeRequest',
  full_name='prophetstor.api.datahub.keycodes.DeleteKeycodeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='keycode', full_name='prophetstor.api.datahub.keycodes.DeleteKeycodeRequest.keycode', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=538,
  serialized_end=577,
)


_GENERATEREGISTRATIONDATARESPONSE = _descriptor.Descriptor(
  name='GenerateRegistrationDataResponse',
  full_name='prophetstor.api.datahub.keycodes.GenerateRegistrationDataResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.api.datahub.keycodes.GenerateRegistrationDataResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data', full_name='prophetstor.api.datahub.keycodes.GenerateRegistrationDataResponse.data', index=1,
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
  serialized_start=579,
  serialized_end=663,
)


_ACTIVATEREGISTRATIONDATAREQUEST = _descriptor.Descriptor(
  name='ActivateRegistrationDataRequest',
  full_name='prophetstor.api.datahub.keycodes.ActivateRegistrationDataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='prophetstor.api.datahub.keycodes.ActivateRegistrationDataRequest.data', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=665,
  serialized_end=712,
)

_ADDKEYCODERESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_ADDKEYCODERESPONSE.fields_by_name['keycode'].message_type = prophetstor_dot_api_dot_datahub_dot_keycodes_dot_keycodes__pb2._KEYCODE
_LISTKEYCODESRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTKEYCODESRESPONSE.fields_by_name['keycodes'].message_type = prophetstor_dot_api_dot_datahub_dot_keycodes_dot_keycodes__pb2._KEYCODE
_LISTKEYCODESRESPONSE.fields_by_name['summary'].message_type = prophetstor_dot_api_dot_datahub_dot_keycodes_dot_keycodes__pb2._KEYCODE
_GENERATEREGISTRATIONDATARESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
DESCRIPTOR.message_types_by_name['AddKeycodeRequest'] = _ADDKEYCODEREQUEST
DESCRIPTOR.message_types_by_name['AddKeycodeResponse'] = _ADDKEYCODERESPONSE
DESCRIPTOR.message_types_by_name['ListKeycodesRequest'] = _LISTKEYCODESREQUEST
DESCRIPTOR.message_types_by_name['ListKeycodesResponse'] = _LISTKEYCODESRESPONSE
DESCRIPTOR.message_types_by_name['DeleteKeycodeRequest'] = _DELETEKEYCODEREQUEST
DESCRIPTOR.message_types_by_name['GenerateRegistrationDataResponse'] = _GENERATEREGISTRATIONDATARESPONSE
DESCRIPTOR.message_types_by_name['ActivateRegistrationDataRequest'] = _ACTIVATEREGISTRATIONDATAREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AddKeycodeRequest = _reflection.GeneratedProtocolMessageType('AddKeycodeRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDKEYCODEREQUEST,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.AddKeycodeRequest)
  })
_sym_db.RegisterMessage(AddKeycodeRequest)

AddKeycodeResponse = _reflection.GeneratedProtocolMessageType('AddKeycodeResponse', (_message.Message,), {
  'DESCRIPTOR' : _ADDKEYCODERESPONSE,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.AddKeycodeResponse)
  })
_sym_db.RegisterMessage(AddKeycodeResponse)

ListKeycodesRequest = _reflection.GeneratedProtocolMessageType('ListKeycodesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTKEYCODESREQUEST,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.ListKeycodesRequest)
  })
_sym_db.RegisterMessage(ListKeycodesRequest)

ListKeycodesResponse = _reflection.GeneratedProtocolMessageType('ListKeycodesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTKEYCODESRESPONSE,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.ListKeycodesResponse)
  })
_sym_db.RegisterMessage(ListKeycodesResponse)

DeleteKeycodeRequest = _reflection.GeneratedProtocolMessageType('DeleteKeycodeRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEKEYCODEREQUEST,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.DeleteKeycodeRequest)
  })
_sym_db.RegisterMessage(DeleteKeycodeRequest)

GenerateRegistrationDataResponse = _reflection.GeneratedProtocolMessageType('GenerateRegistrationDataResponse', (_message.Message,), {
  'DESCRIPTOR' : _GENERATEREGISTRATIONDATARESPONSE,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.GenerateRegistrationDataResponse)
  })
_sym_db.RegisterMessage(GenerateRegistrationDataResponse)

ActivateRegistrationDataRequest = _reflection.GeneratedProtocolMessageType('ActivateRegistrationDataRequest', (_message.Message,), {
  'DESCRIPTOR' : _ACTIVATEREGISTRATIONDATAREQUEST,
  '__module__' : 'prophetstor.api.datahub.keycodes.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.keycodes.ActivateRegistrationDataRequest)
  })
_sym_db.RegisterMessage(ActivateRegistrationDataRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
