protoc -I ./service/myrpc/proto --go_out ./service/myrpc/proto --go_opt paths=source_relative --go-grpc_out ./service/myrpc/proto --go-grpc_opt paths=source_relative service/myrpc/proto/*.proto

protoc -I ./service/myrpc/proto --grpc-gateway_out ./service/myrpc/proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./service/myrpc/proto/*.proto

protoc -I ./service/myrpc/proto --proto_path=../ --swagger_out=logtostderr=true:./service/myrpc/proto/ ./service/myrpc/proto/*.proto

aias -d service/myrpc/proto
complete

@REM echo Set up cross-compilation environment
@REM go env -w GOSUMDB=off
@REM set GOOS=linux
@REM set GOARCH=amd64

@REM go build -o smartHome main.go

@REM echo Clear cross-compilation environment 
@REM set GOARCH=
@REM set GOOS=
@REM go env -w GOSUMDB=on