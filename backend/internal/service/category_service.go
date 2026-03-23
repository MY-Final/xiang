package service

import (
	"errors"
	"strings"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	"xiang/backend/internal/repository"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) List() ([]model.Category, error) {
	return s.repo.List()
}

func (s *CategoryService) Create(req dto.SaveCategoryRequest) (*model.Category, error) {
	category := &model.Category{
		Name:        strings.TrimSpace(req.Name),
		Slug:        strings.TrimSpace(req.Slug),
		Description: strings.TrimSpace(req.Description),
	}
	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Update(id uint, req dto.SaveCategoryRequest) error {
	if current, err := s.repo.FindByID(id); err != nil {
		return err
	} else if current == nil {
		return ErrCategoryNotFound
	}

	category := &model.Category{
		Name:        strings.TrimSpace(req.Name),
		Slug:        strings.TrimSpace(req.Slug),
		Description: strings.TrimSpace(req.Description),
	}
	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id uint) error {
	if current, err := s.repo.FindByID(id); err != nil {
		return err
	} else if current == nil {
		return ErrCategoryNotFound
	}
	return s.repo.Delete(id)
}

func (s *CategoryService) CountAll() (int64, error) {
	return s.repo.CountAll()
}
