#!/bin/bash

echo "Portfolio protobuf generation"
protoc -I/usr/local/include -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
  ./services/portfolio/protobuf/portfolio.proto \
  --go_out=plugins=grpc:./services/portfolio/protobuf
echo "New portfolio.pb.go was generated"


