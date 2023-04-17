FROM golang:1.20-alpine as builder

LABEL maintainer="Tran Nguyen Thuong Truong <"

Run mkdir -p /app
WORKDIR /app
COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls
RUN go clean --modcache
RUN apk update && apk add make && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o main ./cmd/main.go

FROM golang:1.20-alpine as production

RUN mkdir -p /app
WORKDIR /app

COPY --chown=0:0 --from=builder /app/ ./

ENTRYPOINT ["/app/main"]

# CMD ["air", "-c", ".air.toml"]
