package scopes

import (
	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"gorm.io/gorm"
)

func WithStatus(status models.PostStatus) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

func WithTag(tag string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("? = ANY(tags)", tag)
	}
}

func WithSearch(query string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(
			"title ILIKE ? OR content ILIKE ?",
			"%"+query+"%",
			"%"+query+"%",
		)
	}
}
