// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// This file has messages related to keycode managements
//
'use strict';
var grpc = require('grpc');
var datahub_keycodes_services_pb = require('../../datahub/keycodes/services_pb.js');
var datahub_keycodes_keycodes_pb = require('../../datahub/keycodes/keycodes_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var google_rpc_status_pb = require('../../google/rpc/status_pb.js');

function serialize_containersai_datahub_keycodes_ActivateRegistrationDataRequest(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.ActivateRegistrationDataRequest)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.ActivateRegistrationDataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_ActivateRegistrationDataRequest(buffer_arg) {
  return datahub_keycodes_services_pb.ActivateRegistrationDataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_keycodes_AddKeycodeRequest(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.AddKeycodeRequest)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.AddKeycodeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_AddKeycodeRequest(buffer_arg) {
  return datahub_keycodes_services_pb.AddKeycodeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_keycodes_AddKeycodeResponse(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.AddKeycodeResponse)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.AddKeycodeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_AddKeycodeResponse(buffer_arg) {
  return datahub_keycodes_services_pb.AddKeycodeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_keycodes_DeleteKeycodeRequest(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.DeleteKeycodeRequest)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.DeleteKeycodeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_DeleteKeycodeRequest(buffer_arg) {
  return datahub_keycodes_services_pb.DeleteKeycodeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_keycodes_GenerateRegistrationDataResponse(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.GenerateRegistrationDataResponse)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.GenerateRegistrationDataResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_GenerateRegistrationDataResponse(buffer_arg) {
  return datahub_keycodes_services_pb.GenerateRegistrationDataResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_keycodes_ListKeycodesRequest(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.ListKeycodesRequest)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.ListKeycodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_ListKeycodesRequest(buffer_arg) {
  return datahub_keycodes_services_pb.ListKeycodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_keycodes_ListKeycodesResponse(arg) {
  if (!(arg instanceof datahub_keycodes_services_pb.ListKeycodesResponse)) {
    throw new Error('Expected argument of type containersai.datahub.keycodes.ListKeycodesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_keycodes_ListKeycodesResponse(buffer_arg) {
  return datahub_keycodes_services_pb.ListKeycodesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_rpc_Status(arg) {
  if (!(arg instanceof google_rpc_status_pb.Status)) {
    throw new Error('Expected argument of type google.rpc.Status');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_rpc_Status(buffer_arg) {
  return google_rpc_status_pb.Status.deserializeBinary(new Uint8Array(buffer_arg));
}


// *
// Service for providing operation of keycodes
var KeycodesServiceService = exports.KeycodesServiceService = {
  // Used to add a keycode
addKeycode: {
    path: '/containersai.datahub.keycodes.KeycodesService/AddKeycode',
    requestStream: false,
    responseStream: false,
    requestType: datahub_keycodes_services_pb.AddKeycodeRequest,
    responseType: datahub_keycodes_services_pb.AddKeycodeResponse,
    requestSerialize: serialize_containersai_datahub_keycodes_AddKeycodeRequest,
    requestDeserialize: deserialize_containersai_datahub_keycodes_AddKeycodeRequest,
    responseSerialize: serialize_containersai_datahub_keycodes_AddKeycodeResponse,
    responseDeserialize: deserialize_containersai_datahub_keycodes_AddKeycodeResponse,
  },
  // Used to retrieve keycodes detailed information
listKeycodes: {
    path: '/containersai.datahub.keycodes.KeycodesService/ListKeycodes',
    requestStream: false,
    responseStream: false,
    requestType: datahub_keycodes_services_pb.ListKeycodesRequest,
    responseType: datahub_keycodes_services_pb.ListKeycodesResponse,
    requestSerialize: serialize_containersai_datahub_keycodes_ListKeycodesRequest,
    requestDeserialize: deserialize_containersai_datahub_keycodes_ListKeycodesRequest,
    responseSerialize: serialize_containersai_datahub_keycodes_ListKeycodesResponse,
    responseDeserialize: deserialize_containersai_datahub_keycodes_ListKeycodesResponse,
  },
  // Used to delete a keycode
deleteKeycode: {
    path: '/containersai.datahub.keycodes.KeycodesService/DeleteKeycode',
    requestStream: false,
    responseStream: false,
    requestType: datahub_keycodes_services_pb.DeleteKeycodeRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_keycodes_DeleteKeycodeRequest,
    requestDeserialize: deserialize_containersai_datahub_keycodes_DeleteKeycodeRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to generate license registration data
generateRegistrationData: {
    path: '/containersai.datahub.keycodes.KeycodesService/GenerateRegistrationData',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: datahub_keycodes_services_pb.GenerateRegistrationDataResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_containersai_datahub_keycodes_GenerateRegistrationDataResponse,
    responseDeserialize: deserialize_containersai_datahub_keycodes_GenerateRegistrationDataResponse,
  },
  // Used to activate license signature data
activateRegistrationData: {
    path: '/containersai.datahub.keycodes.KeycodesService/ActivateRegistrationData',
    requestStream: false,
    responseStream: false,
    requestType: datahub_keycodes_services_pb.ActivateRegistrationDataRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_keycodes_ActivateRegistrationDataRequest,
    requestDeserialize: deserialize_containersai_datahub_keycodes_ActivateRegistrationDataRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
};

exports.KeycodesServiceClient = grpc.makeGenericClientConstructor(KeycodesServiceService);
