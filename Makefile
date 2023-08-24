include .env
export $(shell type .env | findstr /v /c:"#" /c:"=")

DOCKER_USERNAME ?= thuongtruong1009
APPLICATION_NAME ?= ${APP_NAME}
GIT_HASH ?= $(shell git log --format="%%h" -n 1)
ENTRYPOINT ?= cmd/api/main.go
BUILDPOINT ?= release/latest
MIGRATION_ENTRYPOINT ?= db/migrations

_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile

setup:
	@echo "Installing dependencies..."
	go mod tidy
	go get -v -t -d ./...
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "Done!"

dev:
	air -c .air.toml

run:
	go run ${ENTRYPOINT}

start:
	go run ${BUILDPOINT}

tests:
	@echo "Running tests..."
	go clean -testcache && go vet -v ./... && govulncheck ./... \
  go test ./pkg/validators/... ./pkg/utils/... ./infrastructure/cache/... ./pkg/interceptor/... ./pkg/helpers/... ./pkg/shared/... -v -coverprofile=logs/coverage.txt -covermode=atomic \
  go test -timeout 30s ./pkg/helpers -run ^TestParallelize$ -v
	@echo "Done!"

lint:
	@echo "Running linter..."
	gofmt -w . && goimports -w . && go fmt ./...
	golangci-lint version
	golangci-lint run -c .golangci.yml ./...
	@echo "Done!"

build:
	@ echo "Building ${APPLICATION_NAME}..."
	go build -tags migrate -v -trimpath -o ${BUILDPOINT} ${ENTRYPOINT}
	@ echo "Done!"

docs:
	@ echo "Generating docs..."
	swag fmt && swag init -g ${ENTRYPOINT} -o ./docs --generatedTime=true
	@ echo "Done!"

seed:
	@ echo "Seeding data..."
	go run ./cmd/seed/main.go
	@ echo "Done!"

# Migration

migration-create:
	set /p Name="Please provide name for the migration: " && migrate create -ext sql -dir ${MIGRATION_ENTRYPOINT} -seq %Name%

migration-up:
	@ echo "Migrating up..."
	@ set -p N="How many migration you wants to perform (default value: [all]): " && migrate -path ${MIGRATION_ENTRYPOINT} -verbose -database "${PG_MIGRATE_URI}" up %N:~-1%
	@ echo "Done!"

migration-down:
	@ echo "Migrating down..."
	migrate -path ${MIGRATION_ENTRYPOINT} -verbose -database "${PG_MIGRATE_URI}" down
	@ echo "Done!"

migrate-status:
	migrate -path ${MIGRATION_ENTRYPOINT} -verbose -database "${PG_MIGRATE_URI}" status=

# Docker

docker-build:
	@ echo "Building ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}..."
	docker build --pull --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .
	@ echo "Done!"

docker-dev:
	@ echo "Building ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}..."
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build -d
	@ echo "Running..."

docker-prod:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

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
