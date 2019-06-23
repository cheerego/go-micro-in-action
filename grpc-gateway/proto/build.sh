#!/usr/bin/env bash

protoc  -I. -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. --go_out=plugins=grpc:.  greeter.proto