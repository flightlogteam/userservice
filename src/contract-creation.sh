#!/usr/bin/bash

# Generates proto-contract source-code
protoc --go-grpc_out=. userservice.proto
