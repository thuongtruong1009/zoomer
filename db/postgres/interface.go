package postgres

import (
	"gorm.io/gorm"
	"github.com/thuongtruong1009/zoomer/configs"
)

type PgAdapter interface {
	getInstance(string) *gorm.DB

	ConnectInstance(*configs.Configuration) *gorm.DB

	retryHandler(int, func() (bool, error)) error

	setConnectionPool(*gorm.DB, *configs.Configuration)
}
