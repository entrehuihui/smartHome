#!/bin/bash
protoc -I ./service/myrpc/proto --go_out ./service/myrpc/proto --go_opt paths=source_relative --go-grpc_out ./service/myrpc/proto --go-grpc_opt paths=source_relative service/myrpc/proto/*.proto

protoc -I ./service/myrpc/proto --grpc-gateway_out ./service/myrpc/proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./service/myrpc/proto/*.proto

protoc -I ./service/myrpc/proto --proto_path=../ --swagger_out=logtostderr=true:./service/myrpc/proto/ ./service/myrpc/proto/*.proto

# go install github.com/entrehuihui/aias
aias -d service/myrpc/proto
# go install github.com/entrehuihui/grpa-gateway-complete
grpa-gateway-complete
