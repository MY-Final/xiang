package repository

import (
	"errors"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"

	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

func (r *MediaRepository) Create(media *model.Media) error {
	return r.db.Create(media).Error
}

func (r *MediaRepository) List(query dto.MediaListQuery) ([]model.Media, int64, error) {
	var (
		items []model.Media
		total int64
	)

	db := r.db.Model(&model.Media{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order("id DESC").
		Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).
		Find(&items).Error
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *MediaRepository) FindByID(id uint) (*model.Media, error) {
	var media model.Media
	err := r.db.First(&media, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &media, nil
}

func (r *MediaRepository) Delete(id uint) error {
	return r.db.Delete(&model.Media{}, id).Error
}

func (r *MediaRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&model.Media{}).Count(&count).Error
	return count, err
}
