FROM golang:1.20-alpine as builder

LABEL maintainer="Tran Nguyen Thuong Truong <"

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go clean --modcache

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o main ./cmd/main.go

RUN apk update && apk add make && apk add --no-cache git && apk add --no-cache bash && apk add build-base

FROM golang:1.20-alpine as production

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls

# ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main
# ENTRYPOINT ["./main"]
CMD ["/app/main"]

# CMD ["air", "-c", ".air.toml"]

