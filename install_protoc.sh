#!/bin/bash

# Once Golang is installed eg. 1.24.11, then install envoy see README
# Now install protoc...

echo "See https://protobuf.dev/installation/ for your enviorment."

echo "This script installs the protoc generators in the default Go assuming protobuf is already installed."
# Plugin for basic Go types
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Plugin for gRPC service code
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

