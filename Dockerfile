FROM golang:1.20-alpine as development

LABEL maintainer="Tran Nguyen Thuong Truong <thuongtruongofficial@gmail.com>"

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

# CMD ["/app/app-dev"]

# CMD ["air", "-c", ".air.toml"]
CMD if [ "$TARGET" = "development" ]; then \
        ./app-dev; \
    else \
        ./app-prod; \
    fi
