import * as jspb from "google-protobuf"

import * as common_message_pb from '../common/message_pb';
import * as crawler_lianjia_pb from '../crawler/lianjia_pb';

export class ConnectRequest extends jspb.Message {
  getDatabase(): string;
  setDatabase(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectRequest): ConnectRequest.AsObject;
  static serializeBinaryToWriter(message: ConnectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectRequest;
  static deserializeBinaryFromReader(message: ConnectRequest, reader: jspb.BinaryReader): ConnectRequest;
}

export namespace ConnectRequest {
  export type AsObject = {
    database: string,
  }
}

export class InsertHouseRequest extends jspb.Message {
  getDatabase(): string;
  setDatabase(value: string): void;

  getTable(): string;
  setTable(value: string): void;

  getHouse(): crawler_lianjia_pb.House | undefined;
  setHouse(value?: crawler_lianjia_pb.House): void;
  hasHouse(): boolean;
  clearHouse(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InsertHouseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InsertHouseRequest): InsertHouseRequest.AsObject;
  static serializeBinaryToWriter(message: InsertHouseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InsertHouseRequest;
  static deserializeBinaryFromReader(message: InsertHouseRequest, reader: jspb.BinaryReader): InsertHouseRequest;
}

export namespace InsertHouseRequest {
  export type AsObject = {
    database: string,
    table: string,
    house?: crawler_lianjia_pb.House.AsObject,
  }
}

