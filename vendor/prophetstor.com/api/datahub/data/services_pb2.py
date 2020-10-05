# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prophetstor/api/datahub/data/services.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prophetstor.api.datahub.data import data_pb2 as prophetstor_dot_api_dot_datahub_dot_data_dot_data__pb2
from prophetstor.api.datahub.data import types_pb2 as prophetstor_dot_api_dot_datahub_dot_data_dot_types__pb2
from prophetstor.api.datahub.schemas import types_pb2 as prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='prophetstor/api/datahub/data/services.proto',
  package='prophetstor.api.datahub.data',
  syntax='proto3',
  serialized_options=b'Z prophetstor.com/api/datahub/data',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n+prophetstor/api/datahub/data/services.proto\x12\x1cprophetstor.api.datahub.data\x1a\'prophetstor/api/datahub/data/data.proto\x1a(prophetstor/api/datahub/data/types.proto\x1a+prophetstor/api/datahub/schemas/types.proto\x1a\x17google/rpc/status.proto\"\x91\x01\n\x10WriteDataRequest\x12@\n\x0bschema_meta\x18\x01 \x01(\x0b\x32+.prophetstor.api.datahub.schemas.SchemaMeta\x12;\n\nwrite_data\x18\x02 \x03(\x0b\x32\'.prophetstor.api.datahub.data.WriteData\"\x8e\x01\n\x0fReadDataRequest\x12@\n\x0bschema_meta\x18\x01 \x01(\x0b\x32+.prophetstor.api.datahub.schemas.SchemaMeta\x12\x39\n\tread_data\x18\x02 \x03(\x0b\x32&.prophetstor.api.datahub.data.ReadData\"h\n\x10ReadDataResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x30\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\".prophetstor.api.datahub.data.Data\"\x94\x01\n\x11\x44\x65leteDataRequest\x12@\n\x0bschema_meta\x18\x01 \x01(\x0b\x32+.prophetstor.api.datahub.schemas.SchemaMeta\x12=\n\x0b\x64\x65lete_data\x18\x02 \x03(\x0b\x32(.prophetstor.api.datahub.data.DeleteData\"\x91\x01\n\x10WriteMetaRequest\x12@\n\x0bschema_meta\x18\x01 \x01(\x0b\x32+.prophetstor.api.datahub.schemas.SchemaMeta\x12;\n\nwrite_meta\x18\x02 \x03(\x0b\x32\'.prophetstor.api.datahub.data.WriteMetaB\"Z prophetstor.com/api/datahub/datab\x06proto3'
  ,
  dependencies=[prophetstor_dot_api_dot_datahub_dot_data_dot_data__pb2.DESCRIPTOR,prophetstor_dot_api_dot_datahub_dot_data_dot_types__pb2.DESCRIPTOR,prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_WRITEDATAREQUEST = _descriptor.Descriptor(
  name='WriteDataRequest',
  full_name='prophetstor.api.datahub.data.WriteDataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.api.datahub.data.WriteDataRequest.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_data', full_name='prophetstor.api.datahub.data.WriteDataRequest.write_data', index=1,
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
  serialized_start=231,
  serialized_end=376,
)


_READDATAREQUEST = _descriptor.Descriptor(
  name='ReadDataRequest',
  full_name='prophetstor.api.datahub.data.ReadDataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.api.datahub.data.ReadDataRequest.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='read_data', full_name='prophetstor.api.datahub.data.ReadDataRequest.read_data', index=1,
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
  serialized_start=379,
  serialized_end=521,
)


_READDATARESPONSE = _descriptor.Descriptor(
  name='ReadDataResponse',
  full_name='prophetstor.api.datahub.data.ReadDataResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.api.datahub.data.ReadDataResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data', full_name='prophetstor.api.datahub.data.ReadDataResponse.data', index=1,
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
  serialized_start=523,
  serialized_end=627,
)


_DELETEDATAREQUEST = _descriptor.Descriptor(
  name='DeleteDataRequest',
  full_name='prophetstor.api.datahub.data.DeleteDataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.api.datahub.data.DeleteDataRequest.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='delete_data', full_name='prophetstor.api.datahub.data.DeleteDataRequest.delete_data', index=1,
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
  serialized_start=630,
  serialized_end=778,
)


_WRITEMETAREQUEST = _descriptor.Descriptor(
  name='WriteMetaRequest',
  full_name='prophetstor.api.datahub.data.WriteMetaRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='prophetstor.api.datahub.data.WriteMetaRequest.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_meta', full_name='prophetstor.api.datahub.data.WriteMetaRequest.write_meta', index=1,
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
  serialized_start=781,
  serialized_end=926,
)

_WRITEDATAREQUEST.fields_by_name['schema_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_WRITEDATAREQUEST.fields_by_name['write_data'].message_type = prophetstor_dot_api_dot_datahub_dot_data_dot_data__pb2._WRITEDATA
_READDATAREQUEST.fields_by_name['schema_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_READDATAREQUEST.fields_by_name['read_data'].message_type = prophetstor_dot_api_dot_datahub_dot_data_dot_data__pb2._READDATA
_READDATARESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_READDATARESPONSE.fields_by_name['data'].message_type = prophetstor_dot_api_dot_datahub_dot_data_dot_types__pb2._DATA
_DELETEDATAREQUEST.fields_by_name['schema_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_DELETEDATAREQUEST.fields_by_name['delete_data'].message_type = prophetstor_dot_api_dot_datahub_dot_data_dot_data__pb2._DELETEDATA
_WRITEMETAREQUEST.fields_by_name['schema_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_WRITEMETAREQUEST.fields_by_name['write_meta'].message_type = prophetstor_dot_api_dot_datahub_dot_data_dot_data__pb2._WRITEMETA
DESCRIPTOR.message_types_by_name['WriteDataRequest'] = _WRITEDATAREQUEST
DESCRIPTOR.message_types_by_name['ReadDataRequest'] = _READDATAREQUEST
DESCRIPTOR.message_types_by_name['ReadDataResponse'] = _READDATARESPONSE
DESCRIPTOR.message_types_by_name['DeleteDataRequest'] = _DELETEDATAREQUEST
DESCRIPTOR.message_types_by_name['WriteMetaRequest'] = _WRITEMETAREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WriteDataRequest = _reflection.GeneratedProtocolMessageType('WriteDataRequest', (_message.Message,), {
  'DESCRIPTOR' : _WRITEDATAREQUEST,
  '__module__' : 'prophetstor.api.datahub.data.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.data.WriteDataRequest)
  })
_sym_db.RegisterMessage(WriteDataRequest)

ReadDataRequest = _reflection.GeneratedProtocolMessageType('ReadDataRequest', (_message.Message,), {
  'DESCRIPTOR' : _READDATAREQUEST,
  '__module__' : 'prophetstor.api.datahub.data.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.data.ReadDataRequest)
  })
_sym_db.RegisterMessage(ReadDataRequest)

ReadDataResponse = _reflection.GeneratedProtocolMessageType('ReadDataResponse', (_message.Message,), {
  'DESCRIPTOR' : _READDATARESPONSE,
  '__module__' : 'prophetstor.api.datahub.data.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.data.ReadDataResponse)
  })
_sym_db.RegisterMessage(ReadDataResponse)

DeleteDataRequest = _reflection.GeneratedProtocolMessageType('DeleteDataRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDATAREQUEST,
  '__module__' : 'prophetstor.api.datahub.data.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.data.DeleteDataRequest)
  })
_sym_db.RegisterMessage(DeleteDataRequest)

WriteMetaRequest = _reflection.GeneratedProtocolMessageType('WriteMetaRequest', (_message.Message,), {
  'DESCRIPTOR' : _WRITEMETAREQUEST,
  '__module__' : 'prophetstor.api.datahub.data.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.api.datahub.data.WriteMetaRequest)
  })
_sym_db.RegisterMessage(WriteMetaRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
