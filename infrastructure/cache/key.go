package cache

func AuthUserKey(authUser string) string {
	return "cache_auth_user#" + authUser
}

func AuthTokenKey(authToken string) string {
	return "cache_auth_token#" + authToken
}

func UserRoomKey(userId string) string {
	return "cache_user_room#" + userId
}

func MailOtpKey(email string) string {
	return "cache_mail_otp#" + email
}
