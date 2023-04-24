FROM golang:1.20-alpine as builder

LABEL maintainer="Tran Nguyen Thuong Truong <thuongtruongofficial@gmail.com>"

RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go get -v golang.org/x/tools/gopls
RUN go clean --modcache
RUN apk update && apk add make && apk add --no-cache git && apk add build-base

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o main cmd/main.go

FROM golang:1.20-alpine as production

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app ./

CMD ["./main"]
