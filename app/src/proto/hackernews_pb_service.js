"use strict";
// package: grpc_web_hacker_news
// file: proto/hackernews.proto
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.HackerNewsService = void 0;
var proto_hackernews_pb = __importStar(require("../proto/hackernews_pb"));
var HackerNewsService = /** @class */ (function () {
    function HackerNewsService() {
    }
    HackerNewsService.serviceName = "grpc_web_hacker_news.HackerNewsService";
    return HackerNewsService;
}());
exports.HackerNewsService = HackerNewsService;
(function (HackerNewsService) {
    var ListStories = /** @class */ (function () {
        function ListStories() {
        }
        ListStories.methodName = "ListStories";
        ListStories.service = HackerNewsService;
        ListStories.requestStream = false;
        ListStories.responseStream = true;
        ListStories.requestType = proto_hackernews_pb.ListStoriesRequest;
        ListStories.responseType = proto_hackernews_pb.ListStoriesResponse;
        return ListStories;
    }());
    HackerNewsService.ListStories = ListStories;
})(HackerNewsService = exports.HackerNewsService || (exports.HackerNewsService = {}));
exports.HackerNewsService = HackerNewsService;
