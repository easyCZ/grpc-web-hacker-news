#!/usr/bin/env bash
protoc \
    --go_out=plugins=grpc:./server \
    --plugin=protoc-gen-ts=./app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./app \
    --js_out=import_style=commonjs,binary:./app \
    ./proto/hackernews.proto
