package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"size:64;uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Nickname     string    `gorm:"size:128" json:"nickname"`
	Avatar       string    `gorm:"size:255" json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:128;not null" json:"name"`
	Slug        string         `gorm:"size:128;uniqueIndex;not null" json:"slug"`
	Description string         `gorm:"size:500" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Tag struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:128;not null" json:"name"`
	Slug      string         `gorm:"size:128;uniqueIndex;not null" json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Post struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"size:255;not null" json:"title"`
	Slug            string         `gorm:"size:255;uniqueIndex;not null" json:"slug"`
	Summary         string         `gorm:"size:500" json:"summary"`
	ContentMarkdown string         `gorm:"type:text" json:"content_markdown"`
	ContentHTML     string         `gorm:"type:text" json:"content_html"`
	CoverImage      string         `gorm:"size:255" json:"cover_image"`
	Status          string         `gorm:"size:32;index;default:draft" json:"status"`
	IsTop           bool           `gorm:"index;default:false" json:"is_top"`
	PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
	CategoryID      *uint          `gorm:"index" json:"category_id"`
	Category        *Category      `json:"category,omitempty"`
	Tags            []Tag          `gorm:"many2many:post_tags;" json:"tags,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type PostTag struct {
	PostID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
}

type Media struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FileName  string    `gorm:"size:255;not null" json:"file_name"`
	FileURL   string    `gorm:"size:500;not null" json:"file_url"`
	MimeType  string    `gorm:"size:100;not null" json:"mime_type"`
	FileSize  int64     `gorm:"not null" json:"file_size"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SiteSetting struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SiteName     string    `gorm:"size:255;not null;default:''" json:"site_name"`
	SiteSubtitle string    `gorm:"size:255;not null;default:''" json:"site_subtitle"`
	Logo         string    `gorm:"size:500;not null;default:''" json:"logo"`
	AboutText    string    `gorm:"type:text;not null;default:''" json:"about_text"`
	ThemeColor   string    `gorm:"size:32;not null;default:'#f9a8d4'" json:"theme_color"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Admin{}, &Category{}, &Tag{}, &Post{}, &PostTag{}, &Media{}, &SiteSetting{})
}
