package service

import (
	"errors"
	"strings"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	"xiang/backend/internal/repository"
)

var ErrTagNotFound = errors.New("tag not found")

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) List() ([]model.Tag, error) {
	return s.repo.List()
}

func (s *TagService) Create(req dto.SaveTagRequest) (*model.Tag, error) {
	tag := &model.Tag{
		Name: strings.TrimSpace(req.Name),
		Slug: strings.TrimSpace(req.Slug),
	}
	if err := s.repo.Create(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *TagService) Update(id uint, req dto.SaveTagRequest) error {
	if ok, err := s.repo.Exists(id); err != nil {
		return err
	} else if !ok {
		return ErrTagNotFound
	}

	tag := &model.Tag{
		Name: strings.TrimSpace(req.Name),
		Slug: strings.TrimSpace(req.Slug),
	}
	return s.repo.Update(id, tag)
}

func (s *TagService) Delete(id uint) error {
	if ok, err := s.repo.Exists(id); err != nil {
		return err
	} else if !ok {
		return ErrTagNotFound
	}
	return s.repo.Delete(id)
}

func (s *TagService) CountAll() (int64, error) {
	return s.repo.CountAll()
}
