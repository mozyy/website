/**
 * @fileoverview gRPC-Web generated client stub for database
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as common_message_pb from '../common/message_pb';
import * as crawler_lianjia_pb from '../crawler/lianjia_pb';

import {
  ConnectRequest,
  InsertHouseRequest} from './database_pb';

export class DatabaseServiceClient {
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

  methodInfoConnect = new grpcWeb.AbstractClientBase.MethodInfo(
    common_message_pb.Message,
    (request: ConnectRequest) => {
      return request.serializeBinary();
    },
    common_message_pb.Message.deserializeBinary
  );

  connect(
    request: ConnectRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_message_pb.Message) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/database.DatabaseService/Connect',
      request,
      metadata || {},
      this.methodInfoConnect,
      callback);
  }

  methodInfoInsertHouse = new grpcWeb.AbstractClientBase.MethodInfo(
    common_message_pb.Message,
    (request: InsertHouseRequest) => {
      return request.serializeBinary();
    },
    common_message_pb.Message.deserializeBinary
  );

  insertHouse(
    request: InsertHouseRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_message_pb.Message) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/database.DatabaseService/InsertHouse',
      request,
      metadata || {},
      this.methodInfoInsertHouse,
      callback);
  }

}

