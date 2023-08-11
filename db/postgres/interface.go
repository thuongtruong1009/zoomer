package postgres

import (
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"gorm.io/gorm"
)

type PgAdapter interface {
	getInstance(string) *gorm.DB

	ConnectInstance(*configs.Configuration) *gorm.DB

	retryHandler(int, func() (bool, error)) error

	setConnectionPool(*gorm.DB)
}
