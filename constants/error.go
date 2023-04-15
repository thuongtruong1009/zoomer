package constants

import "errors"

var (
	// auth
	ErrUserNotFound       = errors.New("user not found")
	ErrWrongPassword      = errors.New("wrong password")
	ErrUserExisted        = errors.New("user existed")
	ErrInvalidAccessToken = errors.New("invalid access token")

	// room
	ErrChatNotFound     = errors.New("chat not found")
	ErrChatAccessDenied = errors.New("chat access denied")

	//other
	ErrorInternalServer = errors.New("We had a problem with our server. Try again later")
	ErrorBadRequest   = errors.New("Something went wrong")
	ErrorUnauthorized = errors.New("Your API key is wrong. Try re-authenticating")
	ErrorForbidden    = errors.New("You do not have permission to access this resource")
	ErrorNotFound     = errors.New("The specified resource could not be found")
	ErrorMethodNotAllowed = errors.New("You tried to access a resource with an invalid method")
	ErrorNotAcceptable = errors.New("You requested a format that isn't json")
	ErrorUnprocessableEntity = errors.New("Your input failed validation")
	ErrorTooManyRequests = errors.New("Too many requests")
	// ErrorServiceUnavailable = errors.New("We're temporarily offline for maintenance. Please try again later")
)
