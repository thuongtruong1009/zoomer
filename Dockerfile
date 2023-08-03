FROM golang:1.20-alpine AS development
RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum docs ./
RUN go mod download
RUN go clean --modcache
RUN apk update && apk add make && apk add --no-cache git
COPY . .
RUN make setup
RUN go build -v -o main-dev ./cmd/main.go

FROM golang:1.20-alpine AS production
RUN apk update && apk --no-cache add ca-certificates
RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -tags migrate -o main-prod ./cmd/main.go

FROM scratch
WORKDIR /app
RUN addgroup -S zoomer
RUN adduser -S -D -h /app zoomer zoomer
RUN chown -R zoomer:zoomer /app
USER zoomer
COPY --chown=zoomer:zoomer --from=development /app /app/app-dev
COPY --chown=zoomer:zoomer --from=production /app/main-prod /app
EXPOSE 8080
CMD if [ "$TARGET" = "development" ]; \
    then /app/app-dev/main-dev; \
    else /app/main-prod; \
    fi

LABEL maintainer="Tran Nguyen Thuong Truong <thuongtruongofficial@gmail.com>"
LABEL org.opencontainers.image.authors="thuongtruong1009"
LABEL org.opencontainers.image.version="1.0"
LABEL org.opencontainers.image.description="Official Image of Zoomer application"
LABEL org.opencontainers.image.licenses="Apache-2.0"
LABEL org.opencontainers.image.source="https://github.com/thuongtruong1009/zoomer"
LABEL org.opencontainers.image.documentation="https://github.com/thuongtruong1009/zoomer/blob/main/README.md"
