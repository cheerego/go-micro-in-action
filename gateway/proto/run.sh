#!/usr/bin/env bash

protoc  -I. -I /Users/hkn/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. --go_out=plugins=grpc:. --micro_out=.  greeter.proto