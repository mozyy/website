/**
 * @fileoverview gRPC-Web generated client stub for user
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as common_message_pb from '../common/message_pb';

import {
  RegisterRequest,
  SendRequest,
  ValidateRequest} from './user_pb';

export class SMSServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSend = new grpcWeb.AbstractClientBase.MethodInfo(
    common_message_pb.Message,
    (request: SendRequest) => {
      return request.serializeBinary();
    },
    common_message_pb.Message.deserializeBinary
  );

  send(
    request: SendRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_message_pb.Message) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/user.SMSService/Send',
      request,
      metadata || {},
      this.methodInfoSend,
      callback);
  }

  methodInfoValidate = new grpcWeb.AbstractClientBase.MethodInfo(
    common_message_pb.Message,
    (request: ValidateRequest) => {
      return request.serializeBinary();
    },
    common_message_pb.Message.deserializeBinary
  );

  validate(
    request: ValidateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_message_pb.Message) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/user.SMSService/Validate',
      request,
      metadata || {},
      this.methodInfoValidate,
      callback);
  }

}

export class UserServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoRegister = new grpcWeb.AbstractClientBase.MethodInfo(
    common_message_pb.Message,
    (request: RegisterRequest) => {
      return request.serializeBinary();
    },
    common_message_pb.Message.deserializeBinary
  );

  register(
    request: RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_message_pb.Message) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/user.UserService/Register',
      request,
      metadata || {},
      this.methodInfoRegister,
      callback);
  }

}

