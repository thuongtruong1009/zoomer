package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
	db "github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/auth/presenter"
	"github.com/thuongtruong1009/zoomer/internal/auth/repository"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type authUseCase struct {
	pgAdapter	  db.PgAdapter
	userRepo       repository.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	pgAdapter db.PgAdapter,
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

	_, err := a.pgAdapter.Transaction(func(i interface{})  (interface{}, error) {
		err := a.userRepo.CreateUser(ctx, user)
		return nil, err
	})
	// user, ok := data.(*model.User)

	// if !ok {
	// 	return nil, errors.New("cast error")
	// }

	if err != nil {
		return nil, err
	}

	return a.userRepo.GetUserByUsername(ctx, fmtusername)
}

func (a *authUseCase) SignIn(ctx context.Context, username, password string) (*presenter.LogInResponse, error) {
	user, _ := a.userRepo.GetUserByUsername(ctx, username)
	if user == nil {
		return nil, constants.ErrUserNotFound
	}

	if !user.ComparePassword(password) {
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

	res := &presenter.LogInResponse{
		UserId:   user.Id,
		Username: user.Username,
		Token:    tmp,
	}
	if err != nil {
		return nil, err
	}

	return res, nil
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
