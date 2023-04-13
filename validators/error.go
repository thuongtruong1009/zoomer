package validators

import (
	"gorm.io/gorm"
)

func DBError(db *gorm.DB) {
	if err := db.Error; err != nil {
		panic(err)
	}
}
