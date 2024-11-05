package repository

import (
	"context"
	"log"

	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"github.com/AMANSRI99/aman-blogs/utils/scopes"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type repository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) models.PostRepository {
	return &repository{db: db}
}

func (r *repository) Get(ctx context.Context, page int, pageSize int, orderBy string, orderDir string) ([]models.Post, error) {
	var res []models.Post

	opts := scopes.Query_opts{
		Page:           page,
		PageSize:       pageSize,
		Order:          orderBy,
		OrderDirection: scopes.Direction(orderDir),
	}

	err := r.db.WithContext(ctx).
		Scopes(scopes.Apply(opts)).
		Find(&res).Error

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) GetById(ctx context.Context, id uint) (*models.Post, error) {
	var res models.Post
	err := r.db.WithContext(ctx).First(&res, id).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *repository) Upsert(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *repository) GetByStatus(ctx context.Context, status models.PostStatus, page int, pageSize int) ([]models.Post, error) {
	var res []models.Post

	opts := scopes.Query_opts{
		Page:           page,
		PageSize:       pageSize,
		Order:          "created_at",
		OrderDirection: scopes.Descending,
	}

	db := r.db.Session(&gorm.Session{
		PrepareStmt: false,
		Logger:      r.db.Logger.LogMode(logger.Info),
	}).WithContext(ctx)

	// Log the query being executed
	log.Printf("Executing GetByStatus query with status: %s, page: %d, pageSize: %d", status, page, pageSize)

	err := db.Where("status = ?", string(status)).
		Scopes(scopes.Apply(opts)).
		Find(&res).Error

	if err != nil {
		log.Printf("Error in GetByStatus: %v", err)
		return nil, err
	}

	log.Printf("Found %d posts", len(res))
	return res, nil
}

func (r *repository) GetByTag(ctx context.Context, tag string, page int, pageSize int) ([]models.Post, error) {
	var res []models.Post

	opts := scopes.Query_opts{
		Page:           page,
		PageSize:       pageSize,
		Order:          "created_at",
		OrderDirection: scopes.Descending,
	}

	err := r.db.WithContext(ctx).
		Where("? = ANY(tags)", tag).
		Scopes(scopes.Apply(opts)).
		Find(&res).Error

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) GetRelatedPosts(ctx context.Context, postID uint, limit int) ([]models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return nil, err
	}

	var res []models.Post
	err := r.db.WithContext(ctx).
		Where("id != ? AND tags && ?", postID, post.Tags).
		Limit(limit).
		Order("created_at DESC").
		Find(&res).Error

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) GetPostCount(ctx context.Context, status *models.PostStatus, tag *string) (int64, error) {
	query := r.db.WithContext(ctx).Model(&models.Post{})

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if tag != nil {
		query = query.Where("? = ANY(tags)", *tag)
	}

	var count int64
	err := query.Count(&count).Error
	return count, err
}

func (r *repository) GetBySlug(ctx context.Context, slug string) (*models.Post, error) {
	var post models.Post // Create a post variable

	err := r.db.WithContext(ctx).
		Where("slug = ? AND status = ?", slug, models.Published).
		First(&post).Error // GORM needs address to populate the struct

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Can return nil for not found
		}
		return nil, err
	}

	return &post, nil // Return pointer to the post
}
