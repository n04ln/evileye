
proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=./ --java_out=protobuf/  ./protobuf/*.proto

