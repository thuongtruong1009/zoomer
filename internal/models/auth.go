package models

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
	Id   string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

// type AuthCtxToken struct {
// 	UserId   string `json:"userId"`
// 	Username string `json:"username"`
// 	Email string `json:"email"`
// }
