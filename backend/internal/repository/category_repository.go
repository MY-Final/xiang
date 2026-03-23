package repository

import (
	"errors"

	"xiang/backend/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) List() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Order("id DESC").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(id uint, category *model.Category) error {
	return r.db.Model(&model.Category{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        category.Name,
		"slug":        category.Slug,
		"description": category.Description,
	}).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}

func (r *CategoryRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&model.Category{}).Count(&count).Error
	return count, err
}

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}
