package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/repository"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"net/http"
	"strings"
	"time"
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

func (a *authUseCase) SignUp(ctx context.Context, username string, password string, limit int) (*presenter.SignUpResponse, error) {
	fmtusername := strings.ToLower(username)
	euser, _ := a.userRepo.GetUserByUsername(ctx, fmtusername)

	if euser != nil {
		exceptions.Log(constants.ErrUserExisted, nil)
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
		exceptions.Log(constants.ErrCreateUserFailed, err)
		return nil, err
	}

	return &presenter.SignUpResponse{
		Id:       user.Id,
		Username: user.Username,
		Limit:    user.Limit,
	}, nil
}

func (a *authUseCase) SignIn(ctx context.Context, username, password string) (*presenter.SignInResponse, error) {
	user, _ := a.userRepo.GetUserByUsername(ctx, username)
	if user == nil {
		exceptions.Log(constants.ErrUserNotFound, nil)
		return nil, constants.ErrUserNotFound
	}

	if !user.ComparePassword(password) {
		exceptions.Log(constants.ErrWrongPassword, nil)
		return nil, constants.ErrWrongPassword
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
		exceptions.Log(constants.ErrSigningKey, err)
		return nil, err
	}

	res := &presenter.SignInResponse{
		UserId:   user.Id,
		Username: user.Username,
		Token:    tmp,
	}

	return res, nil
}

func (a *authUseCase) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			exceptions.Log(constants.ErrUnexpectedSigning, token.Header["alg"])
			return nil, fmt.Errorf("%s: %v", constants.ErrUnexpectedSigning, token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		exceptions.Log(constants.ErrParseToken, err)
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	exceptions.Log(constants.ErrInvalidAccessToken, nil)
	return "", constants.ErrInvalidAccessToken
}

func WriteCookie(c echo.Context, name, value string, expire time.Duration, path, domain string, secure, httpOnly bool) {
	cookie := http.Cookie{
		Name:     name,
		Expires:  time.Now().Add(expire),
		Value:    value,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	}

	c.SetCookie(&cookie)
}

// func (a *authUseCase) SearchUserByMatch(c echo.Context, username string) {
// 	users, err := a.userRepo.QueryMatchingFields(c.Request().Context(), username)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// }
