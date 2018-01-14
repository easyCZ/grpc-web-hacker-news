// package: grpc_web_hacker_news
// file: proto/hackernews.proto

import * as jspb from "google-protobuf";

export class Item extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Item.AsObject;
  static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Item, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Item;
  static deserializeBinaryFromReader(message: Item, reader: jspb.BinaryReader): Item;
}

export namespace Item {
  export type AsObject = {
    id: number,
  }
}

export class ListStoriesResponse extends jspb.Message {
  clearStoriesList(): void;
  getStoriesList(): Array<Item>;
  setStoriesList(value: Array<Item>): void;
  addStories(value?: Item, index?: number): Item;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListStoriesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListStoriesResponse): ListStoriesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListStoriesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListStoriesResponse;
  static deserializeBinaryFromReader(message: ListStoriesResponse, reader: jspb.BinaryReader): ListStoriesResponse;
}

export namespace ListStoriesResponse {
  export type AsObject = {
    storiesList: Array<Item.AsObject>,
  }
}

export class ListStoriesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListStoriesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListStoriesRequest): ListStoriesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListStoriesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListStoriesRequest;
  static deserializeBinaryFromReader(message: ListStoriesRequest, reader: jspb.BinaryReader): ListStoriesRequest;
}

export namespace ListStoriesRequest {
  export type AsObject = {
  }
}

