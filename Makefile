
proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:./ --java_out=protobuf/  ./protobuf/api.proto

