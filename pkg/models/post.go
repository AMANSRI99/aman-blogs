package models

import (
	"context"
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/gosimple/slug"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title          string
	Content        string
	EncodedContent string
	Status         PostStatus          `gorm:"type:string;default:'draft'"` // Explicitly specifying type
	FormMeta       map[string]FormMeta `gorm:"-"`
	Slug           string              `gorm:"uniqueIndex"`
	Description    string
	Tags           pq.StringArray `gorm:"type:text[]"`
	ReadTime       int
}

type PostStatus string

const (
	Draft     PostStatus = "draft"
	Published PostStatus = "published"
)

type PostRepository interface {
	Upsert(ctx context.Context, value *Post) error
	GetById(ctx context.Context, id uint) (*Post, error)
	Get(ctx context.Context, page int, pageSize int, orderBy string, orderDir string) ([]Post, error)
	GetByStatus(ctx context.Context, status PostStatus, page int, pageSize int) ([]Post, error)
	GetByTag(ctx context.Context, tag string, page int, pageSize int) ([]Post, error)
	GetRelatedPosts(ctx context.Context, postID uint, limit int) ([]Post, error)
	GetPostCount(ctx context.Context, status *PostStatus, tag *string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
}

// Post methods
func (p *Post) BeforeSave(tx *gorm.DB) error {
	if p.Slug == "" {
		p.Slug = GenerateSlug(p.Title)
	}

	if p.ReadTime == 0 {
		p.ReadTime = CalculateReadTime(p.Content)
	}

	if p.Description == "" {
		p.Description = GenerateDescription(p.Content, 160) // 160 characters for SEO
	}

	return nil
}

// Utility functions
func GenerateSlug(title string) string {
	return slug.Make(title)
}

func CalculateReadTime(content string) int {
	words := len(strings.Fields(content))
	wordsPerMinute := 200                                     // Average reading speed
	readTime := (words + wordsPerMinute - 1) / wordsPerMinute // Round up
	if readTime < 1 {
		return 1
	}
	return readTime
}

func GenerateDescription(content string, maxLength int) string {
	// Strip HTML tags (you might want to add a proper HTML sanitizer)
	content = strings.ReplaceAll(content, "\n", " ")

	if utf8.RuneCountInString(content) <= maxLength {
		return content
	}

	runes := []rune(content)
	truncated := string(runes[:maxLength])
	lastSpace := strings.LastIndex(truncated, " ")

	if lastSpace != -1 {
		truncated = truncated[:lastSpace]
	}

	return truncated + "..."
}

// Validation method
func (p *Post) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return ErrEmptyTitle
	}
	if strings.TrimSpace(p.Content) == "" {
		return ErrEmptyContent
	}
	return nil
}

// Error types
var (
	ErrEmptyTitle   = errors.New("title cannot be empty")
	ErrEmptyContent = errors.New("content cannot be empty")
)
