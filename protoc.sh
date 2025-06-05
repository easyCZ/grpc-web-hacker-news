#!/usr/bin/env bash
protoc \
    --go-grpc_out=./server \
    --go_out=./server \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --plugin=protoc-gen-js=./node_modules/.bin/protoc-gen-js \
    --ts_out=service=true:./app/src \
    --js_out=import_style=commonjs,binary:./app/src \
    ./proto/hackernews.proto