package service

import (
	"strings"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	"xiang/backend/internal/repository"
)

type SiteSettingService struct {
	repo *repository.SiteSettingRepository
}

func NewSiteSettingService(repo *repository.SiteSettingRepository) *SiteSettingService {
	return &SiteSettingService{repo: repo}
}

func (s *SiteSettingService) Get() (*model.SiteSetting, error) {
	return s.repo.Get()
}

func (s *SiteSettingService) Update(req dto.SaveSiteSettingRequest) error {
	data := &model.SiteSetting{
		SiteName:     strings.TrimSpace(req.SiteName),
		SiteSubtitle: strings.TrimSpace(req.SiteSubtitle),
		Logo:         strings.TrimSpace(req.Logo),
		AboutText:    strings.TrimSpace(req.AboutText),
		ThemeColor:   strings.TrimSpace(req.ThemeColor),
	}
	if data.ThemeColor == "" {
		data.ThemeColor = "#f9a8d4"
	}
	return s.repo.Update(data)
}
