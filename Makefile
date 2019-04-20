
proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:./ ./protobuf/api.proto

