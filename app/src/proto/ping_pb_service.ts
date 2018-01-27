// package: grpc_web_hacker_news
// file: proto/ping.proto

import * as proto_ping_pb from "../proto/ping_pb";
export class PingService {
  static serviceName = "grpc_web_hacker_news.PingService";
}
export namespace PingService {
  export class Ping {
    static readonly methodName = "Ping";
    static readonly service = PingService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = proto_ping_pb.PingRequest;
    static readonly responseType = proto_ping_pb.PingResponse;
  }
}
