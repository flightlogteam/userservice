#!/usr/bin/bash

# Generates proto-contract source-code
protoc -I=./ --go_out=plugins=grpc:userservice userservice.proto
