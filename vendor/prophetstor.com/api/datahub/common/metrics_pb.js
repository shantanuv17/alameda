// source: prophetstor/api/datahub/common/metrics.proto
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

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.prophetstor.api.datahub.common.MetricData', null, global);
goog.exportSymbol('proto.prophetstor.api.datahub.common.MetricType', null, global);
goog.exportSymbol('proto.prophetstor.api.datahub.common.ResourceName', null, global);
goog.exportSymbol('proto.prophetstor.api.datahub.common.Sample', null, global);
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
proto.prophetstor.api.datahub.common.Sample = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.api.datahub.common.Sample, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.api.datahub.common.Sample.displayName = 'proto.prophetstor.api.datahub.common.Sample';
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
proto.prophetstor.api.datahub.common.MetricData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.api.datahub.common.MetricData.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.api.datahub.common.MetricData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.api.datahub.common.MetricData.displayName = 'proto.prophetstor.api.datahub.common.MetricData';
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
proto.prophetstor.api.datahub.common.Sample.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.api.datahub.common.Sample.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.api.datahub.common.Sample} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.common.Sample.toObject = function(includeInstance, msg) {
  var f, obj = {
    time: (f = msg.getTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    endTime: (f = msg.getEndTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    numValue: jspb.Message.getFieldWithDefault(msg, 3, "")
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
 * @return {!proto.prophetstor.api.datahub.common.Sample}
 */
proto.prophetstor.api.datahub.common.Sample.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.api.datahub.common.Sample;
  return proto.prophetstor.api.datahub.common.Sample.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.api.datahub.common.Sample} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.api.datahub.common.Sample}
 */
proto.prophetstor.api.datahub.common.Sample.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTime(value);
      break;
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setEndTime(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setNumValue(value);
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
proto.prophetstor.api.datahub.common.Sample.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.api.datahub.common.Sample.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.api.datahub.common.Sample} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.common.Sample.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTime();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getEndTime();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getNumValue();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional google.protobuf.Timestamp time = 1;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.prophetstor.api.datahub.common.Sample.prototype.getTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 1));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.prophetstor.api.datahub.common.Sample} returns this
*/
proto.prophetstor.api.datahub.common.Sample.prototype.setTime = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.common.Sample} returns this
 */
proto.prophetstor.api.datahub.common.Sample.prototype.clearTime = function() {
  return this.setTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.common.Sample.prototype.hasTime = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.Timestamp end_time = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.prophetstor.api.datahub.common.Sample.prototype.getEndTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.prophetstor.api.datahub.common.Sample} returns this
*/
proto.prophetstor.api.datahub.common.Sample.prototype.setEndTime = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.api.datahub.common.Sample} returns this
 */
proto.prophetstor.api.datahub.common.Sample.prototype.clearEndTime = function() {
  return this.setEndTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.api.datahub.common.Sample.prototype.hasEndTime = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional string num_value = 3;
 * @return {string}
 */
proto.prophetstor.api.datahub.common.Sample.prototype.getNumValue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.api.datahub.common.Sample} returns this
 */
proto.prophetstor.api.datahub.common.Sample.prototype.setNumValue = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.api.datahub.common.MetricData.repeatedFields_ = [2];



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
proto.prophetstor.api.datahub.common.MetricData.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.api.datahub.common.MetricData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.api.datahub.common.MetricData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.common.MetricData.toObject = function(includeInstance, msg) {
  var f, obj = {
    metricType: jspb.Message.getFieldWithDefault(msg, 1, 0),
    dataList: jspb.Message.toObjectList(msg.getDataList(),
    proto.prophetstor.api.datahub.common.Sample.toObject, includeInstance),
    granularity: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.prophetstor.api.datahub.common.MetricData}
 */
proto.prophetstor.api.datahub.common.MetricData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.api.datahub.common.MetricData;
  return proto.prophetstor.api.datahub.common.MetricData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.api.datahub.common.MetricData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.api.datahub.common.MetricData}
 */
proto.prophetstor.api.datahub.common.MetricData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.prophetstor.api.datahub.common.MetricType} */ (reader.readEnum());
      msg.setMetricType(value);
      break;
    case 2:
      var value = new proto.prophetstor.api.datahub.common.Sample;
      reader.readMessage(value,proto.prophetstor.api.datahub.common.Sample.deserializeBinaryFromReader);
      msg.addData(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGranularity(value);
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
proto.prophetstor.api.datahub.common.MetricData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.api.datahub.common.MetricData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.api.datahub.common.MetricData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.api.datahub.common.MetricData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMetricType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getDataList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.prophetstor.api.datahub.common.Sample.serializeBinaryToWriter
    );
  }
  f = message.getGranularity();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
};


/**
 * optional MetricType metric_type = 1;
 * @return {!proto.prophetstor.api.datahub.common.MetricType}
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.getMetricType = function() {
  return /** @type {!proto.prophetstor.api.datahub.common.MetricType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.prophetstor.api.datahub.common.MetricType} value
 * @return {!proto.prophetstor.api.datahub.common.MetricData} returns this
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.setMetricType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * repeated Sample data = 2;
 * @return {!Array<!proto.prophetstor.api.datahub.common.Sample>}
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.getDataList = function() {
  return /** @type{!Array<!proto.prophetstor.api.datahub.common.Sample>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.prophetstor.api.datahub.common.Sample, 2));
};


/**
 * @param {!Array<!proto.prophetstor.api.datahub.common.Sample>} value
 * @return {!proto.prophetstor.api.datahub.common.MetricData} returns this
*/
proto.prophetstor.api.datahub.common.MetricData.prototype.setDataList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.api.datahub.common.Sample=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.api.datahub.common.Sample}
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.addData = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.api.datahub.common.Sample, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.api.datahub.common.MetricData} returns this
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.clearDataList = function() {
  return this.setDataList([]);
};


/**
 * optional int64 granularity = 3;
 * @return {number}
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.getGranularity = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.prophetstor.api.datahub.common.MetricData} returns this
 */
proto.prophetstor.api.datahub.common.MetricData.prototype.setGranularity = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * @enum {number}
 */
proto.prophetstor.api.datahub.common.MetricType = {
  METRICS_TYPE_UNDEFINED: 0,
  CPU_SECONDS_TOTAL: 1,
  CPU_CORES_ALLOCATABLE: 2,
  CPU_MILLICORES_TOTAL: 3,
  CPU_MILLICORES_AVAIL: 4,
  CPU_MILLICORES_USAGE: 5,
  CPU_MILLICORES_USAGE_PCT: 6,
  MEMORY_BYTES_ALLOCATABLE: 7,
  MEMORY_BYTES_TOTAL: 8,
  MEMORY_BYTES_AVAIL: 9,
  MEMORY_BYTES_USAGE: 10,
  MEMORY_BYTES_USAGE_PCT: 11,
  FS_BYTES_TOTAL: 12,
  FS_BYTES_AVAIL: 13,
  FS_BYTES_USAGE: 14,
  FS_BYTES_USAGE_PCT: 15,
  HTTP_REQUESTS_COUNT: 16,
  HTTP_REQUESTS_TOTAL: 17,
  HTTP_RESPONSE_COUNT: 18,
  HTTP_RESPONSE_TOTAL: 19,
  DISK_IO_SECONDS_TOTAL: 20,
  DISK_IO_UTILIZATION: 21,
  RESTARTS_TOTAL: 22,
  UNSCHEDULABLE: 23,
  HEALTH: 24,
  POWER_USAGE_WATTS: 25,
  TEMPERATURE_CELSIUS: 26,
  DUTY_CYCLE: 27,
  CURRENT_OFFSET: 28,
  LAG: 29,
  LATENCY: 30,
  NUMBER: 31
};

/**
 * @enum {number}
 */
proto.prophetstor.api.datahub.common.ResourceName = {
  RESOURCE_NAME_UNDEFINED: 0,
  CPU: 1,
  MEMORY: 2
};

goog.object.extend(exports, proto.prophetstor.api.datahub.common);
