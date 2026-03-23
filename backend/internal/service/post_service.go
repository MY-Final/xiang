package service

import (
	"errors"
	"strings"
	"time"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	"xiang/backend/internal/repository"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrSlugConflict  = errors.New("slug already exists")
	ErrInvalidStatus = errors.New("invalid status")
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) ListPublic(query dto.PublicPostListQuery) ([]model.Post, int64, error) {
	query.Normalize()
	return s.postRepo.ListPublic(query)
}

func (s *PostService) GetPublicBySlug(slug string) (*model.Post, error) {
	post, err := s.postRepo.GetPublicBySlug(slug)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}
	return post, nil
}

func (s *PostService) GetPublicByID(id uint) (*model.Post, error) {
	post, err := s.postRepo.GetPublicByID(id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}
	return post, nil
}

func (s *PostService) ListAdmin(query dto.AdminPostListQuery) ([]model.Post, int64, error) {
	query.Normalize()
	return s.postRepo.ListAdmin(query)
}

func (s *PostService) GetAdminByID(id uint) (*model.Post, error) {
	post, err := s.postRepo.GetAdminByID(id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}
	return post, nil
}

func (s *PostService) Create(req dto.SavePostRequest) (uint, error) {
	if err := validateStatus(req.Status); err != nil {
		return 0, err
	}

	exists, err := s.postRepo.SlugExists(req.Slug, 0)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, ErrSlugConflict
	}

	post := &model.Post{
		Title:           req.Title,
		Slug:            strings.TrimSpace(req.Slug),
		Summary:         req.Summary,
		ContentMarkdown: req.ContentMarkdown,
		ContentHTML:     req.ContentHTML,
		CoverImage:      req.CoverImage,
		Status:          normalizeStatus(req.Status),
		IsTop:           req.IsTop,
		PublishedAt:     req.PublishedAt,
		CategoryID:      req.CategoryID,
	}

	if err := s.postRepo.Create(post); err != nil {
		return 0, err
	}

	if err := s.postRepo.SaveTags(post, req.TagIDs); err != nil {
		return 0, err
	}

	return post.ID, nil
}

func (s *PostService) Update(id uint, req dto.SavePostRequest) error {
	if err := validateStatus(req.Status); err != nil {
		return err
	}

	post, err := s.postRepo.GetAdminByID(id)
	if err != nil {
		return err
	}
	if post == nil {
		return ErrPostNotFound
	}

	exists, err := s.postRepo.SlugExists(req.Slug, id)
	if err != nil {
		return err
	}
	if exists {
		return ErrSlugConflict
	}

	post.Title = req.Title
	post.Slug = strings.TrimSpace(req.Slug)
	post.Summary = req.Summary
	post.ContentMarkdown = req.ContentMarkdown
	post.ContentHTML = req.ContentHTML
	post.CoverImage = req.CoverImage
	post.Status = normalizeStatus(req.Status)
	post.IsTop = req.IsTop
	post.PublishedAt = req.PublishedAt
	post.CategoryID = req.CategoryID

	if err := s.postRepo.Update(post); err != nil {
		return err
	}

	return s.postRepo.SaveTags(post, req.TagIDs)
}

func (s *PostService) Delete(id uint) error {
	post, err := s.postRepo.GetAdminByID(id)
	if err != nil {
		return err
	}
	if post == nil {
		return ErrPostNotFound
	}

	return s.postRepo.Delete(id)
}

func (s *PostService) Publish(id uint, publishedAt *time.Time) error {
	post, err := s.postRepo.GetAdminByID(id)
	if err != nil {
		return err
	}
	if post == nil {
		return ErrPostNotFound
	}
	return s.postRepo.Publish(id, publishedAt)
}

func (s *PostService) Unpublish(id uint) error {
	post, err := s.postRepo.GetAdminByID(id)
	if err != nil {
		return err
	}
	if post == nil {
		return ErrPostNotFound
	}
	return s.postRepo.Unpublish(id)
}

func (s *PostService) SetTop(id uint, isTop bool) error {
	post, err := s.postRepo.GetAdminByID(id)
	if err != nil {
		return err
	}
	if post == nil {
		return ErrPostNotFound
	}
	return s.postRepo.SetTop(id, isTop)
}

func (s *PostService) Timeline(year int) ([]model.Post, error) {
	return s.postRepo.Timeline(year)
}

func (s *PostService) CheckSlugAvailable(slug string, excludeID uint) (bool, error) {
	exists, err := s.postRepo.SlugExists(strings.TrimSpace(slug), excludeID)
	if err != nil {
		return false, err
	}
	return !exists, nil
}

func validateStatus(status string) error {
	normalized := normalizeStatus(status)
	switch normalized {
	case "draft", "published", "private":
		return nil
	default:
		return ErrInvalidStatus
	}
}

func normalizeStatus(status string) string {
	trimmed := strings.TrimSpace(status)
	if trimmed == "" {
		return "draft"
	}
	return trimmed
}
