package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) ListPublic(query dto.PublicPostListQuery) ([]model.Post, int64, error) {
	var (
		posts []model.Post
		total int64
	)

	db := r.db.Model(&model.Post{}).
		Preload("Tags").
		Preload("Category").
		Where("status = ?", "published")

	if query.Keyword != "" {
		like := "%" + query.Keyword + "%"
		db = db.Where("title ILIKE ? OR summary ILIKE ?", like, like)
	}
	if query.CategoryID > 0 {
		db = db.Where("category_id = ?", query.CategoryID)
	}
	if query.TagID > 0 {
		db = db.Joins("JOIN post_tags ON post_tags.post_id = posts.id").Where("post_tags.tag_id = ?", query.TagID)
	}
	if query.IsTop != nil {
		db = db.Where("is_top = ?", *query.IsTop)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderBy := "is_top DESC, published_at DESC, created_at DESC"
	if strings.EqualFold(query.Sort, "latest") {
		orderBy = "published_at DESC, created_at DESC"
	}

	err := db.Order(orderBy).
		Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *PostRepository) GetPublicBySlug(slug string) (*model.Post, error) {
	var post model.Post
	err := r.db.Preload("Tags").Preload("Category").
		Where("slug = ? AND status = ?", slug, "published").
		First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetPublicByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.db.Preload("Tags").Preload("Category").
		Where("id = ? AND status = ?", id, "published").
		First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) ListAdmin(query dto.AdminPostListQuery) ([]model.Post, int64, error) {
	var (
		posts []model.Post
		total int64
	)

	db := r.db.Model(&model.Post{}).Preload("Category")

	if query.Keyword != "" {
		like := "%" + query.Keyword + "%"
		db = db.Where("title ILIKE ?", like)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.CategoryID > 0 {
		db = db.Where("category_id = ?", query.CategoryID)
	}
	if query.TagID > 0 {
		db = db.Joins("JOIN post_tags ON post_tags.post_id = posts.id").Where("post_tags.tag_id = ?", query.TagID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order("created_at DESC").
		Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *PostRepository) GetAdminByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.db.Preload("Tags").Preload("Category").First(&post, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) SlugExists(slug string, excludeID uint) (bool, error) {
	var count int64
	db := r.db.Model(&model.Post{}).Where("slug = ?", slug)
	if excludeID > 0 {
		db = db.Where("id <> ?", excludeID)
	}
	if err := db.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) Update(post *model.Post) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(post).Error
}

func (r *PostRepository) SaveTags(post *model.Post, tagIDs []uint) error {
	if len(tagIDs) == 0 {
		return r.db.Model(post).Association("Tags").Clear()
	}

	var tags []model.Tag
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.db.Model(post).Association("Tags").Replace(tags)
}

func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&model.Post{}, id).Error
}

func (r *PostRepository) Publish(id uint, publishedAt *time.Time) error {
	update := map[string]interface{}{"status": "published"}
	if publishedAt != nil {
		update["published_at"] = *publishedAt
	} else {
		update["published_at"] = time.Now().UTC()
	}
	return r.db.Model(&model.Post{}).Where("id = ?", id).Updates(update).Error
}

func (r *PostRepository) Unpublish(id uint) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":       "draft",
		"published_at": nil,
	}).Error
}

func (r *PostRepository) SetTop(id uint, isTop bool) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).Update("is_top", isTop).Error
}

func (r *PostRepository) Timeline(year int) ([]model.Post, error) {
	var posts []model.Post
	db := r.db.Model(&model.Post{}).
		Where("status = ?", "published").
		Where("published_at IS NOT NULL")

	if year > 0 {
		start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(1, 0, 0)
		db = db.Where("published_at >= ? AND published_at < ?", start, end)
	}

	err := db.Order("published_at DESC").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&model.Post{}).Count(&count).Error
	return count, err
}

func (r *PostRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Post{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

func (r *PostRepository) CountByUser(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.Post{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *PostRepository) CountByUserAndStatus(userID uint, status string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Post{}).Where("user_id = ? AND status = ?", userID, status).Count(&count).Error
	return count, err
}

func (r *PostRepository) ListUserPosts(ctx context.Context, userID uint, query dto.PublicPostListQuery) ([]model.Post, int64, error) {
	var (
		posts []model.Post
		total int64
	)

	db := r.db.WithContext(ctx).Model(&model.Post{}).
		Preload("Tags").
		Preload("Category").
		Where("user_id = ?", userID)

	if query.Keyword != "" {
		like := "%" + query.Keyword + "%"
		db = db.Where("title ILIKE ? OR summary ILIKE ?", like, like)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order("created_at DESC").
		Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}
