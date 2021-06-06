#!/usr/bin/bash

# THIS SCRIPT RELIES ON THE FOLLOWING
# - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
# - go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

# Generates proto-contract source-code
protoc --go-grpc_out=. userservice.proto \
  --go_out=. --go_opt=paths=source_relative
mv userservice.pb.go usergrpc/userservice
