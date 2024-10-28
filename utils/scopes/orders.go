package scopes

import (
	"fmt"

	"gorm.io/gorm"
)

func OrderBy(field string, direction Direction) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", field, direction))
	}
}
