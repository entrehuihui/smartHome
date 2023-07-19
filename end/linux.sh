#!/bin/bash
echo "install protoc-gen-go/protoc-gen-go-grpc/protoc-gen-grpc-gateway/protoc-gen-openapiv2/protoc-gen-swagger"
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
protoc -I ./service/myrpc/proto --go_out ./service/myrpc/proto --go_opt paths=source_relative --go-grpc_out ./service/myrpc/proto --go-grpc_opt paths=source_relative service/myrpc/proto/*.proto
protoc -I ./service/myrpc/proto --grpc-gateway_out ./service/myrpc/proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./service/myrpc/proto/*.proto
protoc -I ./service/myrpc/proto --proto_path=../ --swagger_out=logtostderr=true:./service/myrpc/proto/ ./service/myrpc/proto/*.proto

echo "install aias"
go get github.com/entrehuihui/aias
go install github.com/entrehuihui/aias
aias -d service/myrpc/proto

echo "install gateway"
go get github.com/entrehuihui/grpa-gateway-complete
go install github.com/entrehuihui/grpa-gateway-complete
grpa-gateway-complete
