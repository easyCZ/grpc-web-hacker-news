## grpc-web-hacker-news
An example app implementing a Hacker News reader. This example aims to demonstrate usage of grpc-web with React. It additionally shows how to integrate with Redux.

### Running
To start both the Go backend server and the frontend React application, run the following:
```bash
./start.sh
```

The backend server is running on `http://localhost:8900` while the frontend will by default start on `http://localhost:3000`

## Notable setup points

### Disable TSLint for protobuf generated classes
Gernerated proto classes do not confirm to TS lint requirements, [disable linting](https://github.com/easyCZ/grpc-web-hacker-news/blob/master/app/tslint.json#L4)
```json
{
    "linterOptions": {
        "exclude": [
            "src/proto/*"
        ]
    },
}
```

### Configure protobuf compiler script
In this example, we're using a `protoc.sh` script to aid compilation
```bash
protoc \
    --go_out=plugins=grpc:./server \
    --plugin=protoc-gen-ts=./app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./app/src \
    --js_out=import_style=commonjs,binary:./app/src \
    ./proto/hackernews.proto
```

### gRPC-Web Redux Middleware
An example redux middleware is included.
```js
import { Action, Dispatch, Middleware, MiddlewareAPI } from 'redux';
import { Code, grpc, Metadata, Transport } from 'grpc-web-client';
import * as jspb from 'google-protobuf';

const GRPC_WEB_REQUEST = 'GRPC_WEB_REQUEST';
// const GRPC_WEB_INVOKE = 'GRPC_WEB_INVOKE';

// Descriptor of a grpc-web payload
// life-cycle methods mirror grpc-web but allow for an action to be dispatched when triggered
export type GrpcActionPayload<RequestType extends jspb.Message, ResponseType extends jspb.Message> = {
  // The method descriptor to use for a gRPC request, equivalent to grpc.invoke(methodDescriptor, ...)
  methodDescriptor: grpc.MethodDefinition<RequestType, ResponseType>,
  // The transport to use for grpc-web, automatically selected if empty
  transport?: Transport,
  // toggle debug messages
  debug?: boolean,
  // the URL of a host this request should go to
  host: string,
  // An instance of of the request message
  request: RequestType,
  // Additional metadata to attach to the request, the same as grpc-web
  metadata?: Metadata.ConstructorArg,
  // Called immediately before the request is started, useful for toggling a loading status
  onStart?: () => Action | void,
  // Called when response headers are received
  onHeaders?: (headers: Metadata) => Action | void,
  // Called on each incoming message
  onMessage?: (res: ResponseType) => Action | void,
  // Called at the end of a request, make sure to check the exit code
  onEnd: (code: Code, message: string, trailers: Metadata) => Action | void,
};

// Basic type for a gRPC Action
export type GrpcAction<RequestType extends jspb.Message, ResponseType extends jspb.Message> = {
  type: typeof GRPC_WEB_REQUEST,
  payload: GrpcActionPayload<RequestType, ResponseType>,
};

// Action creator, Use it to create a new grpc action
export function grpcRequest<RequestType extends jspb.Message, ResponseType extends jspb.Message>(
  payload: GrpcActionPayload<RequestType, ResponseType>
): GrpcAction<RequestType, ResponseType> {
  return {
    type: GRPC_WEB_REQUEST,
    payload,
  };
}

export function newGrpcMiddleware(): Middleware {
  return ({getState, dispatch}: MiddlewareAPI<{}>) => (next: Dispatch<{}>) => (action: any) => {
    // skip non-grpc actions
    if (!isGrpcWebUnaryAction(action)) {
      return next(action);
    }

    const payload = action.payload;

    if (payload.onStart) {
      payload.onStart();
    }

    grpc.invoke(payload.methodDescriptor, {
      debug: payload.debug,
      host: payload.host,
      request: payload.request,
      metadata: payload.metadata,
      transport: payload.transport,
      onHeaders: headers => {
        if (!payload.onHeaders) { return; }
        const actionToDispatch = payload.onHeaders(headers);
        return actionToDispatch && dispatch(actionToDispatch);
      },
      onMessage: res => {
        if (!payload.onMessage) { return; }
        const actionToDispatch = payload.onMessage(res);
        return actionToDispatch && dispatch(actionToDispatch);
      },
      onEnd: (code, msg, trailers) => {
        const actionToDispatch = payload.onEnd(code, msg, trailers);
        return actionToDispatch && dispatch(actionToDispatch);
      },
    });

    return next(action);
  };
}

function isGrpcWebUnaryAction(action: any): action is GrpcAction<jspb.Message, jspb.Message> {
  return action && action.type && action.type === GRPC_WEB_REQUEST && isGrpcWebPayload(action);
}

function isGrpcWebPayload(action: any): boolean {
  return action &&
    action.payload &&
    action.payload.methodDescriptor &&
    action.payload.request &&
    action.payload.onEnd &&
    action.payload.host;
}


```
