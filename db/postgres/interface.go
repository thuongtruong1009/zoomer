package postgres

import (
	"gorm.io/gorm"
	"github.com/thuongtruong1009/zoomer/configs"
)

type PgAdapter interface {
	GetInstance(*configs.Configuration) *gorm.DB

	RetryHandler(int, func() (bool, error)) error

	SetConnectionPool(*gorm.DB, *configs.Configuration)

	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}
