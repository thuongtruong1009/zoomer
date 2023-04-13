package usecase

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"zoomer/constants"
	"zoomer/internal/auth/repository"
	"zoomer/internal/models"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type authUseCase struct {
	userRepo       repository.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo repository.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL int64) UseCase {
	return &authUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * time.Duration(tokenTTL),
	}
}

func (a *authUseCase) SignUp(ctx context.Context, username string, password string, limit int) (*models.User, error) {
	fmtusername := strings.ToLower(username)
	euser, _ := a.userRepo.GetUserByUsername(ctx, fmtusername)

	if euser != nil {
		return nil, constants.ErrUserExisted
	}

	user := &models.User{
		Id:       uuid.New().String(),
		Username: fmtusername,
		Password: password,
		Limit:    limit,
	}

	user.HashPassword()
	err := a.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return a.userRepo.GetUserByUsername(ctx, fmtusername)
}

func (a *authUseCase) SignIn(ctx context.Context, username, password string) (string, string, string, error) {
	user, _ := a.userRepo.GetUserByUsername(ctx, username)
	if user == nil {
		return "", "", "", constants.ErrUserNotFound
	}

	if !user.ComparePassword(password) {
		return "", "", "", constants.ErrWrongPassword
	}

	claims := AuthClaims{
		Username: user.Username,
		UserId:   user.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    user.Id,
			ExpiresAt: time.Now().Add(a.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tmp, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", "", "", err
	}

	return user.Id, user.Username, tmp, err
}

func (a *authUseCase) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", constants.ErrInvalidAccessToken
}

func WriteCookie(c echo.Context, name string, value string, expire time.Duration, path string, domain string, secure bool, httpOnly bool) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(expire),
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	}
	c.SetCookie(&cookie)
}
