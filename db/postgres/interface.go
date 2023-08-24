package postgres

import "gorm.io/gorm"

type PgAdapter interface {
	getInstance(string) *gorm.DB

	ConnectInstance() *gorm.DB

	retryHandler(int, func() (bool, error)) error

	setConnectionPool(*gorm.DB)
}
