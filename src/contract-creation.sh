#!/usr/bin/bash

# Generates proto-contract source-code
protoc --go-grpc_out=. userservice.proto \
  --go_out=. --go_opt=paths=source_relative
mv userservice.pb.go usergrpc/userservice
