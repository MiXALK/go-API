#!/bin/bash

echo "Generation protobuf go file, gateway file and swagger json for portfolio service"
protoc -I/usr/local/include -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
  ./services/portfolio/protobuf/portfolio.proto \
  --grpc-gateway_out=logtostderr=true:services/api-client \
  --swagger_out=logtostderr=true:swagger-ui/ \
  --go_out=plugins=grpc:services/portfolio/protobuf

echo "Generation protobuf portfolio go file for api client service"
protoc -I/usr/local/include -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
  ./services/portfolio/protobuf/portfolio.proto \
  --go_out=plugins=grpc:services/api-client/services/portfolio/protobuf



