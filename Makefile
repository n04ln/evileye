proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:./ ./protobuf/api.proto

build:
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags="-s -w -X main.commitHash=`git log --pretty=format:%H -n 1` -X main.buildTime=`date +%s`" -o bin/evileye

run:
	GO111MODULE=on go build -ldflags="-s -w -X main.commitHash=`git log --pretty=format:%H -n 1` -X main.buildTime=`date +%s`" -o bin/evileye_for_run
	./bin/evileye_for_run

docker: build
	docker build -t noahorberg/evileye:latest ./

docker-compose: docker
	docker-compose down
	docker-compose up

