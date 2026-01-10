#!/bin/bash

# Start the gRPC server and Envoy proxy
go run cmd/server/main.go &
sleep 3
envoy --service-node ingress --service-cluster ingress -c envoy-config.yaml --log-level debug

# wget http://localhost:8888/v1/plain/small.csv