package utils

import (
	"context"
	"github.com/go-playground/validator/v10"
)

var validate validator.validate

func init() {
	validate = validator.New()
}

func ValidateStruct(ctx context.Context, i interface{}) error {
	return validate.StructCtx(ctx, i)
}
