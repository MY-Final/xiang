package repository

import (
	"errors"

	"xiang/backend/internal/model"

	"gorm.io/gorm"
)

type SiteSettingRepository struct {
	db *gorm.DB
}

func NewSiteSettingRepository(db *gorm.DB) *SiteSettingRepository {
	return &SiteSettingRepository{db: db}
}

func (r *SiteSettingRepository) Get() (*model.SiteSetting, error) {
	var setting model.SiteSetting
	err := r.db.First(&setting).Error
	if err == nil {
		return &setting, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	setting = model.SiteSetting{
		SiteName:     "恋爱小屋",
		SiteSubtitle: "记录我们的日常与心动",
		ThemeColor:   "#f9a8d4",
	}
	if err := r.db.Create(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r *SiteSettingRepository) Update(data *model.SiteSetting) error {
	setting, err := r.Get()
	if err != nil {
		return err
	}

	return r.db.Model(&model.SiteSetting{}).Where("id = ?", setting.ID).Updates(map[string]interface{}{
		"site_name":     data.SiteName,
		"site_subtitle": data.SiteSubtitle,
		"logo":          data.Logo,
		"about_text":    data.AboutText,
		"theme_color":   data.ThemeColor,
	}).Error
}
