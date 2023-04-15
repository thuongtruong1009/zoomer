include .env.example
export $(shell sed 's/=.*//' .env.example)

DOCKER_USERNAME ?= thuongtruong1009
APPLICATION_NAME ?= zoomer
GIT_HASH ?= $(shell git log --format="%h" -n 1)
ENTRYPOINT ?= cmd/main.go

_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile

dev:
	gofmt -w . && go run ${ENTRYPOINT}

test:
	go test -v -race -coverprofile=coverage -covermode=atomic -short ./...

build:
	go build -o ${APPLICATION_NAME} ${ENTRYPOINT}

docker_build:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .

docker_push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}

docker_release:
	docker pull ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}
	docker tag  ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_RELEASE_TAG}

.PHONY: dev_run dev_test dev_build docker_build docker_push docker_release
