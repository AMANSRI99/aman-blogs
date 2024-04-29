package pagination

import (
	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := pageSize * (page - 1)
		return db.Offset(offset).Limit(pageSize)
	}
}

// In this approach, you first define the pagination function using Paginate. This function encapsulates the pagination logic but does not immediately apply it to any database query. Later in your code, when you have a GORM database instance and want to apply pagination, you call the pagination function with the database instance (db). This executes the pagination logic and returns the modified database instance, which you can then use to perform the actual database query.
// By returning a function instead of immediately modifying the database instance, you can defer the execution of pagination logic until it's needed later in the code. This provides flexibility in applying pagination dynamically based on different conditions or criteria.
