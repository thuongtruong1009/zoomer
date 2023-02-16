FROM golang:1.17-alpine as builder

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

RUN apk update
RUN apk add make

FROM golang:1.17-alpine as production

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

ENTRYPOINT ["./main"]

CMD ["air", "-c", ".air.toml"]