// package: grpc_web_hacker_news
// file: proto/hackernews.proto

import * as jspb from "google-protobuf";

export class Story extends jspb.Message {
  getId(): number;
  setId(value: number): void;

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

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Story.AsObject;
  static toObject(includeInstance: boolean, msg: Story): Story.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Story, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Story;
  static deserializeBinaryFromReader(message: Story, reader: jspb.BinaryReader): Story;
}

export namespace Story {
  export type AsObject = {
    id: number,
    score: number,
    title: string,
    by: string,
    time: number,
    url: string,
  }
}

export class ListStoriesResponse extends jspb.Message {
  hasStory(): boolean;
  clearStory(): void;
  getStory(): Story | undefined;
  setStory(value?: Story): void;

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
    story?: Story.AsObject,
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

