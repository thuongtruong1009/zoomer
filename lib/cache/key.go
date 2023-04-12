package cache

import "fmt"

func UserIdKey(userId string) string {
	return "user_id#" + fmt.Sprint(userId)
}

func UsernameKey(username string) string {
	return "user_username#" + username
}

func UserRoomKey(userId string) string {
	return "user_room#" + fmt.Sprint(userId)
}
