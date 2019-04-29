database_file := data.sqlite3
schema_file := schema/schema.sql
env := GO111MODULE=on CGO_ENABLED=1
for_linux := GOOS=linux GOARCH=amd64
image := golang:1.12
tag := $(shell git symbolic-ref --short HEAD | sed 's,/,-,g')
docker_compose_env := EVIL_EYE_IMAGE="noahorberg/evileye:${tag}" DATABASEFILE="/evileye/${database_file}"
commit_hash := $(shell git log --pretty=format:%H -n 1)
build_time := $(shell date +%s)
build_container := "build_evileye"

proto:
	protoc --proto_path=${GOPATH}/src:./protobuf/ --go_out=plugins=grpc:./protobuf ./protobuf/*.proto

buildi: build-for-image
	docker build -t noahorberg/evileye:${tag} ./

push: buildi
	docker push noahorberg/evileye:${tag}

build:
	${for_linux} ${env} go build -ldflags="-s -w -X main.commitHash=${commit_hash} -X main.buildTime=${build_time}" -o bin/evileye

# Q. Why using container image for building binary?
# A. Cannot build a binary for Linux with cgo in MacOS, So build in Linux.
build-for-image:
	docker run --name ${build_container} \
		-v ${PWD}:/go/src/github.com/NoahOrberg/evileye \
		-v ${HOME}/Library/Caches/go-build:/tmp/go-build \
		-v ${GOPATH}/pkg/mod/:/go/pkg/mod/ \
		${image} /bin/sh -c "cd /go/src/github.com/NoahOrberg/evileye; make build"
	docker commit ${build_container} ${image}
	docker rm ${build_container}

run-local:
	${env} go build -ldflags="-s -w -X main.commitHash=`git log --pretty=format:%H -n 1` -X main.buildTime=`date +%s`" -o bin/evileye_for_run
	./bin/evileye_for_run

# docker: build-for-image
	# docker build -t noahorberg/evileye:${tag} ./
docker:
	docker pull noahorberg/evileye:${tag}

docker-compose: docker docker-compose-without-pull

docker-compose-without-pull:
	${docker_compose_env} docker-compose down
	${docker_compose_env} docker-compose up -d
	sleep 2
	docker cp schema/ evileye1:/
	${docker_compose_env} docker-compose exec evileye1 /bin/sh /schema/provisioning.sh /${schema_file} /evileye/${database_file}
	docker cp schema/ evileye2:/
	${docker_compose_env} docker-compose exec evileye2 /bin/sh /schema/provisioning.sh /${schema_file} /evileye/${database_file}
	docker cp schema/ evileye3:/
	${docker_compose_env} docker-compose exec evileye3 /bin/sh /schema/provisioning.sh /${schema_file} /evileye/${database_file}

reset-db:
	-[ -e ${database_file} ] && rm -rf ${database_file}
	touch ${database_file}
	cat ${schema_file} | sqlite3 ${database_file} 

docker-compose-down:
	${docker_compose_env} docker-compose down
