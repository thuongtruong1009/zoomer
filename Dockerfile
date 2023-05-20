LABEL maintainer="Tran Nguyen Thuong Truong <thuongtruongofficial@gmail.com>"
LABEL org.opencontainers.image.authors="thuongtruong1009"
LABEL org.opencontainers.image.version="1.0"
LABEL org.opencontainers.image.description="Official Image of Zoomer application"
LABEL org.opencontainers.image.licenses="Apache-2.0"
LABEL org.opencontainers.image.source="https://github.com/thuongtruong1009/zoomer"
LABEL org.opencontainers.image.documentation="https://github.com/thuongtruong1009/zoomer/blob/main/README.md"

FROM golang:1.20-alpine as development

RUN apk update && apk add make git build-base bash

RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go clean --modcache

COPY . .

RUN go build -o app-dev ./cmd/main.go

FROM golang:1.20-alpine as production

RUN apk update && apk add ca-certificates

RUN mkdir -p /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o app-prod ./cmd/main.go

FROM golang:1.20-alpine

WORKDIR /app

COPY --from=development /app /app/app-dev
COPY --from=production /app /app/app-prod

CMD if [ "$TARGET" = "development" ]; then \
        ./app-dev; \
    else \
        ./app-prod; \
    fi
