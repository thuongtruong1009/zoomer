package validators

import (
	"context"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// type CustomValidator struct {
// 	validate *validator.Validate
// }

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(ctx context.Context, i interface{}) error {
	if err := validate.StructCtx(ctx, i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
