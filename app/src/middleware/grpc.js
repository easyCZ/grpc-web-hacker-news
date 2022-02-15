"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.newGrpcMiddleware = exports.grpcRequest = void 0;
var grpc_web_client_1 = require("grpc-web-client");
var GRPC_WEB_REQUEST = 'GRPC_WEB_REQUEST';
// Action creator, Use it to create a new grpc action
function grpcRequest(payload) {
    return {
        type: GRPC_WEB_REQUEST,
        payload: payload,
    };
}
exports.grpcRequest = grpcRequest;
/* tslint:disable:no-any*/
function newGrpcMiddleware() {
    return function (_a) {
        var getState = _a.getState, dispatch = _a.dispatch;
        return function (next) { return function (action) {
            // skip non-grpc actions
            if (!isGrpcWebUnaryAction(action)) {
                return next(action);
            }
            var payload = action.payload;
            if (payload.onStart) {
                payload.onStart();
            }
            grpc_web_client_1.grpc.invoke(payload.methodDescriptor, {
                debug: payload.debug,
                host: payload.host,
                request: payload.request,
                metadata: payload.metadata,
                transport: payload.transport,
                onHeaders: function (headers) {
                    if (!payload.onHeaders) {
                        return;
                    }
                    var actionToDispatch = payload.onHeaders(headers);
                    return actionToDispatch && dispatch(actionToDispatch);
                },
                onMessage: function (res) {
                    if (!payload.onMessage) {
                        return;
                    }
                    var actionToDispatch = payload.onMessage(res);
                    return actionToDispatch && dispatch(actionToDispatch);
                },
                onEnd: function (code, msg, trailers) {
                    var actionToDispatch = payload.onEnd(code, msg, trailers);
                    return actionToDispatch && dispatch(actionToDispatch);
                },
            });
            return next(action);
        }; };
    };
}
exports.newGrpcMiddleware = newGrpcMiddleware;
function isGrpcWebUnaryAction(action) {
    return action && action.type && action.type === GRPC_WEB_REQUEST && isGrpcWebPayload(action);
}
function isGrpcWebPayload(action) {
    return action &&
        action.payload &&
        action.payload.methodDescriptor &&
        action.payload.request &&
        action.payload.onEnd &&
        action.payload.host;
}
/* tslint:enable:no-any*/
