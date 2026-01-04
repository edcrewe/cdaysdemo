#!/bin/bash

# Init go module
go mod tidy
go run cmd/server/main.go
