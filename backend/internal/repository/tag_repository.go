package repository

import (
	"errors"

	"xiang/backend/internal/model"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) List() ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Order("id DESC").Find(&tags).Error
	return tags, err
}

func (r *TagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) Update(id uint, tag *model.Tag) error {
	return r.db.Model(&model.Tag{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name": tag.Name,
		"slug": tag.Slug,
	}).Error
}

func (r *TagRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

func (r *TagRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Tag{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *TagRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&model.Tag{}).Count(&count).Error
	return count, err
}

func (r *TagRepository) FindByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}
