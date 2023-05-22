There is the most valuable changes log:

### v2.0

**1.Architecture**

- Connect to Redis database for chat service
- Config Cache adapter with Redis
- Config RabbitMQ for chat service

**2. News**

- Update CI for test, build, deploy and publish image
- Set and Read cache for userid, username, roomlist response
- Wakeup socket x2 faster with goroutine
- Wrap context function with custom interceptor
- Refactor code
- Generate swagger document and UI
- Config migration for database

**3. Features**

- Setup Video call service
- Connect done to client
- Add/Connect to new user
- Send real-time messages to user with timestamp record
- Save/Load messages to database
- Rejoin room but not remove old messages
- Load recent friend contact list
- Search and connect new contact

**4. Bugfixes**

- fix: reload old messages when rejoin room
- fix: reconnect to websocket when rejoin room
- fix: display video partner in chat
- fix: reconnect to websocket when rejoin room

### [v1.0 - 28-03-2023](https://github.com/thuongtruong1009/zoomer/releases/tag/v1.0)

**1. Architecture**

- Setup project base on Clean Architecture
- Add basic functionality
- Setup Echo framework
- Http connect for user service
- Websocket connect for chat service
- Setup database with postgresql
- Store file with minio
- Connect to database
- Dockerize
- Hot reload change with air
- Logging with logrus and export to file
- TLS support
- Load balancer with nginx
- Credential & certificate authentication with JWT
- Setup CI/CD
- Setup HTTP2
- Auto migrate in debug mode
- Connect pool for database

**2. Features**

- Initial connect to client
- Register new user and login
- Create/Join/Leave room
- Send messages to all users in all room

**3. Bugfixes**

- fix: cycle import
- fix: conflict between websocket and http handler
- fix: room-searching context

**4. Documentations**

- Decribe Readme.md
