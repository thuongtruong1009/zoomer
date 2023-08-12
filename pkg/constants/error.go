package constants

import "errors"

var (
	ErrorShuttdownServer   error = errors.New("error shutting down server")
	ErrorStartHttps        error = errors.New("error occured when starting the server in HTTPS mode")
	ErrorStartHttp         error = errors.New("error occurred while starting the http api")
	ErrorSetupHttpRouter   error = errors.New("error occurred while setting up http routers")
	ErrorSetupSocketRouter error = errors.New("error occurred while setting up websocket routers")
	ErrorStartupApi		   error = errors.New("error occurred while starting up the server")

	ErrorRedisConnectionFailed error = errors.New("redis connection failed")
	ErrorLoadEnvFile           error = errors.New("error loading env file")
	ErrorEnvKeyNotFound        error = errors.New("env key not found")
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
	// auth
	ErrUserNotFound       error = errors.New("user not found")
	ErrWrongPassword      error = errors.New("wrong password")
	ErrUserExisted        error = errors.New("user existed")
	ErrInvalidAccessToken error = errors.New("invalid access token")
	ErrNoRecord           error = errors.New("no record")
	ErrUsernameInvalid    error = errors.New("username invalid")
	ErrPasswordInvalid    error = errors.New("password invalid")
	ErrUnexpectedSigning  error = errors.New("unexpected signing method")
	ErrSigningKey         error = errors.New("signing key error")
	ErrParseToken         error = errors.New("parse token error")
	ErrCreateUserFailed   error = errors.New("create user failed")

	// room
	ErrChatNotFound     error = errors.New("chat not found")
	ErrChatAccessDenied error = errors.New("chat access denied")

	// stream
	ErrStreamIDMissing error = errors.New("roomID is missing")
)
