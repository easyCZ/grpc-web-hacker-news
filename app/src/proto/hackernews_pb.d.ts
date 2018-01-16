// package: grpc_web_hacker_news
// file: proto/hackernews.proto

import * as jspb from "google-protobuf";

export class ItemId extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemId.AsObject;
  static toObject(includeInstance: boolean, msg: ItemId): ItemId.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemId, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemId;
  static deserializeBinaryFromReader(message: ItemId, reader: jspb.BinaryReader): ItemId;
}

export namespace ItemId {
  export type AsObject = {
    id: number,
  }
}

export class Item extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): ItemId | undefined;
  setId(value?: ItemId): void;

  getScore(): number;
  setScore(value: number): void;

  getTitle(): string;
  setTitle(value: string): void;

  getBy(): string;
  setBy(value: string): void;

  getTime(): number;
  setTime(value: number): void;

  getUrl(): string;
  setUrl(value: string): void;

  getType(): ItemType;
  setType(value: ItemType): void;

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
    id?: ItemId.AsObject,
    score: number,
    title: string,
    by: string,
    time: number,
    url: string,
    type: ItemType,
  }
}

export class ListStoriesResponse extends jspb.Message {
  hasStory(): boolean;
  clearStory(): void;
  getStory(): Item | undefined;
  setStory(value?: Item): void;

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
    story?: Item.AsObject,
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

export class GetStoryRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): ItemId | undefined;
  setId(value?: ItemId): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetStoryRequest): GetStoryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetStoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStoryRequest;
  static deserializeBinaryFromReader(message: GetStoryRequest, reader: jspb.BinaryReader): GetStoryRequest;
}

export namespace GetStoryRequest {
  export type AsObject = {
    id?: ItemId.AsObject,
  }
}

export class GetStoryResponse extends jspb.Message {
  hasStory(): boolean;
  clearStory(): void;
  getStory(): Item | undefined;
  setStory(value?: Item): void;

  getHtml(): Uint8Array | string;
  getHtml_asU8(): Uint8Array;
  getHtml_asB64(): string;
  setHtml(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStoryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetStoryResponse): GetStoryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetStoryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStoryResponse;
  static deserializeBinaryFromReader(message: GetStoryResponse, reader: jspb.BinaryReader): GetStoryResponse;
}

export namespace GetStoryResponse {
  export type AsObject = {
    story?: Item.AsObject,
    html: Uint8Array | string,
  }
}

export enum ItemType {
  UNKNOWN = 0,
  JOB = 1,
  STORY = 2,
  COMMENT = 3,
  POLL = 4,
  POLLOPT = 5,
}

