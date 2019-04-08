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

var common_pb = require('./common_pb.js');
goog.exportSymbol('proto.base.SenderInfo', null, global);
goog.exportSymbol('proto.base.SubscriberRequest', null, global);
goog.exportSymbol('proto.base.Value', null, global);

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
proto.base.SenderInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.base.SenderInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.base.SenderInfo.displayName = 'proto.base.SenderInfo';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.base.SenderInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.base.SenderInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.base.SenderInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.base.SenderInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    clientid: (f = msg.getClientid()) && common_pb.ClientId.toObject(includeInstance, f),
    value: (f = msg.getValue()) && proto.base.Value.toObject(includeInstance, f),
    frequency: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.base.SenderInfo}
 */
proto.base.SenderInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.base.SenderInfo;
  return proto.base.SenderInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.base.SenderInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.base.SenderInfo}
 */
proto.base.SenderInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new common_pb.ClientId;
      reader.readMessage(value,common_pb.ClientId.deserializeBinaryFromReader);
      msg.setClientid(value);
      break;
    case 2:
      var value = new proto.base.Value;
      reader.readMessage(value,proto.base.Value.deserializeBinaryFromReader);
      msg.setValue(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFrequency(value);
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
proto.base.SenderInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.base.SenderInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.base.SenderInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.base.SenderInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClientid();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      common_pb.ClientId.serializeBinaryToWriter
    );
  }
  f = message.getValue();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.base.Value.serializeBinaryToWriter
    );
  }
  f = message.getFrequency();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional ClientId clientId = 1;
 * @return {?proto.base.ClientId}
 */
proto.base.SenderInfo.prototype.getClientid = function() {
  return /** @type{?proto.base.ClientId} */ (
    jspb.Message.getWrapperField(this, common_pb.ClientId, 1));
};


/** @param {?proto.base.ClientId|undefined} value */
proto.base.SenderInfo.prototype.setClientid = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.base.SenderInfo.prototype.clearClientid = function() {
  this.setClientid(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.base.SenderInfo.prototype.hasClientid = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Value value = 2;
 * @return {?proto.base.Value}
 */
proto.base.SenderInfo.prototype.getValue = function() {
  return /** @type{?proto.base.Value} */ (
    jspb.Message.getWrapperField(this, proto.base.Value, 2));
};


/** @param {?proto.base.Value|undefined} value */
proto.base.SenderInfo.prototype.setValue = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.base.SenderInfo.prototype.clearValue = function() {
  this.setValue(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.base.SenderInfo.prototype.hasValue = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional int32 frequency = 3;
 * @return {number}
 */
proto.base.SenderInfo.prototype.getFrequency = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.base.SenderInfo.prototype.setFrequency = function(value) {
  jspb.Message.setField(this, 3, value);
};



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
proto.base.SubscriberRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.base.SubscriberRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.base.SubscriberRequest.displayName = 'proto.base.SubscriberRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.base.SubscriberRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.base.SubscriberRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.base.SubscriberRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.base.SubscriberRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    clientid: (f = msg.getClientid()) && common_pb.ClientId.toObject(includeInstance, f),
    onchange: jspb.Message.getFieldWithDefault(msg, 2, false)
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
 * @return {!proto.base.SubscriberRequest}
 */
proto.base.SubscriberRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.base.SubscriberRequest;
  return proto.base.SubscriberRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.base.SubscriberRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.base.SubscriberRequest}
 */
proto.base.SubscriberRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new common_pb.ClientId;
      reader.readMessage(value,common_pb.ClientId.deserializeBinaryFromReader);
      msg.setClientid(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setOnchange(value);
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
proto.base.SubscriberRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.base.SubscriberRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.base.SubscriberRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.base.SubscriberRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClientid();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      common_pb.ClientId.serializeBinaryToWriter
    );
  }
  f = message.getOnchange();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
};


/**
 * optional ClientId clientId = 1;
 * @return {?proto.base.ClientId}
 */
proto.base.SubscriberRequest.prototype.getClientid = function() {
  return /** @type{?proto.base.ClientId} */ (
    jspb.Message.getWrapperField(this, common_pb.ClientId, 1));
};


/** @param {?proto.base.ClientId|undefined} value */
proto.base.SubscriberRequest.prototype.setClientid = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.base.SubscriberRequest.prototype.clearClientid = function() {
  this.setClientid(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.base.SubscriberRequest.prototype.hasClientid = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional bool onChange = 2;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.base.SubscriberRequest.prototype.getOnchange = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 2, false));
};


/** @param {boolean} value */
proto.base.SubscriberRequest.prototype.setOnchange = function(value) {
  jspb.Message.setField(this, 2, value);
};



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
proto.base.Value = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.base.Value, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.base.Value.displayName = 'proto.base.Value';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.base.Value.prototype.toObject = function(opt_includeInstance) {
  return proto.base.Value.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.base.Value} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.base.Value.toObject = function(includeInstance, msg) {
  var f, obj = {
    payload: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.base.Value}
 */
proto.base.Value.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.base.Value;
  return proto.base.Value.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.base.Value} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.base.Value}
 */
proto.base.Value.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPayload(value);
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
proto.base.Value.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.base.Value.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.base.Value} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.base.Value.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPayload();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
};


/**
 * optional int32 payload = 1;
 * @return {number}
 */
proto.base.Value.prototype.getPayload = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.base.Value.prototype.setPayload = function(value) {
  jspb.Message.setField(this, 1, value);
};


goog.object.extend(exports, proto.base);
