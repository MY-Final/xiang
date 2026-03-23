package service

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/repository"
)

var ErrMediaNotFound = errors.New("media not found")

type MediaService struct {
	repo      *repository.MediaRepository
	uploadDir string
}

func NewMediaService(repo *repository.MediaRepository, uploadDir string) *MediaService {
	return &MediaService{repo: repo, uploadDir: uploadDir}
}

func (s *MediaService) List(query dto.MediaListQuery) (interface{}, error) {
	query.Normalize()
	list, total, err := s.repo.List(query)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"list":      list,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	}, nil
}

func (s *MediaService) Delete(id uint) error {
	media, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if media == nil {
		return ErrMediaNotFound
	}

	if strings.HasPrefix(media.FileURL, "/uploads/") {
		rel := strings.TrimPrefix(media.FileURL, "/uploads/")
		path := filepath.Join(s.uploadDir, filepath.FromSlash(rel))
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	return s.repo.Delete(id)
}

func (s *MediaService) CountAll() (int64, error) {
	return s.repo.CountAll()
}
