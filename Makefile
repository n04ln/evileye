databasefile := data.sqlite3
schemafile := schema/schema.sql
env := GO111MODULE=on CGO_ENABLED=1
for_linux := GOOS=linux GOARCH=amd64
image := golang:1.12

proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:./ ./protobuf/api.proto

build:
	${for_linux} ${env} go build -ldflags="-s -w -X .github.com/NoahOrberg/evileye/controller.commitHash=`git log --pretty=format:%H -n 1` -X .github.com/NoahOrberg/evileye/controller.buildTime=`date +%s`" -o bin/evileye

# Why using container image for build binary?
# A. cannot use cgo in MacOS, So build at Linux.
build-for-image:
	docker run --rm \
		-v ${PWD}:/go/src/github.com/NoahOrberg/evileye \ # for artifact is saved in local machine
		-v ${GOPATH}/pkg/mod/:/go/pkg/mod/ \ # for shaing cache dir
		${image} /bin/sh -c "cd /go/src/github.com/NoahOrberg/evileye; make build"

run:
	${env} go build -ldflags="-s -w -X github.com/NoahOrberg/evileye/controller.commitHash=`git log --pretty=format:%H -n 1` -X github.com/NoahOrberg/evileye/controller.buildTime=`date +%s`" -o bin/evileye_for_run
	./bin/evileye_for_run

docker: build-for-image
	docker build -t noahorberg/evileye:latest ./

docker-compose: docker
	docker-compose down
	docker-compose up -d
	sleep 2
	docker cp schema/ evileye:/
	docker-compose exec evileye /bin/sh /schema/provisioning.sh /${schemafile} /${databasefile}

reset-db:
	-[ -e ${databasefile} ] && rm -rf ${databasefile}
	touch ${databasefile}
	cat ${schemafile} | sqlite3 ${databasefile} 
