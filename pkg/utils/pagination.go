package utils

import "gorm.io/gorm"

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		// if offset below 0, set to 1
		if page <= 0 {
			page = 1
		}

		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset = (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
