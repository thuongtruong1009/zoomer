package models

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email string `json:"email"`
	UserId   string `json:"userId"`
}
