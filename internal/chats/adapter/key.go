package adapter

import (
	"fmt"
	"time"
)

func UserSetKey() string {
	return "users"
}

func SessionKey(client string) string {
	return "session#" + client
}

func ChatKey() string {
	return fmt.Sprintf("chat#%d", time.Now().UnixMilli())
}

func ChatIndex() string {
	return "idx#chats"
}

func ContactListZKey(username string) string {
	return "contacts:" + username
}
