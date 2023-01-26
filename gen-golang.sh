#!/usr/bin/env bash

protoDir="./hellostream"

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ${protoDir}/*proto