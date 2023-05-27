package constants

import "errors"

var (
	// auth
	ErrUserNotFound       error = errors.New("user not found")
	ErrWrongPassword      error = errors.New("wrong password")
	ErrUserExisted        error = errors.New("user existed")
	ErrInvalidAccessToken error = errors.New("invalid access token")
	ErrNoRecord           error = errors.New("no record")
	ErrUsernameInvalid    error = errors.New("username invalid")
	ErrPasswordInvalid    error = errors.New("password invalid")

	// room
	ErrChatNotFound     error = errors.New("chat not found")
	ErrChatAccessDenied error = errors.New("chat access denied")

	// stream
	ErrStreamIDMissing error = errors.New("roomID is missing")

	//other
	ErrorInternalServer      error = errors.New("We had a problem with our server. Try again later")
	ErrorBadRequest          error = errors.New("Something went wrong")
	ErrorUnauthorized        error = errors.New("Your API key is wrong. Try re-authenticating")
	ErrorForbidden           error = errors.New("You do not have permission to access this resource")
	ErrorNotFound            error = errors.New("The specified resource could not be found")
	ErrorMethodNotAllowed    error = errors.New("You tried to access a resource with an invalid method")
	ErrorNotAcceptable       error = errors.New("You requested a format that isn't json")
	ErrorUnprocessableEntity error = errors.New("Your input failed validation")
	ErrorTooManyRequests     error = errors.New("Too many requests")
	// ErrorServiceUnavailable = errors.New("We're temporarily offline for maintenance. Please try again later")
)
