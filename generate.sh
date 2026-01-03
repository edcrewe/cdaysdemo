#!/bin/bash

# Create directory for Google APIs
mkdir -p googleapis/google/api
mkdir -p generated/go

# Download the required proto files
curl -o googleapis/google/api/annotations.proto \
https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
curl -o googleapis/google/api/http.proto \
https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
curl -o googleapis/google/api/httpbody.proto \
https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto

export GOOGLEAPIS_DIR=googleapis

# Generate gRPC service code
protoc -I=${GOOGLEAPIS_DIR} --proto_path=proto --go_out=./generated/go  --go_opt paths=source_relative --go-grpc_out=./generated/go --go-grpc_opt paths=source_relative csv.proto