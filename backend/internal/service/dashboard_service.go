package service

import "xiang/backend/internal/repository"

type DashboardService struct {
	postRepo     *repository.PostRepository
	tagRepo      *repository.TagRepository
	categoryRepo *repository.CategoryRepository
	mediaRepo    *repository.MediaRepository
}

func NewDashboardService(
	postRepo *repository.PostRepository,
	tagRepo *repository.TagRepository,
	categoryRepo *repository.CategoryRepository,
	mediaRepo *repository.MediaRepository,
) *DashboardService {
	return &DashboardService{
		postRepo:     postRepo,
		tagRepo:      tagRepo,
		categoryRepo: categoryRepo,
		mediaRepo:    mediaRepo,
	}
}

func (s *DashboardService) Stats() (map[string]int64, error) {
	postCount, err := s.postRepo.CountAll()
	if err != nil {
		return nil, err
	}
	publishedCount, err := s.postRepo.CountByStatus("published")
	if err != nil {
		return nil, err
	}
	draftCount, err := s.postRepo.CountByStatus("draft")
	if err != nil {
		return nil, err
	}
	tagCount, err := s.tagRepo.CountAll()
	if err != nil {
		return nil, err
	}
	categoryCount, err := s.categoryRepo.CountAll()
	if err != nil {
		return nil, err
	}
	mediaCount, err := s.mediaRepo.CountAll()
	if err != nil {
		return nil, err
	}

	return map[string]int64{
		"post_count":      postCount,
		"published_count": publishedCount,
		"draft_count":     draftCount,
		"tag_count":       tagCount,
		"category_count":  categoryCount,
		"media_count":     mediaCount,
	}, nil
}
