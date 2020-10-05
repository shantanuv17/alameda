// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// This file has messages and services related to Federator.ai
//
'use strict';
var grpc = require('grpc');
var prophetstor_api_datahub_applications_services_pb = require('../../../prophetstor/api/datahub/applications/services_pb.js');
var prophetstor_api_datahub_data_services_pb = require('../../../prophetstor/api/datahub/data/services_pb.js');
var prophetstor_api_datahub_events_services_pb = require('../../../prophetstor/api/datahub/events/services_pb.js');
var prophetstor_api_datahub_gpu_services_pb = require('../../../prophetstor/api/datahub/gpu/services_pb.js');
var prophetstor_api_datahub_keycodes_services_pb = require('../../../prophetstor/api/datahub/keycodes/services_pb.js');
var prophetstor_api_datahub_licenses_services_pb = require('../../../prophetstor/api/datahub/licenses/services_pb.js');
var prophetstor_api_datahub_metrics_services_pb = require('../../../prophetstor/api/datahub/metrics/services_pb.js');
var prophetstor_api_datahub_plannings_services_pb = require('../../../prophetstor/api/datahub/plannings/services_pb.js');
var prophetstor_api_datahub_predictions_services_pb = require('../../../prophetstor/api/datahub/predictions/services_pb.js');
var prophetstor_api_datahub_rawdata_services_pb = require('../../../prophetstor/api/datahub/rawdata/services_pb.js');
var prophetstor_api_datahub_recommendations_services_pb = require('../../../prophetstor/api/datahub/recommendations/services_pb.js');
var prophetstor_api_datahub_resources_services_pb = require('../../../prophetstor/api/datahub/resources/services_pb.js');
var prophetstor_api_datahub_schemas_services_pb = require('../../../prophetstor/api/datahub/schemas/services_pb.js');
var prophetstor_api_datahub_scores_services_pb = require('../../../prophetstor/api/datahub/scores/services_pb.js');
var prophetstor_api_datahub_weavescope_services_pb = require('../../../prophetstor/api/datahub/weavescope/services_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var google_rpc_status_pb = require('../../../google/rpc/status_pb.js');

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

function serialize_prophetstor_api_datahub_applications_CreateApplicationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_applications_services_pb.CreateApplicationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.applications.CreateApplicationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_applications_CreateApplicationsRequest(buffer_arg) {
  return prophetstor_api_datahub_applications_services_pb.CreateApplicationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_applications_DeleteApplicationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_applications_services_pb.DeleteApplicationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.applications.DeleteApplicationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_applications_DeleteApplicationsRequest(buffer_arg) {
  return prophetstor_api_datahub_applications_services_pb.DeleteApplicationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_applications_ListApplicationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_applications_services_pb.ListApplicationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.applications.ListApplicationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_applications_ListApplicationsRequest(buffer_arg) {
  return prophetstor_api_datahub_applications_services_pb.ListApplicationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_applications_ListApplicationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_applications_services_pb.ListApplicationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.applications.ListApplicationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_applications_ListApplicationsResponse(buffer_arg) {
  return prophetstor_api_datahub_applications_services_pb.ListApplicationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_data_DeleteDataRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_data_services_pb.DeleteDataRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.data.DeleteDataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_data_DeleteDataRequest(buffer_arg) {
  return prophetstor_api_datahub_data_services_pb.DeleteDataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_data_ReadDataRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_data_services_pb.ReadDataRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.data.ReadDataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_data_ReadDataRequest(buffer_arg) {
  return prophetstor_api_datahub_data_services_pb.ReadDataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_data_ReadDataResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_data_services_pb.ReadDataResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.data.ReadDataResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_data_ReadDataResponse(buffer_arg) {
  return prophetstor_api_datahub_data_services_pb.ReadDataResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_data_WriteDataRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_data_services_pb.WriteDataRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.data.WriteDataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_data_WriteDataRequest(buffer_arg) {
  return prophetstor_api_datahub_data_services_pb.WriteDataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_data_WriteMetaRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_data_services_pb.WriteMetaRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.data.WriteMetaRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_data_WriteMetaRequest(buffer_arg) {
  return prophetstor_api_datahub_data_services_pb.WriteMetaRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_events_CreateEventsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_events_services_pb.CreateEventsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.events.CreateEventsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_events_CreateEventsRequest(buffer_arg) {
  return prophetstor_api_datahub_events_services_pb.CreateEventsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_events_ListEventsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_events_services_pb.ListEventsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.events.ListEventsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_events_ListEventsRequest(buffer_arg) {
  return prophetstor_api_datahub_events_services_pb.ListEventsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_events_ListEventsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_events_services_pb.ListEventsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.events.ListEventsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_events_ListEventsResponse(buffer_arg) {
  return prophetstor_api_datahub_events_services_pb.ListEventsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_CreateGpuPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.CreateGpuPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.CreateGpuPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_CreateGpuPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.CreateGpuPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_ListGpuMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.ListGpuMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.ListGpuMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_ListGpuMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.ListGpuMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_ListGpuMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.ListGpuMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.ListGpuMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_ListGpuMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.ListGpuMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_ListGpuPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.ListGpuPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.ListGpuPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_ListGpuPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.ListGpuPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_ListGpuPredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.ListGpuPredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.ListGpuPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_ListGpuPredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.ListGpuPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_ListGpusRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.ListGpusRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.ListGpusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_ListGpusRequest(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.ListGpusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_gpu_ListGpusResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_gpu_services_pb.ListGpusResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.gpu.ListGpusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_gpu_ListGpusResponse(buffer_arg) {
  return prophetstor_api_datahub_gpu_services_pb.ListGpusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_ActivateRegistrationDataRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.ActivateRegistrationDataRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.ActivateRegistrationDataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_ActivateRegistrationDataRequest(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.ActivateRegistrationDataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_AddKeycodeRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.AddKeycodeRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.AddKeycodeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_AddKeycodeRequest(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.AddKeycodeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_AddKeycodeResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.AddKeycodeResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.AddKeycodeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_AddKeycodeResponse(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.AddKeycodeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_DeleteKeycodeRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.DeleteKeycodeRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.DeleteKeycodeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_DeleteKeycodeRequest(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.DeleteKeycodeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_GenerateRegistrationDataResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.GenerateRegistrationDataResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.GenerateRegistrationDataResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_GenerateRegistrationDataResponse(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.GenerateRegistrationDataResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_ListKeycodesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.ListKeycodesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.ListKeycodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_ListKeycodesRequest(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.ListKeycodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_keycodes_ListKeycodesResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_keycodes_services_pb.ListKeycodesResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.keycodes.ListKeycodesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_keycodes_ListKeycodesResponse(buffer_arg) {
  return prophetstor_api_datahub_keycodes_services_pb.ListKeycodesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_licenses_GetLicenseResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_licenses_services_pb.GetLicenseResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.licenses.GetLicenseResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_licenses_GetLicenseResponse(buffer_arg) {
  return prophetstor_api_datahub_licenses_services_pb.GetLicenseResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreateApplicationMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreateApplicationMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreateApplicationMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreateApplicationMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreateApplicationMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreateClusterMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreateClusterMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreateClusterMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreateClusterMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreateClusterMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreateControllerMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreateControllerMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreateControllerMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreateControllerMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreateControllerMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreateMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreateMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreateMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreateMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreateMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreateNamespaceMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreateNamespaceMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreateNamespaceMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreateNamespaceMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreateNamespaceMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreateNodeMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreateNodeMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreateNodeMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreateNodeMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreateNodeMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_CreatePodMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.CreatePodMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.CreatePodMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_CreatePodMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.CreatePodMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListApplicationMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListApplicationMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListApplicationMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListApplicationMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListApplicationMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListApplicationMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListApplicationMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListApplicationMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListApplicationMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListApplicationMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListClusterMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListClusterMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListClusterMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListClusterMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListClusterMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListClusterMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListClusterMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListClusterMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListClusterMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListClusterMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListControllerMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListControllerMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListControllerMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListControllerMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListControllerMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListControllerMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListControllerMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListControllerMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListControllerMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListControllerMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListNamespaceMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListNamespaceMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListNamespaceMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListNamespaceMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListNamespaceMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListNamespaceMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListNodeMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListNodeMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListNodeMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListNodeMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListNodeMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListNodeMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListNodeMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListNodeMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListNodeMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListNodeMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListPodMetricsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListPodMetricsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListPodMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListPodMetricsRequest(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListPodMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_metrics_ListPodMetricsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_metrics_services_pb.ListPodMetricsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.metrics.ListPodMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_metrics_ListPodMetricsResponse(buffer_arg) {
  return prophetstor_api_datahub_metrics_services_pb.ListPodMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreateApplicationPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreateApplicationPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreateApplicationPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreateApplicationPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreateApplicationPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreateClusterPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreateClusterPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreateClusterPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreateClusterPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreateClusterPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreateControllerPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreateControllerPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreateControllerPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreateControllerPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreateControllerPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreateNamespacePlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreateNamespacePlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreateNamespacePlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreateNamespacePlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreateNamespacePlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreateNodePlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreateNodePlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreateNodePlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreateNodePlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreateNodePlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreatePlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreatePlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreatePlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreatePlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreatePlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_CreatePodPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.CreatePodPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.CreatePodPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_CreatePodPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.CreatePodPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListApplicationPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListApplicationPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListApplicationPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListApplicationPlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListApplicationPlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListApplicationPlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListClusterPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListClusterPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListClusterPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListClusterPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListClusterPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListClusterPlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListClusterPlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListClusterPlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListClusterPlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListClusterPlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListControllerPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListControllerPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListControllerPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListControllerPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListControllerPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListControllerPlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListControllerPlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListControllerPlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListControllerPlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListControllerPlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListNamespacePlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListNamespacePlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListNamespacePlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListNamespacePlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListNamespacePlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListNamespacePlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListNodePlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListNodePlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListNodePlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListNodePlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListNodePlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListNodePlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListNodePlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListNodePlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListNodePlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListNodePlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListPlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListPlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListPlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListPlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListPlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListPodPlanningsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListPodPlanningsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListPodPlanningsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListPodPlanningsRequest(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListPodPlanningsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_plannings_ListPodPlanningsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_plannings_services_pb.ListPodPlanningsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.plannings.ListPodPlanningsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_plannings_ListPodPlanningsResponse(buffer_arg) {
  return prophetstor_api_datahub_plannings_services_pb.ListPodPlanningsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreateApplicationPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreateApplicationPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreateApplicationPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreateApplicationPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreateApplicationPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreateClusterPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreateClusterPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreateClusterPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreateClusterPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreateClusterPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreateControllerPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreateControllerPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreateControllerPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreateControllerPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreateControllerPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreateNamespacePredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreateNamespacePredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreateNamespacePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreateNamespacePredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreateNamespacePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreateNodePredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreateNodePredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreateNodePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreateNodePredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreateNodePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreatePodPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreatePodPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreatePodPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreatePodPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreatePodPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_CreatePredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.CreatePredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.CreatePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_CreatePredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.CreatePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListApplicationPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListApplicationPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListApplicationPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListApplicationPredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListApplicationPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListApplicationPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListClusterPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListClusterPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListClusterPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListClusterPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListClusterPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListClusterPredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListClusterPredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListClusterPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListClusterPredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListClusterPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListControllerPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListControllerPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListControllerPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListControllerPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListControllerPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListControllerPredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListControllerPredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListControllerPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListControllerPredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListControllerPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListNamespacePredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListNamespacePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListNamespacePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListNamespacePredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListNamespacePredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListNamespacePredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListNodePredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListNodePredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListNodePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListNodePredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListNodePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListNodePredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListNodePredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListNodePredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListNodePredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListNodePredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListPodPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListPodPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListPodPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListPodPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListPodPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListPodPredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListPodPredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListPodPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListPodPredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListPodPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListPredictionsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListPredictionsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListPredictionsRequest(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_predictions_ListPredictionsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_predictions_services_pb.ListPredictionsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.predictions.ListPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_predictions_ListPredictionsResponse(buffer_arg) {
  return prophetstor_api_datahub_predictions_services_pb.ListPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_rawdata_ReadRawdataRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_rawdata_services_pb.ReadRawdataRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.rawdata.ReadRawdataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_rawdata_ReadRawdataRequest(buffer_arg) {
  return prophetstor_api_datahub_rawdata_services_pb.ReadRawdataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_rawdata_ReadRawdataResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_rawdata_services_pb.ReadRawdataResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.rawdata.ReadRawdataResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_rawdata_ReadRawdataResponse(buffer_arg) {
  return prophetstor_api_datahub_rawdata_services_pb.ReadRawdataResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_rawdata_WriteRawdataRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_rawdata_services_pb.WriteRawdataRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.rawdata.WriteRawdataRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_rawdata_WriteRawdataRequest(buffer_arg) {
  return prophetstor_api_datahub_rawdata_services_pb.WriteRawdataRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreateApplicationRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreateApplicationRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreateApplicationRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreateApplicationRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreateApplicationRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreateClusterRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreateClusterRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreateClusterRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreateClusterRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreateClusterRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreateControllerRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreateControllerRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreateControllerRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreateControllerRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreateControllerRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreateNamespaceRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreateNamespaceRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreateNamespaceRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreateNamespaceRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreateNamespaceRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreateNodeRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreateNodeRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreateNodeRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreateNodeRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreateNodeRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreatePodRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreatePodRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreatePodRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreatePodRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreatePodRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_CreateRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.CreateRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.CreateRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_CreateRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.CreateRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListApplicationRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListApplicationRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListApplicationRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListApplicationRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListClusterRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListClusterRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListClusterRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListClusterRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListClusterRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListClusterRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListControllerRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListControllerRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListControllerRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListControllerRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListControllerRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListControllerRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListNamespaceRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListNamespaceRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListNamespaceRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListNamespaceRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListNodeRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListNodeRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListNodeRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListNodeRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListNodeRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListNodeRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListPodRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListPodRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListRecommendationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListRecommendationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListRecommendationsRequest(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_recommendations_ListRecommendationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_recommendations_services_pb.ListRecommendationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.recommendations.ListRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_recommendations_ListRecommendationsResponse(buffer_arg) {
  return prophetstor_api_datahub_recommendations_services_pb.ListRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_CreateApplicationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.CreateApplicationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.CreateApplicationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_CreateApplicationsRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.CreateApplicationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_CreateClustersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.CreateClustersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.CreateClustersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_CreateClustersRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.CreateClustersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_CreateControllersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.CreateControllersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.CreateControllersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_CreateControllersRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.CreateControllersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_CreateNamespacesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.CreateNamespacesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.CreateNamespacesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_CreateNamespacesRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.CreateNamespacesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_CreateNodesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.CreateNodesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.CreateNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_CreateNodesRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.CreateNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_CreatePodsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.CreatePodsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.CreatePodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_CreatePodsRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.CreatePodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_DeleteApplicationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.DeleteApplicationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.DeleteApplicationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_DeleteApplicationsRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.DeleteApplicationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_DeleteClustersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.DeleteClustersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.DeleteClustersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_DeleteClustersRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.DeleteClustersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_DeleteControllersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.DeleteControllersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.DeleteControllersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_DeleteControllersRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.DeleteControllersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_DeleteNamespacesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.DeleteNamespacesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.DeleteNamespacesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_DeleteNamespacesRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.DeleteNamespacesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_DeleteNodesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.DeleteNodesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.DeleteNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_DeleteNodesRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.DeleteNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_DeletePodsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.DeletePodsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.DeletePodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_DeletePodsRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.DeletePodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListApplicationsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListApplicationsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListApplicationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListApplicationsRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListApplicationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListApplicationsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListApplicationsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListApplicationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListApplicationsResponse(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListApplicationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListClustersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListClustersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListClustersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListClustersRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListClustersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListClustersResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListClustersResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListClustersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListClustersResponse(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListClustersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListControllersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListControllersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListControllersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListControllersRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListControllersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListControllersResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListControllersResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListControllersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListControllersResponse(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListControllersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListNamespacesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListNamespacesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListNamespacesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListNamespacesRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListNamespacesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListNamespacesResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListNamespacesResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListNamespacesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListNamespacesResponse(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListNamespacesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListNodesRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListNodesRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListNodesRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListNodesResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListNodesResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListNodesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListNodesResponse(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListNodesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListPodsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListPodsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListPodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListPodsRequest(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListPodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_resources_ListPodsResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_resources_services_pb.ListPodsResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.resources.ListPodsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_resources_ListPodsResponse(buffer_arg) {
  return prophetstor_api_datahub_resources_services_pb.ListPodsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_schemas_CreateSchemasRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_schemas_services_pb.CreateSchemasRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.schemas.CreateSchemasRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_schemas_CreateSchemasRequest(buffer_arg) {
  return prophetstor_api_datahub_schemas_services_pb.CreateSchemasRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_schemas_DeleteSchemasRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_schemas_services_pb.DeleteSchemasRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.schemas.DeleteSchemasRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_schemas_DeleteSchemasRequest(buffer_arg) {
  return prophetstor_api_datahub_schemas_services_pb.DeleteSchemasRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_schemas_ListSchemasRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_schemas_services_pb.ListSchemasRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.schemas.ListSchemasRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_schemas_ListSchemasRequest(buffer_arg) {
  return prophetstor_api_datahub_schemas_services_pb.ListSchemasRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_schemas_ListSchemasResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_schemas_services_pb.ListSchemasResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.schemas.ListSchemasResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_schemas_ListSchemasResponse(buffer_arg) {
  return prophetstor_api_datahub_schemas_services_pb.ListSchemasResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_scores_CreateSimulatedSchedulingScoresRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_scores_services_pb.CreateSimulatedSchedulingScoresRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.scores.CreateSimulatedSchedulingScoresRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_scores_CreateSimulatedSchedulingScoresRequest(buffer_arg) {
  return prophetstor_api_datahub_scores_services_pb.CreateSimulatedSchedulingScoresRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_scores_services_pb.ListSimulatedSchedulingScoresRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresRequest(buffer_arg) {
  return prophetstor_api_datahub_scores_services_pb.ListSimulatedSchedulingScoresRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.scores.ListSimulatedSchedulingScoresResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresResponse(buffer_arg) {
  return prophetstor_api_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.weavescope.ListWeaveScopeContainersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest(buffer_arg) {
  return prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeHostsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeHostsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.weavescope.ListWeaveScopeHostsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeHostsRequest(buffer_arg) {
  return prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeHostsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_weavescope_ListWeaveScopePodsRequest(arg) {
  if (!(arg instanceof prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopePodsRequest)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.weavescope.ListWeaveScopePodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopePodsRequest(buffer_arg) {
  return prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopePodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse(arg) {
  if (!(arg instanceof prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse)) {
    throw new Error('Expected argument of type prophetstor.api.datahub.weavescope.WeaveScopeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse(buffer_arg) {
  return prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Service for providing data stored in the backend
var DatahubServiceService = exports.DatahubServiceService = {
  createApps: {
    path: '/prophetstor.api.datahub.DatahubService/CreateApps',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_applications_services_pb.CreateApplicationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_applications_CreateApplicationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_applications_CreateApplicationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  listApps: {
    path: '/prophetstor.api.datahub.DatahubService/ListApps',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_applications_services_pb.ListApplicationsRequest,
    responseType: prophetstor_api_datahub_applications_services_pb.ListApplicationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_applications_ListApplicationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_applications_ListApplicationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_applications_ListApplicationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_applications_ListApplicationsResponse,
  },
  deleteApps: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteApps',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_applications_services_pb.DeleteApplicationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_applications_DeleteApplicationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_applications_DeleteApplicationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to read data based on alameda specific schemas
readData: {
    path: '/prophetstor.api.datahub.DatahubService/ReadData',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_data_services_pb.ReadDataRequest,
    responseType: prophetstor_api_datahub_data_services_pb.ReadDataResponse,
    requestSerialize: serialize_prophetstor_api_datahub_data_ReadDataRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_data_ReadDataRequest,
    responseSerialize: serialize_prophetstor_api_datahub_data_ReadDataResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_data_ReadDataResponse,
  },
  // Used to write data based on alameda specific schemas
writeData: {
    path: '/prophetstor.api.datahub.DatahubService/WriteData',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_data_services_pb.WriteDataRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_data_WriteDataRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_data_WriteDataRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to delete data based on alameda specific schemas
deleteData: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteData',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_data_services_pb.DeleteDataRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_data_DeleteDataRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_data_DeleteDataRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to write metadata based on alameda specific schemas
writeMeta: {
    path: '/prophetstor.api.datahub.DatahubService/WriteMeta',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_data_services_pb.WriteMetaRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_data_WriteMetaRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_data_WriteMetaRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create alameda specific events
createEvents: {
    path: '/prophetstor.api.datahub.DatahubService/CreateEvents',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_events_services_pb.CreateEventsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_events_CreateEventsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_events_CreateEventsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to list alameda specific events
listEvents: {
    path: '/prophetstor.api.datahub.DatahubService/ListEvents',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_events_services_pb.ListEventsRequest,
    responseType: prophetstor_api_datahub_events_services_pb.ListEventsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_events_ListEventsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_events_ListEventsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_events_ListEventsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_events_ListEventsResponse,
  },
  // Used to create GPU predictions
createGpuPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreateGpuPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_gpu_services_pb.CreateGpuPredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_gpu_CreateGpuPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_gpu_CreateGpuPredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to list GPU need to be predicted
listGpus: {
    path: '/prophetstor.api.datahub.DatahubService/ListGpus',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_gpu_services_pb.ListGpusRequest,
    responseType: prophetstor_api_datahub_gpu_services_pb.ListGpusResponse,
    requestSerialize: serialize_prophetstor_api_datahub_gpu_ListGpusRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_gpu_ListGpusRequest,
    responseSerialize: serialize_prophetstor_api_datahub_gpu_ListGpusResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_gpu_ListGpusResponse,
  },
  // Used to list GPU metrics data
listGpuMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListGpuMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_gpu_services_pb.ListGpuMetricsRequest,
    responseType: prophetstor_api_datahub_gpu_services_pb.ListGpuMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_gpu_ListGpuMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_gpu_ListGpuMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_gpu_ListGpuMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_gpu_ListGpuMetricsResponse,
  },
  // Used to list GPU predictions
listGpuPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListGpuPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_gpu_services_pb.ListGpuPredictionsRequest,
    responseType: prophetstor_api_datahub_gpu_services_pb.ListGpuPredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_gpu_ListGpuPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_gpu_ListGpuPredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_gpu_ListGpuPredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_gpu_ListGpuPredictionsResponse,
  },
  // Used to add a keycode
addKeycode: {
    path: '/prophetstor.api.datahub.DatahubService/AddKeycode',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_keycodes_services_pb.AddKeycodeRequest,
    responseType: prophetstor_api_datahub_keycodes_services_pb.AddKeycodeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_keycodes_AddKeycodeRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_keycodes_AddKeycodeRequest,
    responseSerialize: serialize_prophetstor_api_datahub_keycodes_AddKeycodeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_keycodes_AddKeycodeResponse,
  },
  // Used to retrieve keycodes detailed information
listKeycodes: {
    path: '/prophetstor.api.datahub.DatahubService/ListKeycodes',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_keycodes_services_pb.ListKeycodesRequest,
    responseType: prophetstor_api_datahub_keycodes_services_pb.ListKeycodesResponse,
    requestSerialize: serialize_prophetstor_api_datahub_keycodes_ListKeycodesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_keycodes_ListKeycodesRequest,
    responseSerialize: serialize_prophetstor_api_datahub_keycodes_ListKeycodesResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_keycodes_ListKeycodesResponse,
  },
  // Used to delete a keycode
deleteKeycode: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteKeycode',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_keycodes_services_pb.DeleteKeycodeRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_keycodes_DeleteKeycodeRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_keycodes_DeleteKeycodeRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to generate license registration data
generateRegistrationData: {
    path: '/prophetstor.api.datahub.DatahubService/GenerateRegistrationData',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: prophetstor_api_datahub_keycodes_services_pb.GenerateRegistrationDataResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_prophetstor_api_datahub_keycodes_GenerateRegistrationDataResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_keycodes_GenerateRegistrationDataResponse,
  },
  // Used to activate license signature data
activateRegistrationData: {
    path: '/prophetstor.api.datahub.DatahubService/ActivateRegistrationData',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_keycodes_services_pb.ActivateRegistrationDataRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_keycodes_ActivateRegistrationDataRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_keycodes_ActivateRegistrationDataRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to get datahub license information
getLicense: {
    path: '/prophetstor.api.datahub.DatahubService/GetLicense',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: prophetstor_api_datahub_licenses_services_pb.GetLicenseResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_prophetstor_api_datahub_licenses_GetLicenseResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_licenses_GetLicenseResponse,
  },
  createMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreateMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreateMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreateMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreateMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createPodMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePodMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreatePodMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreatePodMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreatePodMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createControllerMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreateControllerMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreateControllerMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreateControllerMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreateControllerMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createApplicationMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreateApplicationMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreateApplicationMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreateApplicationMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreateApplicationMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createNamespaceMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNamespaceMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreateNamespaceMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreateNamespaceMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreateNamespaceMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createNodeMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNodeMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreateNodeMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreateNodeMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreateNodeMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createClusterMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/CreateClusterMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.CreateClusterMetricsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_CreateClusterMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_CreateClusterMetricsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  listMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListMetricsResponse,
  },
  // Used to list pod metric data
listPodMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListPodMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListPodMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListPodMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListPodMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListPodMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListPodMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListPodMetricsResponse,
  },
  // Used to list controller metric data
listControllerMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListControllerMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListControllerMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListControllerMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListControllerMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListControllerMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListControllerMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListControllerMetricsResponse,
  },
  // Used to list alameda scaler metric data
listApplicationMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListApplicationMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListApplicationMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListApplicationMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListApplicationMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListApplicationMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListApplicationMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListApplicationMetricsResponse,
  },
  // Used to list namespace metric data
listNamespaceMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListNamespaceMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListNamespaceMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListNamespaceMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListNamespaceMetricsResponse,
  },
  // Used to list node metric data
listNodeMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListNodeMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListNodeMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListNodeMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListNodeMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListNodeMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListNodeMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListNodeMetricsResponse,
  },
  // Used to list cluster metric data
listClusterMetrics: {
    path: '/prophetstor.api.datahub.DatahubService/ListClusterMetrics',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_metrics_services_pb.ListClusterMetricsRequest,
    responseType: prophetstor_api_datahub_metrics_services_pb.ListClusterMetricsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_metrics_ListClusterMetricsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_metrics_ListClusterMetricsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_metrics_ListClusterMetricsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_metrics_ListClusterMetricsResponse,
  },
  // Used to check if datahub is still alive
ping: {
    path: '/prophetstor.api.datahub.DatahubService/Ping',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreatePlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreatePlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreatePlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createPodPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePodPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreatePodPlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreatePodPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreatePodPlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createControllerPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreateControllerPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreateControllerPlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreateControllerPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreateControllerPlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createApplicationPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreateApplicationPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreateApplicationPlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreateApplicationPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreateApplicationPlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createNamespacePlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNamespacePlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreateNamespacePlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreateNamespacePlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreateNamespacePlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createNodePlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNodePlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreateNodePlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreateNodePlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreateNodePlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createClusterPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/CreateClusterPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.CreateClusterPlanningsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_CreateClusterPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_CreateClusterPlanningsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  listPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListPlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListPlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListPlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListPlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListPlanningsResponse,
  },
  listPodPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListPodPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListPodPlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListPodPlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListPodPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListPodPlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListPodPlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListPodPlanningsResponse,
  },
  listControllerPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListControllerPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListControllerPlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListControllerPlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListControllerPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListControllerPlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListControllerPlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListControllerPlanningsResponse,
  },
  listApplicationPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListApplicationPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListApplicationPlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListApplicationPlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListApplicationPlanningsResponse,
  },
  listNamespacePlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListNamespacePlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListNamespacePlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListNamespacePlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListNamespacePlanningsResponse,
  },
  listNodePlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListNodePlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListNodePlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListNodePlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListNodePlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListNodePlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListNodePlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListNodePlanningsResponse,
  },
  listClusterPlannings: {
    path: '/prophetstor.api.datahub.DatahubService/ListClusterPlannings',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_plannings_services_pb.ListClusterPlanningsRequest,
    responseType: prophetstor_api_datahub_plannings_services_pb.ListClusterPlanningsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_plannings_ListClusterPlanningsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_plannings_ListClusterPlanningsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_plannings_ListClusterPlanningsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_plannings_ListClusterPlanningsResponse,
  },
  createPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreatePredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreatePredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreatePredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create predictions of pods
createPodPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePodPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreatePodPredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreatePodPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreatePodPredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create predictions of controllers
createControllerPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreateControllerPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreateControllerPredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreateControllerPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreateControllerPredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create predictions of alameda scalers
createApplicationPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreateApplicationPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreateApplicationPredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreateApplicationPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreateApplicationPredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create predictions of namespaces
createNamespacePredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNamespacePredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreateNamespacePredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreateNamespacePredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreateNamespacePredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create predictions of nodes
createNodePredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNodePredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreateNodePredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreateNodePredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreateNodePredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create predictions of clusters
createClusterPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/CreateClusterPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.CreateClusterPredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_CreateClusterPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_CreateClusterPredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  listPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListPredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListPredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListPredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListPredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListPredictionsResponse,
  },
  // Used to list pod predictions
listPodPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListPodPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListPodPredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListPodPredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListPodPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListPodPredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListPodPredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListPodPredictionsResponse,
  },
  // Used to list controller predictions
listControllerPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListControllerPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListControllerPredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListControllerPredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListControllerPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListControllerPredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListControllerPredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListControllerPredictionsResponse,
  },
  // Used to list alameda scaler predictions
listApplicationPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListApplicationPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListApplicationPredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListApplicationPredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListApplicationPredictionsResponse,
  },
  // Used to list namespace predictions
listNamespacePredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListNamespacePredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListNamespacePredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListNamespacePredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListNamespacePredictionsResponse,
  },
  // Used to list node predictions
listNodePredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListNodePredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListNodePredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListNodePredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListNodePredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListNodePredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListNodePredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListNodePredictionsResponse,
  },
  // Used to list cluster predictions
listClusterPredictions: {
    path: '/prophetstor.api.datahub.DatahubService/ListClusterPredictions',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_predictions_services_pb.ListClusterPredictionsRequest,
    responseType: prophetstor_api_datahub_predictions_services_pb.ListClusterPredictionsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_predictions_ListClusterPredictionsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_predictions_ListClusterPredictionsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_predictions_ListClusterPredictionsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_predictions_ListClusterPredictionsResponse,
  },
  // Rawdata --------------------------------------------------
readRawdata: {
    path: '/prophetstor.api.datahub.DatahubService/ReadRawdata',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_rawdata_services_pb.ReadRawdataRequest,
    responseType: prophetstor_api_datahub_rawdata_services_pb.ReadRawdataResponse,
    requestSerialize: serialize_prophetstor_api_datahub_rawdata_ReadRawdataRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_rawdata_ReadRawdataRequest,
    responseSerialize: serialize_prophetstor_api_datahub_rawdata_ReadRawdataResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_rawdata_ReadRawdataResponse,
  },
  writeRawdata: {
    path: '/prophetstor.api.datahub.DatahubService/WriteRawdata',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_rawdata_services_pb.WriteRawdataRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_rawdata_WriteRawdataRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_rawdata_WriteRawdataRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  createRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreateRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreateRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreateRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreateRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create recommendations of pods
createPodRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePodRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreatePodRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreatePodRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreatePodRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create recommendations of controllers
createControllerRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreateControllerRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreateControllerRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreateControllerRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreateControllerRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create recommendations of alameda scalers
createApplicationRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreateApplicationRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreateApplicationRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreateApplicationRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreateApplicationRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create recommendations of namespaces
createNamespaceRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNamespaceRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreateNamespaceRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreateNamespaceRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreateNamespaceRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create recommendations of nodes
createNodeRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNodeRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreateNodeRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreateNodeRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreateNodeRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to create recommendations of clusters
createClusterRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/CreateClusterRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.CreateClusterRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_CreateClusterRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_CreateClusterRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  listRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListRecommendationsResponse,
  },
  // Used to list pod recommenations
listPodRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListPodRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsResponse,
  },
  // Used to list available pod recommenations
listAvailablePodRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListAvailablePodRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListPodRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListPodRecommendationsResponse,
  },
  // Used to list controller recommenations
listControllerRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListControllerRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListControllerRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListControllerRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListControllerRecommendationsResponse,
  },
  // Used to list alameda scaler recommenations
listApplicationRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListApplicationRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListApplicationRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListApplicationRecommendationsResponse,
  },
  // Used to list namespace recommenations
listNamespaceRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListNamespaceRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListNamespaceRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListNamespaceRecommendationsResponse,
  },
  // Used to list node recommenations
listNodeRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListNodeRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListNodeRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListNodeRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListNodeRecommendationsResponse,
  },
  // Used to list cluster recommenations
listClusterRecommendations: {
    path: '/prophetstor.api.datahub.DatahubService/ListClusterRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_recommendations_services_pb.ListClusterRecommendationsRequest,
    responseType: prophetstor_api_datahub_recommendations_services_pb.ListClusterRecommendationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_recommendations_ListClusterRecommendationsResponse,
  },
  // Used to add pods that need to be predicted
createPods: {
    path: '/prophetstor.api.datahub.DatahubService/CreatePods',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.CreatePodsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_CreatePodsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_CreatePodsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to add controllers that need to be predicted
createControllers: {
    path: '/prophetstor.api.datahub.DatahubService/CreateControllers',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.CreateControllersRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_CreateControllersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_CreateControllersRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to add alameda scalers that need to be predicted
createApplications: {
    path: '/prophetstor.api.datahub.DatahubService/CreateApplications',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.CreateApplicationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_CreateApplicationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_CreateApplicationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to add namespaces that need to be predicted
createNamespaces: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNamespaces',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.CreateNamespacesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_CreateNamespacesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_CreateNamespacesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to add nodes that need to be predicted
createNodes: {
    path: '/prophetstor.api.datahub.DatahubService/CreateNodes',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.CreateNodesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_CreateNodesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_CreateNodesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to add clusters that need to be predicted
createClusters: {
    path: '/prophetstor.api.datahub.DatahubService/CreateClusters',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.CreateClustersRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_CreateClustersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_CreateClustersRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to list pods need to be predicted
listPods: {
    path: '/prophetstor.api.datahub.DatahubService/ListPods',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.ListPodsRequest,
    responseType: prophetstor_api_datahub_resources_services_pb.ListPodsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_resources_ListPodsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_ListPodsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_resources_ListPodsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_resources_ListPodsResponse,
  },
  // Used to list contollers need to be predicted
listControllers: {
    path: '/prophetstor.api.datahub.DatahubService/ListControllers',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.ListControllersRequest,
    responseType: prophetstor_api_datahub_resources_services_pb.ListControllersResponse,
    requestSerialize: serialize_prophetstor_api_datahub_resources_ListControllersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_ListControllersRequest,
    responseSerialize: serialize_prophetstor_api_datahub_resources_ListControllersResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_resources_ListControllersResponse,
  },
  // Used to list alameda scalers need to be predicted
listApplications: {
    path: '/prophetstor.api.datahub.DatahubService/ListApplications',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.ListApplicationsRequest,
    responseType: prophetstor_api_datahub_resources_services_pb.ListApplicationsResponse,
    requestSerialize: serialize_prophetstor_api_datahub_resources_ListApplicationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_ListApplicationsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_resources_ListApplicationsResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_resources_ListApplicationsResponse,
  },
  // Used to list namespaces need to be predicted
listNamespaces: {
    path: '/prophetstor.api.datahub.DatahubService/ListNamespaces',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.ListNamespacesRequest,
    responseType: prophetstor_api_datahub_resources_services_pb.ListNamespacesResponse,
    requestSerialize: serialize_prophetstor_api_datahub_resources_ListNamespacesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_ListNamespacesRequest,
    responseSerialize: serialize_prophetstor_api_datahub_resources_ListNamespacesResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_resources_ListNamespacesResponse,
  },
  // Used to list nodes' information
listNodes: {
    path: '/prophetstor.api.datahub.DatahubService/ListNodes',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.ListNodesRequest,
    responseType: prophetstor_api_datahub_resources_services_pb.ListNodesResponse,
    requestSerialize: serialize_prophetstor_api_datahub_resources_ListNodesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_ListNodesRequest,
    responseSerialize: serialize_prophetstor_api_datahub_resources_ListNodesResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_resources_ListNodesResponse,
  },
  // Used to list clusters' information
listClusters: {
    path: '/prophetstor.api.datahub.DatahubService/ListClusters',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.ListClustersRequest,
    responseType: prophetstor_api_datahub_resources_services_pb.ListClustersResponse,
    requestSerialize: serialize_prophetstor_api_datahub_resources_ListClustersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_ListClustersRequest,
    responseSerialize: serialize_prophetstor_api_datahub_resources_ListClustersResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_resources_ListClustersResponse,
  },
  // Used to stop generating predictions for the pods
deletePods: {
    path: '/prophetstor.api.datahub.DatahubService/DeletePods',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.DeletePodsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_DeletePodsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_DeletePodsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to stop generating predictions for the controllers
deleteControllers: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteControllers',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.DeleteControllersRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_DeleteControllersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_DeleteControllersRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to stop generating predictions for the applications
deleteApplications: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteApplications',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.DeleteApplicationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_DeleteApplicationsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_DeleteApplicationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to stop generating predictions for the namespaces
deleteNamespaces: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteNamespaces',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.DeleteNamespacesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_DeleteNamespacesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_DeleteNamespacesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to stop generating predictions for the nodes
deleteNodes: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteNodes',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.DeleteNodesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_DeleteNodesRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_DeleteNodesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to stop generating predictions for the clusters
deleteClusters: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteClusters',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_resources_services_pb.DeleteClustersRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_resources_DeleteClustersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_resources_DeleteClustersRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Schemas --------------------------------------------------
createSchemas: {
    path: '/prophetstor.api.datahub.DatahubService/CreateSchemas',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_schemas_services_pb.CreateSchemasRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_schemas_CreateSchemasRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_schemas_CreateSchemasRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  listSchemas: {
    path: '/prophetstor.api.datahub.DatahubService/ListSchemas',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_schemas_services_pb.ListSchemasRequest,
    responseType: prophetstor_api_datahub_schemas_services_pb.ListSchemasResponse,
    requestSerialize: serialize_prophetstor_api_datahub_schemas_ListSchemasRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_schemas_ListSchemasRequest,
    responseSerialize: serialize_prophetstor_api_datahub_schemas_ListSchemasResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_schemas_ListSchemasResponse,
  },
  deleteSchemas: {
    path: '/prophetstor.api.datahub.DatahubService/DeleteSchemas',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_schemas_services_pb.DeleteSchemasRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_schemas_DeleteSchemasRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_schemas_DeleteSchemasRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Scores --------------------------------------------------
createSimulatedSchedulingScores: {
    path: '/prophetstor.api.datahub.DatahubService/CreateSimulatedSchedulingScores',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_scores_services_pb.CreateSimulatedSchedulingScoresRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_prophetstor_api_datahub_scores_CreateSimulatedSchedulingScoresRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_scores_CreateSimulatedSchedulingScoresRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // Used to list system scores
listSimulatedSchedulingScores: {
    path: '/prophetstor.api.datahub.DatahubService/ListSimulatedSchedulingScores',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_scores_services_pb.ListSimulatedSchedulingScoresRequest,
    responseType: prophetstor_api_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse,
    requestSerialize: serialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresRequest,
    responseSerialize: serialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_scores_ListSimulatedSchedulingScoresResponse,
  },
  // weave scope --------------------------------------------------
listWeaveScopeHosts: {
    path: '/prophetstor.api.datahub.DatahubService/ListWeaveScopeHosts',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeHostsRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeHostsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeHostsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  getWeaveScopeHostDetails: {
    path: '/prophetstor.api.datahub.DatahubService/GetWeaveScopeHostDetails',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeHostsRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeHostsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeHostsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  listWeaveScopePods: {
    path: '/prophetstor.api.datahub.DatahubService/ListWeaveScopePods',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopePodsRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopePodsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopePodsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  getWeaveScopePodDetails: {
    path: '/prophetstor.api.datahub.DatahubService/GetWeaveScopePodDetails',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopePodsRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopePodsRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopePodsRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  listWeaveScopeContainers: {
    path: '/prophetstor.api.datahub.DatahubService/ListWeaveScopeContainers',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  listWeaveScopeContainersByHostname: {
    path: '/prophetstor.api.datahub.DatahubService/ListWeaveScopeContainersByHostname',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  listWeaveScopeContainersByImage: {
    path: '/prophetstor.api.datahub.DatahubService/ListWeaveScopeContainersByImage',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
  getWeaveScopeContainerDetails: {
    path: '/prophetstor.api.datahub.DatahubService/GetWeaveScopeContainerDetails',
    requestStream: false,
    responseStream: false,
    requestType: prophetstor_api_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
    responseType: prophetstor_api_datahub_weavescope_services_pb.WeaveScopeResponse,
    requestSerialize: serialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    requestDeserialize: deserialize_prophetstor_api_datahub_weavescope_ListWeaveScopeContainersRequest,
    responseSerialize: serialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
    responseDeserialize: deserialize_prophetstor_api_datahub_weavescope_WeaveScopeResponse,
  },
};

exports.DatahubServiceClient = grpc.makeGenericClientConstructor(DatahubServiceService);
