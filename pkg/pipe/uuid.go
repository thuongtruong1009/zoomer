package pipe

import (
	"github.com/google/uuid"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func IsValidUUID(input string) error {
	if input == "" {
		return constants.ErrRequiredUUID
	}

	_, err := uuid.Parse(input)
	if err != nil {
		return constants.ErrInvalidUUID
	}

	return nil
}
