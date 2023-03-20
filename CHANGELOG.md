There is the most valuable changes log:

### v1.0.0

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

- Send message to specific user in specific room
- Send message to all users in specific room

**3. Bugfixes**

- fix: cycle import
- fix: conflict between websocket and http handler
- fix: room-searching context

**4. Documentations**

- Decribe Readme.md
