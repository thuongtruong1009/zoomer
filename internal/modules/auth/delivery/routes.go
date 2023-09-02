package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/middlewares"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapAuthRoutes(authGroup *echo.Group, h AuthHandler, mw *middlewares.AuthMiddleware) {
	authGroup.POST(constants.SignUpEndPoint, h.SignUp())
	authGroup.POST(constants.SignInEndPoint, h.SignIn())
	authGroup.POST(constants.SignOutEndPoint, h.SignOut(), mw.JWTValidation)
	authGroup.POST(constants.ForgotPassword, h.ForgotPassword())
	authGroup.POST(constants.VerifyResetPasswordOtp, h.VerifyResetPasswordOtp())
	authGroup.PATCH(constants.ResetPassword, h.ResetPassword())
	authGroup.PATCH(constants.UpdatePassword, h.UpdatePassword(), mw.JWTValidation)
}
