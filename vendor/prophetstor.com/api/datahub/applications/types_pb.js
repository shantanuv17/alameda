// source: prophetstor/api/datahub/applications/types.proto
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

var prophetstor_api_datahub_common_rawdata_pb = require('../../../../prophetstor/api/datahub/common/rawdata_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_common_rawdata_pb);
var prophetstor_api_datahub_schemas_types_pb = require('../../../../prophetstor/api/datahub/schemas/types_pb.js');
goog.object.extend(proto, prophetstor_api_datahub_schemas_types_pb);
goog.exportSymbol('proto.prophetstor.api.datahub.applications.Application', null, global);
goog.exportSymbol('proto.prophetstor.api.datahub.applications.ApplicationData', null, global);
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
proto.prophetstor.api.datahub.applications.Application = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.api.datahub.applications.Application.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.api.datahub.applications.Application, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.api.datahub.applications.Application.displayName = 'proto.prophetstor.api.datahub.applications.Application';
}
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
proto.prophetstor.api.datahub.applications.ApplicationData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.api.datahub.applications.ApplicationData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.api.datahub.applications.ApplicationData.displayName = 'proto.prophetstor.api.datahub.applications.ApplicationData';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.api.datahub.applications.Application.repeatedFields_ = [2];



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
proto.prophetstor.api.datahub.applications.Application.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.api.datahub.applications.Application.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.api.datahub.applications.Application} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.applications.Application.toObject = function(includeInstance, msg) {
  var f, obj = {
    schemaMeta: (f = msg.getSchemaMeta()) && prophetstor_api_datahub_schemas_types_pb.SchemaMeta.toObject(includeInstance, f),
    applicationDataList: jspb.Message.toObjectList(msg.getApplicationDataList(),
    proto.prophetstor.api.datahub.applications.ApplicationData.toObject, includeInstance)
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
 * @return {!proto.prophetstor.api.datahub.applications.Application}
 */
proto.prophetstor.api.datahub.applications.Application.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.api.datahub.applications.Application;
  return proto.prophetstor.api.datahub.applications.Application.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.api.datahub.applications.Application} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.api.datahub.applications.Application}
 */
proto.prophetstor.api.datahub.applications.Application.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new prophetstor_api_datahub_schemas_types_pb.SchemaMeta;
      reader.readMessage(value,prophetstor_api_datahub_schemas_types_pb.SchemaMeta.deserializeBinaryFromReader);
      msg.setSchemaMeta(value);
      break;
    case 2:
      var value = new proto.prophetstor.api.datahub.applications.ApplicationData;
      reader.readMessage(value,proto.prophetstor.api.datahub.applications.ApplicationData.deserializeBinaryFromReader);
      msg.addApplicationData(value);
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
proto.prophetstor.api.datahub.applications.Application.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.api.datahub.applications.Application.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.api.datahub.applications.Application} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.applications.Application.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSchemaMeta();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      prophetstor_api_datahub_schemas_types_pb.SchemaMeta.serializeBinaryToWriter
    );
  }
  f = message.getApplicationDataList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.prophetstor.api.datahub.applications.ApplicationData.serializeBinaryToWriter
    );
  }
};


/**
 * optional prophetstor.api.datahub.schemas.SchemaMeta schema_meta = 1;
 * @return {?proto.prophetstor.api.datahub.schemas.SchemaMeta}
 */
proto.prophetstor.api.datahub.applications.Application.prototype.getSchemaMeta = function() {
  return /** @type{?proto.prophetstor.api.datahub.schemas.SchemaMeta} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_schemas_types_pb.SchemaMeta, 1));
};


/**
 * @param {?proto.prophetstor.api.datahub.schemas.SchemaMeta|undefined} value
 * @return {!proto.prophetstor.api.datahub.applications.Application} returns this
*/
proto.prophetstor.api.datahub.applications.Application.prototype.setSchemaMeta = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.applications.Application} returns this
 */
proto.prophetstor.api.datahub.applications.Application.prototype.clearSchemaMeta = function() {
  return this.setSchemaMeta(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.applications.Application.prototype.hasSchemaMeta = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated ApplicationData application_data = 2;
 * @return {!Array<!proto.prophetstor.api.datahub.applications.ApplicationData>}
 */
proto.prophetstor.api.datahub.applications.Application.prototype.getApplicationDataList = function() {
  return /** @type{!Array<!proto.prophetstor.api.datahub.applications.ApplicationData>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.prophetstor.api.datahub.applications.ApplicationData, 2));
};


/**
 * @param {!Array<!proto.prophetstor.api.datahub.applications.ApplicationData>} value
 * @return {!proto.prophetstor.api.datahub.applications.Application} returns this
*/
proto.prophetstor.api.datahub.applications.Application.prototype.setApplicationDataList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.api.datahub.applications.ApplicationData=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.api.datahub.applications.ApplicationData}
 */
proto.prophetstor.api.datahub.applications.Application.prototype.addApplicationData = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.api.datahub.applications.ApplicationData, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.api.datahub.applications.Application} returns this
 */
proto.prophetstor.api.datahub.applications.Application.prototype.clearApplicationDataList = function() {
  return this.setApplicationDataList([]);
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
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.api.datahub.applications.ApplicationData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.api.datahub.applications.ApplicationData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.applications.ApplicationData.toObject = function(includeInstance, msg) {
  var f, obj = {
    measurement: jspb.Message.getFieldWithDefault(msg, 1, ""),
    readData: (f = msg.getReadData()) && prophetstor_api_datahub_common_rawdata_pb.ReadData.toObject(includeInstance, f)
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
 * @return {!proto.prophetstor.api.datahub.applications.ApplicationData}
 */
proto.prophetstor.api.datahub.applications.ApplicationData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.api.datahub.applications.ApplicationData;
  return proto.prophetstor.api.datahub.applications.ApplicationData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.api.datahub.applications.ApplicationData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.api.datahub.applications.ApplicationData}
 */
proto.prophetstor.api.datahub.applications.ApplicationData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMeasurement(value);
      break;
    case 2:
      var value = new prophetstor_api_datahub_common_rawdata_pb.ReadData;
      reader.readMessage(value,prophetstor_api_datahub_common_rawdata_pb.ReadData.deserializeBinaryFromReader);
      msg.setReadData(value);
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
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.api.datahub.applications.ApplicationData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.api.datahub.applications.ApplicationData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.applications.ApplicationData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMeasurement();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getReadData();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      prophetstor_api_datahub_common_rawdata_pb.ReadData.serializeBinaryToWriter
    );
  }
};


/**
 * optional string measurement = 1;
 * @return {string}
 */
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.getMeasurement = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.api.datahub.applications.ApplicationData} returns this
 */
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.setMeasurement = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional prophetstor.api.datahub.common.ReadData read_data = 2;
 * @return {?proto.prophetstor.api.datahub.common.ReadData}
 */
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.getReadData = function() {
  return /** @type{?proto.prophetstor.api.datahub.common.ReadData} */ (
    jspb.Message.getWrapperField(this, prophetstor_api_datahub_common_rawdata_pb.ReadData, 2));
};


/**
 * @param {?proto.prophetstor.api.datahub.common.ReadData|undefined} value
 * @return {!proto.prophetstor.api.datahub.applications.ApplicationData} returns this
*/
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.setReadData = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.applications.ApplicationData} returns this
 */
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.clearReadData = function() {
  return this.setReadData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.applications.ApplicationData.prototype.hasReadData = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.prophetstor.api.datahub.applications);
