package service

import (
	"context"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	"xiang/backend/internal/repository"
)

type UserPostService struct {
	postRepo *repository.PostRepository
}

func NewUserPostService(postRepo *repository.PostRepository) *UserPostService {
	return &UserPostService{postRepo: postRepo}
}

// ListUserPosts 获取用户的日记列表
func (s *UserPostService) ListUserPosts(ctx context.Context, userID uint, query dto.PublicPostListQuery) ([]model.Post, int64, error) {
	query.Normalize()
	return s.postRepo.ListUserPosts(ctx, userID, query)
}

// GetUserPostByID 获取用户的单篇日记
func (s *UserPostService) GetUserPostByID(ctx context.Context, userID, postID uint) (*model.Post, error) {
	post, err := s.postRepo.GetAdminByID(postID)
	if err != nil {
		return nil, err
	}
	if post == nil || post.UserID != userID {
		return nil, nil
	}
	return post, nil
}

// CreateUserPost 创建用户的日记
func (s *UserPostService) CreateUserPost(ctx context.Context, userID uint, req dto.SavePostRequest) (uint, error) {
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
		UserID:          userID,
		Title:           req.Title,
		Slug:            req.Slug,
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

// UpdateUserPost 更新用户的日记
func (s *UserPostService) UpdateUserPost(ctx context.Context, userID, postID uint, req dto.SavePostRequest) error {
	post, err := s.postRepo.GetAdminByID(postID)
	if err != nil {
		return err
	}
	if post == nil || post.UserID != userID {
		return ErrPostNotFound
	}

	exists, err := s.postRepo.SlugExists(req.Slug, postID)
	if err != nil {
		return err
	}
	if exists {
		return ErrSlugConflict
	}

	post.Title = req.Title
	post.Slug = req.Slug
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

// DeleteUserPost 删除用户的日记
func (s *UserPostService) DeleteUserPost(ctx context.Context, userID, postID uint) error {
	post, err := s.postRepo.GetAdminByID(postID)
	if err != nil {
		return err
	}
	if post == nil || post.UserID != userID {
		return ErrPostNotFound
	}

	return s.postRepo.Delete(postID)
}

// GetUserStats 获取用户统计信息
func (s *UserPostService) GetUserStats(ctx context.Context, userID uint) (total int64, published int64, draft int64, err error) {
	total, err = s.postRepo.CountByUser(userID)
	if err != nil {
		return
	}
	published, err = s.postRepo.CountByUserAndStatus(userID, "published")
	if err != nil {
		return
	}
	draft, err = s.postRepo.CountByUserAndStatus(userID, "draft")
	if err != nil {
		return
	}
	return
}
