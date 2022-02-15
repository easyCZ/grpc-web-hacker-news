"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.selectStory = exports.listStories = exports.listStoriesInit = exports.addStory = exports.SELECT_STORY = exports.ADD_STORY = exports.STORIES_INIT = void 0;
var hackernews_pb_1 = require("../proto/hackernews_pb");
var grpc_1 = require("../middleware/grpc");
var hackernews_pb_service_1 = require("../proto/hackernews_pb_service");
exports.STORIES_INIT = 'STORIES_INIT';
exports.ADD_STORY = 'ADD_STORY';
exports.SELECT_STORY = 'SELECT_STORY';
var addStory = function (story) { return ({ type: exports.ADD_STORY, payload: story }); };
exports.addStory = addStory;
var listStoriesInit = function () { return ({ type: exports.STORIES_INIT }); };
exports.listStoriesInit = listStoriesInit;
var listStories = function () {
    return (0, grpc_1.grpcRequest)({
        request: new hackernews_pb_1.ListStoriesRequest(),
        onStart: function () { return (0, exports.listStoriesInit)(); },
        onEnd: function (code, message, trailers) {
            console.log(code, message, trailers);
            return;
        },
        host: 'http://localhost:8900',
        methodDescriptor: hackernews_pb_service_1.HackerNewsService.ListStories,
        onMessage: function (message) {
            var story = message.getStory();
            if (story) {
                return (0, exports.addStory)(story);
            }
            return;
        },
    });
};
exports.listStories = listStories;
var selectStory = function (storyId) { return ({ type: exports.SELECT_STORY, payload: storyId }); };
exports.selectStory = selectStory;
