package validators

import "github.com/labstack/echo/v4"

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}

	if err := ctx.Validate(request); err != nil {
		return err
	}
	return ValidateStruct(ctx.Request().Context(), request)
}
