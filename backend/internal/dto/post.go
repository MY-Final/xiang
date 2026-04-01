package dto

import "time"

type PublicPostListQuery struct {
	PaginationQuery
	Keyword    string `form:"keyword"`
	TagID      uint   `form:"tag_id"`
	CategoryID uint   `form:"category_id"`
	Sort       string `form:"sort"`
	IsTop      *bool  `form:"is_top"`
	Status     string `form:"status"`
}

type AdminPostListQuery struct {
	PaginationQuery
	Keyword    string `form:"keyword"`
	Status     string `form:"status"`
	CategoryID uint   `form:"category_id"`
	TagID      uint   `form:"tag_id"`
}

type SavePostRequest struct {
	Title           string     `json:"title" binding:"required"`
	Slug            string     `json:"slug" binding:"required"`
	Summary         string     `json:"summary"`
	ContentMarkdown string     `json:"content_markdown"`
	ContentHTML     string     `json:"content_html"`
	CoverImage      string     `json:"cover_image"`
	Status          string     `json:"status"`
	IsTop           bool       `json:"is_top"`
	CategoryID      *uint      `json:"category_id"`
	TagIDs          []uint     `json:"tag_ids"`
	PublishedAt     *time.Time `json:"published_at"`
}

type PublishPostRequest struct {
	PublishedAt *time.Time `json:"published_at"`
}

type TopPostRequest struct {
	IsTop bool `json:"is_top"`
}

type CheckSlugQuery struct {
	Slug string `form:"slug" binding:"required"`
}
