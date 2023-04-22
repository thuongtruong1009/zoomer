package main

import (
	"net/http"
)

func main() {
	services.InitializeOAuthFacebook()
	services.InitializeOAuthGoogle()

	http.HandleFunc("/facebook-login", services.HandleFacebookLogin)
	http.HandleFunc("/facebook-callback", services.CallBackFromFacebook)
	http.HandleFunc("/google-login", services.HandleGoogleLogin)
	http.HandleFunc("/google-callback", services.CallBackFromGoogle)
}
