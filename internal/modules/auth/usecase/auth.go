package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	authRepository "github.com/thuongtruong1009/zoomer/internal/modules/auth/repository"
	userRepository "github.com/thuongtruong1009/zoomer/internal/modules/users/repository"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/infrastructure/cache"
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
	authRepo       authRepository.UserRepository
	userRepo 	 userRepository.IUserRepository
	cfg *configs.Configuration
	paramCfg *parameter.ParameterConfig
}

func NewAuthUseCase(
	authRepo authRepository.UserRepository,
	userRepo userRepository.IUserRepository,
	cfg *configs.Configuration,
	paramCfg *parameter.ParameterConfig,
	) UseCase {
	return &authUseCase{
		authRepo:       authRepo,
		userRepo: userRepo,
		cfg: cfg,
		paramCfg: paramCfg,
	}
}

func (a *authUseCase) SignUp(ctx context.Context, username string, password string, limit int) (*presenter.SignUpResponse, error) {
	fmtusername := strings.ToLower(username)
	euser, _ := a.userRepo.GetUserByIdOrName(ctx, fmtusername)

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

	err := a.authRepo.CreateUser(ctx, user)

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

func (au *authUseCase) SignIn(ctx context.Context, username, password string) (*presenter.SignInResponse, error) {
	user, _ := au.userRepo.GetUserByIdOrName(ctx, username)
	if user == nil {
		exceptions.Log(constants.ErrUserNotFound, nil)
		return nil, constants.ErrUserNotFound
	}

	if !user.ComparePassword(password) {
		exceptions.Log(constants.ErrWrongPassword, nil)
		return nil, constants.ErrWrongPassword
	}

	res := &presenter.SignInResponse{
		UserId:   user.Id,
		Username: user.Username,
		Token:    "",
	}

	cachekey := cache.TokenKey(user.Id + user.Username)
	userInCache := cache.GetCache(cachekey)
	if userInCache != nil {
		res.Token = userInCache.(string)
	} else {
		claims := AuthClaims{
			Username: user.Username,
			UserId:   user.Id,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				Issuer:    user.Id,
				ExpiresAt: time.Now().Add(time.Second * time.Duration(au.paramCfg.TokenTimeout)).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tmp, err := token.SignedString([]byte(au.cfg.SigningKey))
		if err != nil {
			exceptions.Log(constants.ErrSigningKey, err)
			return nil, err
		}

		res.Token = tmp
		cache.SetCache(cachekey, tmp, time.Second * time.Duration(au.paramCfg.TokenTimeout))
	}

	return res, nil
}

func (au *authUseCase) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			exceptions.Log(constants.ErrUnexpectedSigning, token.Header["alg"])
			return nil, fmt.Errorf("%s: %v", constants.ErrUnexpectedSigning, token.Header["alg"])
		}
		return []byte(au.cfg.SigningKey), nil
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
// 	users, err := a.authRepo.QueryMatchingFields(c.Request().Context(), username)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// }
