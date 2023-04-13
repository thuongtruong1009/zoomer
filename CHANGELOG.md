There is the most valuable changes log:

### v1.1

**1.Architecture**

- Connect to Redis database for chat service
- Connect done to client
- Config Cache adapter with Redis

**2. Features**

- Add/Connect to new user
- Send real-time messages to user with timestamp record
- Save/Load messages to database
- Rejoin room but not remove old messages
- Load recent friend contact list
- Cache response userid, username, roomlist

**3. Bugfixes**

- fix: reload old messages when rejoin room
- fix: reconnect to websocket when rejoin room

### v1.0

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
- Apply HTTP2
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
