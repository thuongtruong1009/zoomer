include .env
export $(shell sed 's/=.*//' .env)

DOCKER_USERNAME ?= thuongtruong1009
APPLICATION_NAME ?= ${APP_NAME}
GIT_HASH ?= $(shell git log --format="%h" -n 1)
ENTRYPOINT ?= cmd/main.go

_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile

setup:
  go get -u ./...
	go mod tidy
	go install github.com/cosmtrek/air@v1.27.3
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

dev:
	go run ${ENTRYPOINT}

air:
	gofmt -w . && air

test:
	go test -v -race -coverprofile=coverage -covermode=atomic -short ./...

build:
	go build -o ${APPLICATION_NAME} ${ENTRYPOINT}

docs:
	swag i --dir ./cmd/,\
	./modules/,\
	./pkg/wrapper,\
	./pkg/contexts

# Migration

migration-create:
	migrate create -ext sql -dir db/migrations/sql -seq $(name)

migration-up:
	migrate -path db/migrations/sql -verbose -database "${PG_MIGRATE_URI}" up

migration-down:
	migrate -path db/migrations/sql -verbose -database "${PG_MIGRATE_URI}" down

migrate-status:
	migrate -path db/migrations/sql -verbose -database "${PG_MIGRATE_URI}" status

seed:
	go run ./cmd/seed/main.go

# Docker

docker-build:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .

docker-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

docker-prod:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

docker-push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}

docker-release:
	docker pull ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}
	docker tag  ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_RELEASE_TAG}

git-hooks:
	echo "Installing git hooks..." && \
	rm -rf .git/hooks/pre-commit && \
	ln -s ../../scripts/pre-commit.sh .git/hooks/pre-commit && \
	chmod +x .git/hooks/pre-commit && \
	echo "Done!"

.PHONY: setup dev air test build docs seed migrate-create migrate-up migrate down migrate-down migrate-status docker-build docker-dev docker-prod docker-push docker-release git-hooks
