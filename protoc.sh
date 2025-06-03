#!/usr/bin/env bash
protoc \
    --go-grpc_out=./server \
    --go_out=./server \
    --plugin=protoc-gen-ts=./app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./app/src \
    --js_out=import_style=commonjs,binary:./app/src \
    ./proto/hackernews.proto