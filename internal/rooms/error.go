package chats

import "errors"

var (
	ErrChatNotFound     = errors.New("chat not found")
	ErrChatAccessDenied = errors.New("chat access denied")
)
