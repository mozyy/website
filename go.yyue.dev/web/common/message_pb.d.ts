import * as jspb from "google-protobuf"

export class Message extends jspb.Message {
  getState(): Status;
  setState(value: Status): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    state: Status,
    message: string,
  }
}

export enum Status { 
  SUCCESS = 0,
  INFO = 1,
  WARNING = 2,
  ERROR = 3,
}
