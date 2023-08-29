package constants

import "errors"

var (
	ErrorShuttdownServer   error = errors.New("error shutting down server")
	ErrorStartHttps        error = errors.New("error occured when starting the server in HTTPS mode")
	ErrorStartHttp         error = errors.New("error occurred while starting the http api")
	ErrorSetupHttpRouter   error = errors.New("error occurred while setting up http routers")
	ErrorSetupSocketRouter error = errors.New("error occurred while setting up websocket routers")
	ErrorStartupApi        error = errors.New("error occurred while starting up the server")
	ErrorLoadEnvFile       error = errors.New("error loading env file")
	ErrorEnvKeyNotFound    error = errors.New("env key not found")
	ErrorContextTimeout    error = errors.New("error context timeout")
)

var (
	ErrorInternalServer      error = errors.New("we had a problem with our server. Try again later")
	ErrorBadRequest          error = errors.New("something went wrong")
	ErrorUnauthorized        error = errors.New("your API key is wrong. Try re-authenticating")
	ErrorForbidden           error = errors.New("you do not have permission to access this resource")
	ErrorNotFound            error = errors.New("the specified resource could not be found")
	ErrorMethodNotAllowed    error = errors.New("you tried to access a resource with an invalid method")
	ErrorNotAcceptable       error = errors.New("you requested a format that isn't json")
	ErrorUnprocessableEntity error = errors.New("your input failed validation")
	ErrorTooManyRequests     error = errors.New("too many requests")
	ErrorTimeout             error = errors.New("your request timed out")
	ErrorNoContent           error = errors.New("no content")
)

var (
	ErrorRedisConnectionFailed error = errors.New("redis connection failed")
	ErrRedisSyncUser           error = errors.New("error when sync user data to redis")
	ErrRedisAddUser            error = errors.New("error when add redis user data to redis")

	ErrorPostgresConnectionFailed error = errors.New("failed to connect to postgres database")
	ErrorPostgresGetResponse      error = errors.New("error when get ping response from postgres")
	ErrorPostgresReconnect        error = errors.New("error when reconnect to postgres")
	ErrorPostgresAutoMigration    error = errors.New("error when run auto migrate postgres")
)

var (
	ErrRequiredUUID error = errors.New("id is required")
	ErrInvalidUUID  error = errors.New("uuid is invalid")
)

var (
	ErrUserNotFound     error = errors.New("user not found")
	ErrUserExisted      error = errors.New("user existed")
	ErrReqiredUsername  error = errors.New("username is required")
	ErrLenUsername      error = errors.New("username must be between 4 and 20 characters")
	ErrAlphaNumUsername error = errors.New("username can only contain letters or numbers")
	ErrCreateUserFailed error = errors.New("create user failed")
	ErrSpaceUsername    error = errors.New("username is not allowed to have spaces")

	ErrRequiredEmail error = errors.New("email is required")
	ErrInvalidEmail  error = errors.New("email is invalid")
	ErrSpaceEmail    error = errors.New("email is not allowed to have spaces")
	ErrLenEmail      error = errors.New("email must be between 8 and 20 characters")

	ErrRequiredPassword error = errors.New("password is required")
	ErrSpacePassword    error = errors.New("password is not allowed to have spaces")
	ErrLenPassword      error = errors.New("password must be between 8 and 32 characters")
	ErrAlphaNumPassword error = errors.New("password must contain at least one letter - one number - one special characters - one uppercase letter - one lowercase letter")
	ErrHashPassword     error = errors.New("error when encrypt password")
	ErrComparePassword  error = errors.New("password not match")

	ErrNoRecord           error = errors.New("no record")
	ErrInvalidAccessToken error = errors.New("invalid access token")
	ErrUnexpectedSigning  error = errors.New("unexpected signing method")
	ErrSigningKey         error = errors.New("signing key error")
	ErrParseToken         error = errors.New("parse token error")
)

var (
	// room
	ErrChatNotFound     error = errors.New("chat not found")
	ErrChatAccessDenied error = errors.New("chat access denied")
	ErrInvalidRoomLimit error = errors.New("room limit must be positive number")
	ErrMaxRoomLimit     error = errors.New("max room limit")
)

var (
	// stream
	ErrStreamIDMissing error = errors.New("roomID is missing")
)
