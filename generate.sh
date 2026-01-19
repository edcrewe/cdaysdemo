#!/bin/bash

# Create directory for Google API
rm -fr generated
mkdir -p googleapis/google/api
mkdir -p generated/go/v1

# Download the required proto files
curl -o googleapis/google/api/annotations.proto \
https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
curl -o googleapis/google/api/http.proto \
https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
curl -o googleapis/google/api/httpbody.proto \
https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto
curl -L -o protovalidate.tar.gz https://github.com/bufbuild/protovalidate/archive/refs/tags/v1.0.0.tar.gz
tar -xzf protovalidate.tar.gz --strip-components=2 "protovalidate-1.0.0/proto"
rm protovalidate.tar.gz
rm -fr protovalidate-testing

# Generate gRPC service code
protoc -I=googleapis -I=protovalidate --proto_path=proto --go_out=./generated/go/v1  --go_opt paths=source_relative --go-grpc_out=./generated/go/v1 --go-grpc_opt paths=source_relative demo.proto

# Generate proto descriptor file for Envoy
protoc -I=googleapis -I=protovalidate --proto_path=proto --include_imports      --descriptor_set_out=proto.pb demo.proto

go mod vendor;go mod tidy