include .env
export $(shell sed 's/=.*//' .env)

DOCKER_USERNAME ?= thuongtruong1009
APPLICATION_NAME ?= zoomer
GIT_HASH ?= $(shell git log --format="%h" -n 1)
ENTRYPOINT ?= cmd/main.go

_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile

setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

dev:
	go run ${ENTRYPOINT}

air:
	air -c .air.toml -d

test:
	go test -v -race -coverprofile=coverage -covermode=atomic -short ./...

build:
	go build -o ${APPLICATION_NAME} ${ENTRYPOINT}

# Migration

migration-create:
	migrate create -ext sql -dir migrations/sql $(name)

migration-up:
	migrate -path migrations/sql -verbose -database "${DATABASE_URL}" up

migration-down:
	migrate -path migrations/sql -verbose -database "${DATABASE_URL}" down

setup:
	go get -u github.com/cosmtrek/air
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

docs:
	swag i --dir ./cmd/,\
	./modules/,\
	./pkg/wrapper,\
	./pkg/contexts

# Docker

docker_build:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .

docker_push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}

docker_release:
	docker pull ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}
	docker tag  ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_RELEASE_TAG}

git-hooks:
	echo "Installing git hooks..." && \
	rm -rf .git/hooks/pre-commit && \
	ln -s ../../scripts/pre-commit.sh .git/hooks/pre-commit && \
	chmod +x .git/hooks/pre-commit && \
	echo "Done!"

.PHONY: dev air test build docker_build docker_push docker_release git-hooks
