#!/bin/bash

# *.pb.gw.go
protoc -I . \
   -I$GOPATH/src \
   -I$GOPATH/src/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	user.proto

# *_grpc.pb.go
protoc -I. -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/grpc-gateway/third_party/googleapis \
  --go-grpc_out=. --go-grpc_opt paths=source_relative \
  user.proto

# *.pb.go
protoc -I=. \
   -I$GOPATH/src \
   -I$GOPATH/src/grpc-gateway/third_party/googleapis \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative \
	user.proto