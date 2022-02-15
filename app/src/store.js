"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var redux_1 = require("redux");
var stories_1 = __importDefault(require("./reducers/stories"));
var grpc_1 = require("./middleware/grpc");
var reducers = (0, redux_1.combineReducers)({
    stories: stories_1.default,
});
exports.default = (0, redux_1.createStore)(reducers, (0, redux_1.applyMiddleware)((0, grpc_1.newGrpcMiddleware)()));
