// source: prophetstor/api/datahub/keycodes/keycodes.proto
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

var prophetstor_api_datahub_keycodes_types_pb = require('../../../../prophetstor/api/datahub/keycodes/types_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_keycodes_types_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.prophetstor.api.datahub.keycodes.Keycode', null, global);
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
proto.prophetstor.api.datahub.keycodes.Keycode = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.api.datahub.keycodes.Keycode, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.api.datahub.keycodes.Keycode.displayName = 'proto.prophetstor.api.datahub.keycodes.Keycode';
}



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
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.api.datahub.keycodes.Keycode.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.api.datahub.keycodes.Keycode} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.keycodes.Keycode.toObject = function(includeInstance, msg) {
  var f, obj = {
    keycode: jspb.Message.getFieldWithDefault(msg, 1, ""),
    keycodeType: jspb.Message.getFieldWithDefault(msg, 2, ""),
    keycodeVersion: jspb.Message.getFieldWithDefault(msg, 3, 0),
    applyTime: (f = msg.getApplyTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    expireTime: (f = msg.getExpireTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    licenseState: jspb.Message.getFieldWithDefault(msg, 6, ""),
    registered: jspb.Message.getBooleanFieldWithDefault(msg, 7, false),
    capacity: (f = msg.getCapacity()) && prophetstor_api_datahub_keycodes_types_pb.Capacity.toObject(includeInstance, f),
    functionality: (f = msg.getFunctionality()) && prophetstor_api_datahub_keycodes_types_pb.Functionality.toObject(includeInstance, f),
    retention: (f = msg.getRetention()) && prophetstor_api_datahub_keycodes_types_pb.Retention.toObject(includeInstance, f),
    serviceAgreement: (f = msg.getServiceAgreement()) && prophetstor_api_datahub_keycodes_types_pb.ServiceAgreement.toObject(includeInstance, f)
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
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.api.datahub.keycodes.Keycode;
  return proto.prophetstor.api.datahub.keycodes.Keycode.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.api.datahub.keycodes.Keycode} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setKeycode(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setKeycodeType(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setKeycodeVersion(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setApplyTime(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setExpireTime(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setLicenseState(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRegistered(value);
      break;
    case 8:
      var value = new prophetstor_api_datahub_keycodes_types_pb.Capacity;
      reader.readMessage(value,prophetstor_api_datahub_keycodes_types_pb.Capacity.deserializeBinaryFromReader);
      msg.setCapacity(value);
      break;
    case 9:
      var value = new prophetstor_api_datahub_keycodes_types_pb.Functionality;
      reader.readMessage(value,prophetstor_api_datahub_keycodes_types_pb.Functionality.deserializeBinaryFromReader);
      msg.setFunctionality(value);
      break;
    case 10:
      var value = new prophetstor_api_datahub_keycodes_types_pb.Retention;
      reader.readMessage(value,prophetstor_api_datahub_keycodes_types_pb.Retention.deserializeBinaryFromReader);
      msg.setRetention(value);
      break;
    case 11:
      var value = new prophetstor_api_datahub_keycodes_types_pb.ServiceAgreement;
      reader.readMessage(value,prophetstor_api_datahub_keycodes_types_pb.ServiceAgreement.deserializeBinaryFromReader);
      msg.setServiceAgreement(value);
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
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.api.datahub.keycodes.Keycode.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.api.datahub.keycodes.Keycode} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.keycodes.Keycode.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKeycode();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getKeycodeType();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getKeycodeVersion();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getApplyTime();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getExpireTime();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getLicenseState();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getRegistered();
  if (f) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getCapacity();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      prophetstor_api_datahub_keycodes_types_pb.Capacity.serializeBinaryToWriter
    );
  }
  f = message.getFunctionality();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      prophetstor_api_datahub_keycodes_types_pb.Functionality.serializeBinaryToWriter
    );
  }
  f = message.getRetention();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      prophetstor_api_datahub_keycodes_types_pb.Retention.serializeBinaryToWriter
    );
  }
  f = message.getServiceAgreement();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      prophetstor_api_datahub_keycodes_types_pb.ServiceAgreement.serializeBinaryToWriter
    );
  }
};


/**
 * optional string keycode = 1;
 * @return {string}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getKeycode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setKeycode = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string keycode_type = 2;
 * @return {string}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getKeycodeType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setKeycodeType = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 keycode_version = 3;
 * @return {number}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getKeycodeVersion = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setKeycodeVersion = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp apply_time = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getApplyTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
*/
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setApplyTime = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.clearApplyTime = function() {
  return this.setApplyTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.hasApplyTime = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.Timestamp expire_time = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getExpireTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
*/
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setExpireTime = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.clearExpireTime = function() {
  return this.setExpireTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.hasExpireTime = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string license_state = 6;
 * @return {string}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getLicenseState = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setLicenseState = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional bool registered = 7;
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getRegistered = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 7, false));
};


/**
 * @param {boolean} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setRegistered = function(value) {
  return jspb.Message.setProto3BooleanField(this, 7, value);
};


/**
 * optional Capacity capacity = 8;
 * @return {?proto.prophetstor.api.datahub.keycodes.Capacity}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getCapacity = function() {
  return /** @type{?proto.prophetstor.api.datahub.keycodes.Capacity} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_keycodes_types_pb.Capacity, 8));
};


/**
 * @param {?proto.prophetstor.api.datahub.keycodes.Capacity|undefined} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
*/
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setCapacity = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.clearCapacity = function() {
  return this.setCapacity(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.hasCapacity = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional Functionality functionality = 9;
 * @return {?proto.prophetstor.api.datahub.keycodes.Functionality}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getFunctionality = function() {
  return /** @type{?proto.prophetstor.api.datahub.keycodes.Functionality} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_keycodes_types_pb.Functionality, 9));
};


/**
 * @param {?proto.prophetstor.api.datahub.keycodes.Functionality|undefined} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
*/
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setFunctionality = function(value) {
  return jspb.Message.setWrapperField(this, 9, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.clearFunctionality = function() {
  return this.setFunctionality(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.hasFunctionality = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional Retention retention = 10;
 * @return {?proto.prophetstor.api.datahub.keycodes.Retention}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getRetention = function() {
  return /** @type{?proto.prophetstor.api.datahub.keycodes.Retention} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_keycodes_types_pb.Retention, 10));
};


/**
 * @param {?proto.prophetstor.api.datahub.keycodes.Retention|undefined} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
*/
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setRetention = function(value) {
  return jspb.Message.setWrapperField(this, 10, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.clearRetention = function() {
  return this.setRetention(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.hasRetention = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional ServiceAgreement service_agreement = 11;
 * @return {?proto.prophetstor.api.datahub.keycodes.ServiceAgreement}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.getServiceAgreement = function() {
  return /** @type{?proto.prophetstor.api.datahub.keycodes.ServiceAgreement} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_keycodes_types_pb.ServiceAgreement, 11));
};


/**
 * @param {?proto.prophetstor.api.datahub.keycodes.ServiceAgreement|undefined} value
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
*/
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.setServiceAgreement = function(value) {
  return jspb.Message.setWrapperField(this, 11, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.keycodes.Keycode} returns this
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.clearServiceAgreement = function() {
  return this.setServiceAgreement(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.keycodes.Keycode.prototype.hasServiceAgreement = function() {
  return jspb.Message.getField(this, 11) != null;
};


goog.object.extend(exports, proto.prophetstor.api.datahub.keycodes);
