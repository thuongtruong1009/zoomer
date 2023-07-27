<h1 align="center">ZOOMER</h1>

## Preview

<div>
  <img src="./public/demo/signup.PNG" alt="signup" width="400" /><img src="./public/demo/signin.PNG" alt="signin" width="400" />
</div>

## **Technical stuff**

- Architecture: Clean architecture
- Framework: Echo
- ORM: Gorm
- DB: Postgres, Redis
- Deployment: Docker
- Hot-Reloader: Air
- Cache: Redis
- Message Queue: RabbitMQ
- Stream: WebRTC
- Swagger: Echo-Swagger
- Peer connection: HTTP, WebSocket
- UML diagram: Diagram.net

## **What news**

- [x] Support JWT
- [x] Limit rooms per user in a day
- [x] Users can signup and login
- [x] The only way the user can message have to know the receiver user name.
- [x] Users can access their chat history.
- [x] Users can block each other
- [x] Dockerize and scalability
- [x] Support Redis cache
- [x] Support Peer connection
- [ ] Support OAuth2
- [ ] Support Swagger
- [ ] Support notification
- [ ] Support file upload/download
- [ ] Add missing intergration tests and Unit Test Coverage
- [ ] Add role/permission based validation
- [ ] Implement new features

## Architecture

![Architecture](./public/architecture.png)

## Usecase

![Usecase](./public/usecase.png)

## **Realtime Chat**

- First, we have the hub running on a separate goroutine which is the central place that manages different channels and contains a map of rooms. The hub has a Register and an Unregister channel to register/unregister clients, and a Broadcast channel that receives a message and broadcasts it out to all the other clients in the same room.

A room is initially empty. Only when a client hits the `/chats/joinRoom` endpoint, that will create a new client object in the room and it will be registered through the hub's Register channel.

Each client has a `writeMessage` and a `readMessage` method. `readMessage` reads the message through the client's websocket connection and send the message to the Broadcast channel in the hub, which will then broadcast the message out to every client in the same room. The `writeMessage` method in each of those clients will write the message to its websocket connection, which will be handled on the frontend side to display the messages accordingly.

## **How to run the code locally**

##### 1. Clone this repository

##### 2. Update .env file

##### 3. Run the code

- With local

```bash
go run cmd/main.go
```

- With Docker

Update .env file (change host to postgresql)

```bash
docker-compose up -d
```

#### **Testing**

- Unit tests

```bash
go test -v -cover ./...
```

- Integration tests

```bash
go test -v ./integration-tests
```

<!-- **References**

- [Clean architecture 1](https://github.dev/manakuro/golang-clean-architecture)

- [Clean architecture 2](https://github.dev/evrone/go-clean-template)

- [Web socket chat](https://www.youtube.com/watch?v=W9SuX9c40s8)

- [Status user ref](https://anonystick.com/blog-developer/check-user-online-hay-offline-nhu-facebook-voi-1-dong-code-javascript-2020112018731223)

- [V1 ref](https://www.thepolyglotdeveloper.com/2016/12/create-real-time-chat-app-golang-angular-2-websockets/)

- [Chat ref](https://github.com/ong-gtp/go-chat)

- [Redis Cache](https://dev.to/aseemwangoo/using-redis-for-caching-2022-2og5)

- [Sample template](https://github.dev/dzungtran/echo-rest-api)

- [Streaming file with gRPC 1](http://www.inanzzz.com/index.php/post/152g/transferring-files-with-grpc-client-side-streams-using-golang)

- [Streaming file with gRPC 2](https://dev.to/techschoolguru/upload-file-in-chunks-with-client-streaming-grpc-golang-4loc)

- [MinIO uploader](https://dev.to/minhblues/easy-file-uploads-in-go-fiber-with-minio-393c)

- [Goapp template](https://github.com/bnkamalesh/goapp)

- [RabbitMQ docs](https://topdev.vn/blog/ket-noi-amqp-client-voi-rabbitmq-server/)

- [RabbitMQ example 1](https://github.com/lakhinsu/rabbitmq-go-example)

- [Gin template](https://github.com/lakhinsu/gin-boilerplate)

- [Elastic search](https://github.dev/NooBeeID/simple-go-search/tree/develop) -->
