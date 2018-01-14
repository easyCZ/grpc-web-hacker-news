// package: grpc_web_hacker_news
// file: proto/hackernews.proto

import * as proto_hackernews_pb from "../proto/hackernews_pb";
export class HackerNewsService {
  static serviceName = "grpc_web_hacker_news.HackerNewsService";
}
export namespace HackerNewsService {
  export class ListStories {
    static readonly methodName = "ListStories";
    static readonly service = HackerNewsService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = proto_hackernews_pb.ListStoriesRequest;
    static readonly responseType = proto_hackernews_pb.ListStoriesResponse;
  }
}
