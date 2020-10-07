// source: prophetstor/api/datahub/configs/configs.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var prophetstor_api_datahub_configs_scaler_pb = require('../../../../prophetstor/api/datahub/configs/scaler_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_configs_scaler_pb);
var prophetstor_api_datahub_configs_organization_pb = require('../../../../prophetstor/api/datahub/configs/organization_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_configs_organization_pb);
var prophetstor_api_datahub_configs_detection_pb = require('../../../../prophetstor/api/datahub/configs/detection_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_configs_detection_pb);
var prophetstor_api_datahub_configs_service_pb = require('../../../../prophetstor/api/datahub/configs/service_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_configs_service_pb);
var prophetstor_api_datahub_configs_notification_topic_pb = require('../../../../prophetstor/api/datahub/configs/notification_topic_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_configs_notification_topic_pb);
var prophetstor_api_datahub_configs_notification_channel_pb = require('../../../../prophetstor/api/datahub/configs/notification_channel_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_configs_notification_channel_pb);
goog.exportSymbol('proto.prophetstor.api.datahub.configs.Config', null, global);
goog.exportSymbol('proto.prophetstor.api.datahub.configs.Config.KindCase', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.prophetstor.api.datahub.configs.Config = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.prophetstor.api.datahub.configs.Config.oneofGroups_);
};
goog.inherits(proto.prophetstor.api.datahub.configs.Config, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.api.datahub.configs.Config.displayName = 'proto.prophetstor.api.datahub.configs.Config';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.prophetstor.api.datahub.configs.Config.oneofGroups_ = [[1,2,3,4,5,6]];

/**
 * @enum {number}
 */
proto.prophetstor.api.datahub.configs.Config.KindCase = {
  KIND_NOT_SET: 0,
  SCALER_CONFIG: 1,
  ORGANIZATION_CONFIG: 2,
  DETECTION_CONFIG: 3,
  SERVICE_CONFIG: 4,
  NOTIFICATION_CHANNEL_CONFIG: 5,
  NOTIFICATION_TOPIC_CONFIG: 6
};

/**
 * @return {proto.prophetstor.api.datahub.configs.Config.KindCase}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getKindCase = function() {
  return /** @type {proto.prophetstor.api.datahub.configs.Config.KindCase} */(jspb.Message.computeOneofCase(this, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.api.datahub.configs.Config.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.api.datahub.configs.Config} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.configs.Config.toObject = function(includeInstance, msg) {
  var f, obj = {
    scalerConfig: (f = msg.getScalerConfig()) && prophetstor_api_datahub_configs_scaler_pb.Scaler.toObject(includeInstance, f),
    organizationConfig: (f = msg.getOrganizationConfig()) && prophetstor_api_datahub_configs_organization_pb.Organization.toObject(includeInstance, f),
    detectionConfig: (f = msg.getDetectionConfig()) && prophetstor_api_datahub_configs_detection_pb.Detection.toObject(includeInstance, f),
    serviceConfig: (f = msg.getServiceConfig()) && prophetstor_api_datahub_configs_service_pb.Service.toObject(includeInstance, f),
    notificationChannelConfig: (f = msg.getNotificationChannelConfig()) && prophetstor_api_datahub_configs_notification_channel_pb.NotificationChannel.toObject(includeInstance, f),
    notificationTopicConfig: (f = msg.getNotificationTopicConfig()) && prophetstor_api_datahub_configs_notification_topic_pb.NotificationTopic.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.prophetstor.api.datahub.configs.Config}
 */
proto.prophetstor.api.datahub.configs.Config.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.api.datahub.configs.Config;
  return proto.prophetstor.api.datahub.configs.Config.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.api.datahub.configs.Config} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.api.datahub.configs.Config}
 */
proto.prophetstor.api.datahub.configs.Config.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new prophetstor_api_datahub_configs_scaler_pb.Scaler;
      reader.readMessage(value,prophetstor_api_datahub_configs_scaler_pb.Scaler.deserializeBinaryFromReader);
      msg.setScalerConfig(value);
      break;
    case 2:
      var value = new prophetstor_api_datahub_configs_organization_pb.Organization;
      reader.readMessage(value,prophetstor_api_datahub_configs_organization_pb.Organization.deserializeBinaryFromReader);
      msg.setOrganizationConfig(value);
      break;
    case 3:
      var value = new prophetstor_api_datahub_configs_detection_pb.Detection;
      reader.readMessage(value,prophetstor_api_datahub_configs_detection_pb.Detection.deserializeBinaryFromReader);
      msg.setDetectionConfig(value);
      break;
    case 4:
      var value = new prophetstor_api_datahub_configs_service_pb.Service;
      reader.readMessage(value,prophetstor_api_datahub_configs_service_pb.Service.deserializeBinaryFromReader);
      msg.setServiceConfig(value);
      break;
    case 5:
      var value = new prophetstor_api_datahub_configs_notification_channel_pb.NotificationChannel;
      reader.readMessage(value,prophetstor_api_datahub_configs_notification_channel_pb.NotificationChannel.deserializeBinaryFromReader);
      msg.setNotificationChannelConfig(value);
      break;
    case 6:
      var value = new prophetstor_api_datahub_configs_notification_topic_pb.NotificationTopic;
      reader.readMessage(value,prophetstor_api_datahub_configs_notification_topic_pb.NotificationTopic.deserializeBinaryFromReader);
      msg.setNotificationTopicConfig(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.api.datahub.configs.Config.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.api.datahub.configs.Config} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.configs.Config.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getScalerConfig();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      prophetstor_api_datahub_configs_scaler_pb.Scaler.serializeBinaryToWriter
    );
  }
  f = message.getOrganizationConfig();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      prophetstor_api_datahub_configs_organization_pb.Organization.serializeBinaryToWriter
    );
  }
  f = message.getDetectionConfig();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      prophetstor_api_datahub_configs_detection_pb.Detection.serializeBinaryToWriter
    );
  }
  f = message.getServiceConfig();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      prophetstor_api_datahub_configs_service_pb.Service.serializeBinaryToWriter
    );
  }
  f = message.getNotificationChannelConfig();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      prophetstor_api_datahub_configs_notification_channel_pb.NotificationChannel.serializeBinaryToWriter
    );
  }
  f = message.getNotificationTopicConfig();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      prophetstor_api_datahub_configs_notification_topic_pb.NotificationTopic.serializeBinaryToWriter
    );
  }
};


/**
 * optional Scaler scaler_config = 1;
 * @return {?proto.prophetstor.api.datahub.configs.Scaler}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getScalerConfig = function() {
  return /** @type{?proto.prophetstor.api.datahub.configs.Scaler} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_configs_scaler_pb.Scaler, 1));
};


/**
 * @param {?proto.prophetstor.api.datahub.configs.Scaler|undefined} value
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
*/
proto.prophetstor.api.datahub.configs.Config.prototype.setScalerConfig = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
 */
proto.prophetstor.api.datahub.configs.Config.prototype.clearScalerConfig = function() {
  return this.setScalerConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.hasScalerConfig = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Organization organization_config = 2;
 * @return {?proto.prophetstor.api.datahub.configs.Organization}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getOrganizationConfig = function() {
  return /** @type{?proto.prophetstor.api.datahub.configs.Organization} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_configs_organization_pb.Organization, 2));
};


/**
 * @param {?proto.prophetstor.api.datahub.configs.Organization|undefined} value
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
*/
proto.prophetstor.api.datahub.configs.Config.prototype.setOrganizationConfig = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
 */
proto.prophetstor.api.datahub.configs.Config.prototype.clearOrganizationConfig = function() {
  return this.setOrganizationConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.hasOrganizationConfig = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Detection detection_config = 3;
 * @return {?proto.prophetstor.api.datahub.configs.Detection}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getDetectionConfig = function() {
  return /** @type{?proto.prophetstor.api.datahub.configs.Detection} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_configs_detection_pb.Detection, 3));
};


/**
 * @param {?proto.prophetstor.api.datahub.configs.Detection|undefined} value
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
*/
proto.prophetstor.api.datahub.configs.Config.prototype.setDetectionConfig = function(value) {
  return jspb.Message.setOneofWrapperField(this, 3, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
 */
proto.prophetstor.api.datahub.configs.Config.prototype.clearDetectionConfig = function() {
  return this.setDetectionConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.hasDetectionConfig = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Service service_config = 4;
 * @return {?proto.prophetstor.api.datahub.configs.Service}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getServiceConfig = function() {
  return /** @type{?proto.prophetstor.api.datahub.configs.Service} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_configs_service_pb.Service, 4));
};


/**
 * @param {?proto.prophetstor.api.datahub.configs.Service|undefined} value
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
*/
proto.prophetstor.api.datahub.configs.Config.prototype.setServiceConfig = function(value) {
  return jspb.Message.setOneofWrapperField(this, 4, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
 */
proto.prophetstor.api.datahub.configs.Config.prototype.clearServiceConfig = function() {
  return this.setServiceConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.hasServiceConfig = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional NotificationChannel notification_channel_config = 5;
 * @return {?proto.prophetstor.api.datahub.configs.NotificationChannel}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getNotificationChannelConfig = function() {
  return /** @type{?proto.prophetstor.api.datahub.configs.NotificationChannel} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_configs_notification_channel_pb.NotificationChannel, 5));
};


/**
 * @param {?proto.prophetstor.api.datahub.configs.NotificationChannel|undefined} value
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
*/
proto.prophetstor.api.datahub.configs.Config.prototype.setNotificationChannelConfig = function(value) {
  return jspb.Message.setOneofWrapperField(this, 5, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
 */
proto.prophetstor.api.datahub.configs.Config.prototype.clearNotificationChannelConfig = function() {
  return this.setNotificationChannelConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.hasNotificationChannelConfig = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional NotificationTopic notification_topic_config = 6;
 * @return {?proto.prophetstor.api.datahub.configs.NotificationTopic}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.getNotificationTopicConfig = function() {
  return /** @type{?proto.prophetstor.api.datahub.configs.NotificationTopic} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_configs_notification_topic_pb.NotificationTopic, 6));
};


/**
 * @param {?proto.prophetstor.api.datahub.configs.NotificationTopic|undefined} value
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
*/
proto.prophetstor.api.datahub.configs.Config.prototype.setNotificationTopicConfig = function(value) {
  return jspb.Message.setOneofWrapperField(this, 6, proto.prophetstor.api.datahub.configs.Config.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.configs.Config} returns this
 */
proto.prophetstor.api.datahub.configs.Config.prototype.clearNotificationTopicConfig = function() {
  return this.setNotificationTopicConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.configs.Config.prototype.hasNotificationTopicConfig = function() {
  return jspb.Message.getField(this, 6) != null;
};


goog.object.extend(exports, proto.prophetstor.api.datahub.configs);
