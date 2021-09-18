package scope

import (
	"fmt"

	"gorm.io/gorm"
)

func SimpleScope(field string, value interface{}, params ...string) func(db *gorm.DB) *gorm.DB {
	if len(params) == 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s %s ?", field, params[0]), value)
	}
}
