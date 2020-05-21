#!/bin/bash

echo "Portfolio protobuf generation"
protoc -I ./services/portfolio/protobuf --go_out=plugins=grpc:./services/portfolio/protobuf \
  ./services/portfolio/protobuf/portfolio.proto
echo "New portfolio.pb.go was generated"