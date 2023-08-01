include .env
export $(shell type .env | findstr /v /c:"#" /c:"=")

DOCKER_USERNAME ?= thuongtruong1009
APPLICATION_NAME ?= ${APP_NAME}
GIT_HASH ?= $(shell git log --format="%%h" -n 1)
ENTRYPOINT ?= cmd/main.go
BUILDPOINT ?= release/latest

_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile

setup:
	@echo "Installing dependencies..."
	go mod tidy
	go install github.com/cosmtrek/air@v1.27.3
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Done!"

dev:
	air -c .air.toml

run:
	go run ${ENTRYPOINT}

start:
	go run ${BUILDPOINT}

tests:
	@echo "Running tests..."
	cd scripts && run-tests.sh
	@echo "Done!"

lint:
	@echo "Running linter..."
	gofmt -w . && goimports -w . && go fmt ./...
	golangci-lint version
	golangci-lint run -c .golangci.yml ./...
	@echo "Done!"

build:
	@ echo "Building ${APPLICATION_NAME}..."
	@ go build -trimpath -o ${BUILDPOINT} ${ENTRYPOINT}
	@ echo "Done!"

docs:
	@ echo "Generating docs..."
	swag i --dir ./cmd/, ./internal/auth/delivery/, ./internal/rooms/delivery/, ./internal/stream/delivery/, ./internal/chats/delivery/, ./internal/resources/delivery/, ./internal/search/delivery
	swag init -g ./cmd/main.go --output ./docs
	@ echo "Done!"

# Migration

migration-create:
	set /p Name="Please provide name for the migration: " && migrate create -ext sql -dir db/migrations/sql -seq %Name%

migration-up:
	@ echo "Migrating up..."
	@ set -p N="How many migration you wants to perform (default value: [all]): " && migrate -path db/migrations/sql -verbose -database "${PG_MIGRATE_URI}" up %N:~-1%
	@ echo "Done!"

migration-down:
	@ echo "Migrating down..."
	migrate -path db/migrations/sql -verbose -database "${PG_MIGRATE_URI}" down
	@ echo "Done!"

migrate-status:
	migrate -path db/migrations/sql -verbose -database "${PG_MIGRATE_URI}" status=

seed:
	@ echo "Seeding data..."
	go run ./cmd/seed/main.go
	@ echo "Done!"

# Docker

docker-build:
	@ echo "Building ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}..."
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .
	@ echo "Done!"

docker-dev:
	@ echo "Building ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}..."
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
	@ echo "Running..."

docker-prod:
	@ echo "Building ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}..."
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
	@ echo "Done!"

docker-push:
	@ echo "Pushing ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} to docker hub..."
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}
	@ echo "Done!"

docker-release:
	docker pull ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}
	docker tag  ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_RELEASE_TAG}

git-hooks:
	@ echo "Installing git hooks..."
	rmdir /s /q .git\hooks
	mklink /H .git\hooks\pre-commit scripts\pre-commit.sh
	@ echo "Done!"


.PHONY: setup dev run start test build docs seed migrate-create migrate-up migrate down migrate-down migrate-status docker-build docker-dev docker-prod docker-push docker-release git-hooks
