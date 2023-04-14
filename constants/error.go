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
	ErrorInternalServer = errors.New("Error occured in our own server")
	ErrorBadRequest   = errors.New("Bad request")
)
