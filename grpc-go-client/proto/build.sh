#!/usr/bin/env bash
protoc --go_out=plugins=grpc:. --micro_out=. greeter.proto
