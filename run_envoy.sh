#!/bin/bash

# Start the Envoy proxy
envoy --service-node ingress --service-cluster ingress -c envoy-config.yaml --log-level debug

# wget http://localhost:8888/v1/plain/small.csv