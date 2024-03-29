package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/thuongtruong1009/zoomer/infrastructure/cache"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/infrastructure/mail"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	authRepository "github.com/thuongtruong1009/zoomer/internal/modules/auth/repository"
	userRepository "github.com/thuongtruong1009/zoomer/internal/modules/users/repository"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"strings"
	"time"
)

type authUsecase struct {
	authRepo authRepository.UserRepository
	userRepo userRepository.IUserRepository
	cfg      *configs.Configuration
	paramCfg *parameter.ParameterConfig
	mail     mail.IMail
}

func NewAuthUseCase(
	authRepo authRepository.UserRepository,
	userRepo userRepository.IUserRepository,
	cfg *configs.Configuration,
	paramCfg *parameter.ParameterConfig,
	mail mail.IMail,
) UseCase {
	return &authUsecase{
		authRepo: authRepo,
		userRepo: userRepo,
		cfg:      cfg,
		paramCfg: paramCfg,
		mail:     mail,
	}
}

func (a *authUsecase) SignUp(ctx context.Context, dto *presenter.SignUpRequest) (*presenter.SignUpResponse, error) {
	fmtusername := strings.ToLower(dto.Username)

	euser, _ := a.userRepo.GetUserByIdOrName(ctx, fmtusername)
	if euser != nil {
		exceptions.Log(constants.ErrUserExisted, nil)
		return nil, constants.ErrUserExisted
	}

	user := &models.User{
		Id:       uuid.New().String(),
		Username: fmtusername,
		Email:    dto.Email,
		Password: dto.Password,
		Limit:    dto.Limit,
	}

	hashedPassword, err := helpers.Encrypt(user.Password)
	if err != nil {
		exceptions.Log(constants.ErrHashPassword, err)
		return nil, err
	}

	user.Password = string(hashedPassword)

	if err := a.authRepo.CreateUser(ctx, user); err != nil {
		exceptions.Log(constants.ErrCreateUserFailed, err)
		return nil, err
	}

	return &presenter.SignUpResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Limit:    user.Limit,
	}, nil
}

func (au *authUsecase) SignIn(ctx context.Context, dto *presenter.SignInRequest) (*presenter.SignInResponse, error) {
	user, err := au.userRepo.GetUserByIdOrName(ctx, dto.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	if err := helpers.Decrypt(user.Password, dto.Password); err != nil {
		exceptions.Log(constants.ErrComparePassword, err)
		return nil, err
	}

	res := &presenter.SignInResponse{
		UserId:   user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    "",
	}

	cachekey := cache.AuthTokenKey(user.Id + user.Username)
	userInCache := cache.GetCache(cachekey)
	if userInCache != nil {
		res.Token = userInCache.(string)
	} else {
		claims := models.AuthClaims{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				Issuer:    user.Id,
				ExpiresAt: time.Now().Add(helpers.DurationSecond(au.paramCfg.TokenTimeout)).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tmp, err := token.SignedString([]byte(au.cfg.SigningKey))
		if err != nil {
			exceptions.Log(constants.ErrSigningKey, err)
			return nil, constants.ErrSigningKey
		}

		res.Token = tmp
		cache.SetCache(cachekey, tmp, helpers.DurationSecond(au.paramCfg.TokenTimeout))
	}

	return res, nil
}

func (au *authUsecase) ParseToken(ctx context.Context, accessToken string) (*models.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			exceptions.Log(constants.ErrUnexpectedSigning, token.Header["alg"])
			return nil, fmt.Errorf("%s: %v", constants.ErrUnexpectedSigning, token.Header["alg"])
		}
		return []byte(au.cfg.SigningKey), nil
	})

	if err != nil {
		exceptions.Log(constants.ErrParseToken, err)
		return nil, constants.ErrParseToken
	}

	if claims, ok := token.Claims.(*models.AuthClaims); ok && token.Valid {
		return claims, nil
	}

	exceptions.Log(constants.ErrInvalidAccessToken, nil)
	return nil, constants.ErrInvalidAccessToken
}

func (au *authUsecase) ForgotPassword(ctx context.Context, email string) error {
	newOtp := helpers.RandomChain(constants.RandomTypeNumber, 6)

	newEmail := &mail.Mail{
		To:      email,
		Subject: "Reset Zoomer password",
		Body:    "Your OTP code is: " + newOtp,
	}

	cache.SetCache(cache.MailOtpKey(email), newOtp, helpers.DurationSecond(au.paramCfg.OtpTimeout))

	return au.mail.SendingMail(newEmail)
}

func (au *authUsecase) VerifyOtp(ctx context.Context, otpCode string) error {
	sentOtp := cache.GetCache(cache.MailOtpKey(otpCode))

	if sentOtp == nil {
		exceptions.Log(constants.ErrOtpExpired, nil)
		return constants.ErrOtpExpired
	}

	if sentOtp.(string) != otpCode {
		exceptions.Log(constants.ErrOtpInvalid, nil)
		return constants.ErrOtpInvalid
	}

	return nil
}

func (au *authUsecase) ResetPassword(ctx context.Context, dto *presenter.UpdatePassword) error {
	return au.authRepo.UpdatePassword(ctx, dto.Email, dto.NewPassword)
}

func (au *authUsecase) UpdatePassword(ctx context.Context, dto *presenter.UpdatePassword) error {
	return au.authRepo.UpdatePassword(ctx, dto.Email, dto.NewPassword)
}
