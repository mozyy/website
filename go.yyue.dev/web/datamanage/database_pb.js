// source: datamanage/database.proto
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

var common_message_pb = require('../common/message_pb.js');
goog.object.extend(proto, common_message_pb);
var crawler_lianjia_pb = require('../crawler/lianjia_pb.js');
goog.object.extend(proto, crawler_lianjia_pb);
goog.exportSymbol('proto.database.ConnectRequest', null, global);
goog.exportSymbol('proto.database.InsertHouseRequest', null, global);
goog.exportSymbol('proto.database.InsertHouseSummaryRequest', null, global);
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
proto.database.ConnectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.database.ConnectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.database.ConnectRequest.displayName = 'proto.database.ConnectRequest';
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
proto.database.InsertHouseSummaryRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.database.InsertHouseSummaryRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.database.InsertHouseSummaryRequest.displayName = 'proto.database.InsertHouseSummaryRequest';
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
proto.database.InsertHouseRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.database.InsertHouseRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.database.InsertHouseRequest.displayName = 'proto.database.InsertHouseRequest';
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
proto.database.ConnectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.database.ConnectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.database.ConnectRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.database.ConnectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    database: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.database.ConnectRequest}
 */
proto.database.ConnectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.database.ConnectRequest;
  return proto.database.ConnectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.database.ConnectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.database.ConnectRequest}
 */
proto.database.ConnectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatabase(value);
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
proto.database.ConnectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.database.ConnectRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.database.ConnectRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.database.ConnectRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabase();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string database = 1;
 * @return {string}
 */
proto.database.ConnectRequest.prototype.getDatabase = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.database.ConnectRequest} returns this
 */
proto.database.ConnectRequest.prototype.setDatabase = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.database.InsertHouseSummaryRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.database.InsertHouseSummaryRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.database.InsertHouseSummaryRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.database.InsertHouseSummaryRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    database: jspb.Message.getFieldWithDefault(msg, 1, ""),
    table: jspb.Message.getFieldWithDefault(msg, 2, ""),
    house: (f = msg.getHouse()) && crawler_lianjia_pb.HouseSummary.toObject(includeInstance, f)
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
 * @return {!proto.database.InsertHouseSummaryRequest}
 */
proto.database.InsertHouseSummaryRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.database.InsertHouseSummaryRequest;
  return proto.database.InsertHouseSummaryRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.database.InsertHouseSummaryRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.database.InsertHouseSummaryRequest}
 */
proto.database.InsertHouseSummaryRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatabase(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTable(value);
      break;
    case 3:
      var value = new crawler_lianjia_pb.HouseSummary;
      reader.readMessage(value,crawler_lianjia_pb.HouseSummary.deserializeBinaryFromReader);
      msg.setHouse(value);
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
proto.database.InsertHouseSummaryRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.database.InsertHouseSummaryRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.database.InsertHouseSummaryRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.database.InsertHouseSummaryRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabase();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTable();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getHouse();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      crawler_lianjia_pb.HouseSummary.serializeBinaryToWriter
    );
  }
};


/**
 * optional string database = 1;
 * @return {string}
 */
proto.database.InsertHouseSummaryRequest.prototype.getDatabase = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.database.InsertHouseSummaryRequest} returns this
 */
proto.database.InsertHouseSummaryRequest.prototype.setDatabase = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string table = 2;
 * @return {string}
 */
proto.database.InsertHouseSummaryRequest.prototype.getTable = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.database.InsertHouseSummaryRequest} returns this
 */
proto.database.InsertHouseSummaryRequest.prototype.setTable = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional lianjia.HouseSummary house = 3;
 * @return {?proto.lianjia.HouseSummary}
 */
proto.database.InsertHouseSummaryRequest.prototype.getHouse = function() {
  return /** @type{?proto.lianjia.HouseSummary} */ (
    jspb.Message.getWrapperField(this, crawler_lianjia_pb.HouseSummary, 3));
};


/**
 * @param {?proto.lianjia.HouseSummary|undefined} value
 * @return {!proto.database.InsertHouseSummaryRequest} returns this
*/
proto.database.InsertHouseSummaryRequest.prototype.setHouse = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.database.InsertHouseSummaryRequest} returns this
 */
proto.database.InsertHouseSummaryRequest.prototype.clearHouse = function() {
  return this.setHouse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.database.InsertHouseSummaryRequest.prototype.hasHouse = function() {
  return jspb.Message.getField(this, 3) != null;
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
proto.database.InsertHouseRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.database.InsertHouseRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.database.InsertHouseRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.database.InsertHouseRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    database: jspb.Message.getFieldWithDefault(msg, 1, ""),
    table: jspb.Message.getFieldWithDefault(msg, 2, ""),
    house: (f = msg.getHouse()) && crawler_lianjia_pb.House.toObject(includeInstance, f)
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
 * @return {!proto.database.InsertHouseRequest}
 */
proto.database.InsertHouseRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.database.InsertHouseRequest;
  return proto.database.InsertHouseRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.database.InsertHouseRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.database.InsertHouseRequest}
 */
proto.database.InsertHouseRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatabase(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTable(value);
      break;
    case 3:
      var value = new crawler_lianjia_pb.House;
      reader.readMessage(value,crawler_lianjia_pb.House.deserializeBinaryFromReader);
      msg.setHouse(value);
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
proto.database.InsertHouseRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.database.InsertHouseRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.database.InsertHouseRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.database.InsertHouseRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabase();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTable();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getHouse();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      crawler_lianjia_pb.House.serializeBinaryToWriter
    );
  }
};


/**
 * optional string database = 1;
 * @return {string}
 */
proto.database.InsertHouseRequest.prototype.getDatabase = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.database.InsertHouseRequest} returns this
 */
proto.database.InsertHouseRequest.prototype.setDatabase = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string table = 2;
 * @return {string}
 */
proto.database.InsertHouseRequest.prototype.getTable = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.database.InsertHouseRequest} returns this
 */
proto.database.InsertHouseRequest.prototype.setTable = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional lianjia.House house = 3;
 * @return {?proto.lianjia.House}
 */
proto.database.InsertHouseRequest.prototype.getHouse = function() {
  return /** @type{?proto.lianjia.House} */ (
    jspb.Message.getWrapperField(this, crawler_lianjia_pb.House, 3));
};


/**
 * @param {?proto.lianjia.House|undefined} value
 * @return {!proto.database.InsertHouseRequest} returns this
*/
proto.database.InsertHouseRequest.prototype.setHouse = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.database.InsertHouseRequest} returns this
 */
proto.database.InsertHouseRequest.prototype.clearHouse = function() {
  return this.setHouse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.database.InsertHouseRequest.prototype.hasHouse = function() {
  return jspb.Message.getField(this, 3) != null;
};


goog.object.extend(exports, proto.database);
