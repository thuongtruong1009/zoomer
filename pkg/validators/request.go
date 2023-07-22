package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}

	if err := validate.StructCtx(ctx.Request().Context(), request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

