package controller

import (
	"errors"
	"io"
	"sort"
	"strconv"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/model"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{service: service}
}

// ListPublic 获取文章列表
// @Summary 获取文章列表
// @Tags Post
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Param tag_id query int false "标签ID"
// @Param category_id query int false "分类ID"
// @Param sort query string false "排序"
// @Param is_top query bool false "是否置顶"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/posts [get]
func (ctl *PostController) ListPublic(c *gin.Context) {
	var query dto.PublicPostListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}
	query.Normalize()

	list, total, err := ctl.service.ListPublic(query)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取文章列表失败")
		return
	}

	response.Success(c, gin.H{
		"list":      list,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

// GetPublicBySlug 按 slug 获取文章详情
// @Summary 按 slug 获取文章详情
// @Tags Post
// @Produce json
// @Param slug path string true "slug"
// @Success 200 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/posts/slug/{slug} [get]
func (ctl *PostController) GetPublicBySlug(c *gin.Context) {
	slug := c.Param("slug")
	post, err := ctl.service.GetPublicBySlug(slug)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "获取文章详情失败")
		return
	}

	response.Success(c, post)
}

// GetPublicByID 按 ID 获取文章详情
// @Summary 按 ID 获取文章详情
// @Tags Post
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/posts/{id} [get]
func (ctl *PostController) GetPublicByID(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	post, err := ctl.service.GetPublicByID(id)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "获取文章详情失败")
		return
	}

	response.Success(c, post)
}

// Timeline 获取时间线
// @Summary 获取时间线
// @Tags Post
// @Produce json
// @Param year query int false "年份"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/timeline [get]
func (ctl *PostController) Timeline(c *gin.Context) {
	year, err := parseOptionalYear(c.Query("year"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	posts, err := ctl.service.Timeline(year)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取时间线失败")
		return
	}

	response.Success(c, buildTimeline(posts))
}

// ListAdmin 获取后台文章列表
// @Summary 获取后台文章列表
// @Tags AdminPost
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Param status query string false "状态"
// @Param category_id query int false "分类ID"
// @Param tag_id query int false "标签ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts [get]
func (ctl *PostController) ListAdmin(c *gin.Context) {
	var query dto.AdminPostListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}
	query.Normalize()

	list, total, err := ctl.service.ListAdmin(query)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取后台文章列表失败")
		return
	}

	response.Success(c, gin.H{
		"list":      list,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

// GetAdminByID 获取后台文章详情
// @Summary 获取后台文章详情
// @Tags AdminPost
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/{id} [get]
func (ctl *PostController) GetAdminByID(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	post, err := ctl.service.GetAdminByID(id)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "获取后台文章详情失败")
		return
	}

	response.Success(c, post)
}

// Create 新建文章
// @Summary 新建文章
// @Tags AdminPost
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body dto.SavePostRequest true "文章参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 409 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts [post]
func (ctl *PostController) Create(c *gin.Context) {
	var req dto.SavePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	id, err := ctl.service.Create(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrSlugConflict):
			response.Fail(c, 409, apperror.CodeConflict, "slug 已存在")
		case errors.Is(err, service.ErrInvalidStatus):
			response.Fail(c, 400, apperror.CodeBadRequest, "状态参数错误")
		default:
			response.Fail(c, 500, apperror.CodeInternal, "创建文章失败")
		}
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "创建成功", "data": gin.H{"id": id}})
}

// Update 更新文章
// @Summary 更新文章
// @Tags AdminPost
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Param payload body dto.SavePostRequest true "文章参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 409 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/{id} [put]
func (ctl *PostController) Update(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	var req dto.SavePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Update(id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrPostNotFound):
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
		case errors.Is(err, service.ErrSlugConflict):
			response.Fail(c, 409, apperror.CodeConflict, "slug 已存在")
		case errors.Is(err, service.ErrInvalidStatus):
			response.Fail(c, 400, apperror.CodeBadRequest, "状态参数错误")
		default:
			response.Fail(c, 500, apperror.CodeInternal, "更新文章失败")
		}
		return
	}

	response.Success(c, nil)
}

// Delete 删除文章
// @Summary 删除文章
// @Tags AdminPost
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/{id} [delete]
func (ctl *PostController) Delete(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Delete(id)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "删除文章失败")
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "删除成功", "data": nil})
}

// Publish 发布文章
// @Summary 发布文章
// @Tags AdminPost
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Param payload body dto.PublishPostRequest false "发布时间"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/{id}/publish [patch]
func (ctl *PostController) Publish(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	var req dto.PublishPostRequest
	if err := c.ShouldBindJSON(&req); err != nil && !errors.Is(err, io.EOF) {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Publish(id, req.PublishedAt)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "发布失败")
		return
	}

	response.Success(c, nil)
}

// Unpublish 下架文章
// @Summary 下架文章
// @Tags AdminPost
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/{id}/unpublish [patch]
func (ctl *PostController) Unpublish(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Unpublish(id)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "下架失败")
		return
	}

	c.JSON(200, gin.H{"code": 0, "message": "下架成功", "data": nil})
}

// SetTop 设置置顶状态
// @Summary 置顶或取消置顶
// @Tags AdminPost
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Param payload body dto.TopPostRequest true "置顶参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/{id}/top [patch]
func (ctl *PostController) SetTop(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	var req dto.TopPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.SetTop(id, req.IsTop)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "文章不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "操作失败")
		return
	}

	response.Success(c, nil)
}

// CheckSlug 检查 slug 是否可用
// @Summary 检查 slug 是否可用
// @Tags AdminPost
// @Produce json
// @Security BearerAuth
// @Param slug query string true "slug"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/posts/check-slug [get]
func (ctl *PostController) CheckSlug(c *gin.Context) {
	var query dto.CheckSlugQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	available, err := ctl.service.CheckSlugAvailable(query.Slug, 0)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "检查 slug 失败")
		return
	}

	response.Success(c, gin.H{"available": available})
}

func parseUintID(raw string) (uint, error) {
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func parseOptionalYear(raw string) (int, error) {
	if raw == "" {
		return 0, nil
	}
	return strconv.Atoi(raw)
}

func buildTimeline(posts []model.Post) []gin.H {
	type monthEntry struct {
		month int
		posts []gin.H
	}
	type yearEntry struct {
		year   int
		months map[int]*monthEntry
	}

	years := map[int]*yearEntry{}

	for _, post := range posts {
		if post.PublishedAt == nil {
			continue
		}
		t := post.PublishedAt.UTC()
		y := t.Year()
		m := int(t.Month())

		if _, exists := years[y]; !exists {
			years[y] = &yearEntry{year: y, months: map[int]*monthEntry{}}
		}
		if _, exists := years[y].months[m]; !exists {
			years[y].months[m] = &monthEntry{month: m}
		}

		years[y].months[m].posts = append(years[y].months[m].posts, gin.H{
			"id":           post.ID,
			"title":        post.Title,
			"slug":         post.Slug,
			"cover_image":  post.CoverImage,
			"published_at": post.PublishedAt,
		})
	}

	yearKeys := make([]int, 0, len(years))
	for y := range years {
		yearKeys = append(yearKeys, y)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(yearKeys)))

	result := make([]gin.H, 0, len(yearKeys))
	for _, y := range yearKeys {
		monthKeys := make([]int, 0, len(years[y].months))
		for m := range years[y].months {
			monthKeys = append(monthKeys, m)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(monthKeys)))

		months := make([]gin.H, 0, len(monthKeys))
		for _, m := range monthKeys {
			months = append(months, gin.H{
				"month": years[y].months[m].month,
				"posts": years[y].months[m].posts,
			})
		}

		result = append(result, gin.H{
			"year":   y,
			"months": months,
		})
	}

	return result
}
