package cache

import "fmt"

func UserIdKey(userId string) string {
	return "cache_uid#" + fmt.Sprint(userId)
}

func UsernameKey(username string) string {
	return "cache_username#" + username
}

func UserRoomKey(userId string) string {
	return "cache_userroom#" + fmt.Sprint(userId)
}

func TokenKey(token string) string {
	return "cache_token#" + token
}
