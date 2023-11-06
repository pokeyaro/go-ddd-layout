package orm

import "gorm.io/gorm"

const (
	MaxPageSize       = 100
	DefaultPageSize   = 20
	DefaultPageNumber = 1
)

// Paginator returns a function that applies pagination to the query using the given page size and number.
func Paginator(pageSize, pageNumber int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if pageNumber <= 0 {
			pageNumber = DefaultPageNumber
		}

		switch {
		case pageSize > MaxPageSize:
			pageSize = MaxPageSize
		case pageSize <= 0:
			pageSize = DefaultPageSize
		}

		offset := (pageNumber - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
