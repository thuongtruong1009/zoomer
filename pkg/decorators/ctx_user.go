package decorators

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func BindCurrentUser(ctx echo.Context, user *models.AuthClaims) error {
	ctx.Set(constants.AuthCtxUserKey, user)
	return nil
}

func DetectCurrentUser(ctx echo.Context) *models.AuthClaims {
	user := ctx.Get(constants.AuthCtxUserKey)
	if user == nil {
		return nil
	}

	return user.(*models.AuthClaims)
}
