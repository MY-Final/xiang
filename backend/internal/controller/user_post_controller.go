package controller

import (
	"errors"
	"fmt"
	"net/http"
	"xiang/backend/internal/dto"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserPostController struct {
	postService *service.UserPostService
}

func NewUserPostController(postService *service.UserPostService) *UserPostController {
	return &UserPostController{postService: postService}
}

// ListMyPosts 获取我的日记列表
func (c *UserPostController) ListMyPosts(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	var query dto.PublicPostListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "请求参数错误")
		return
	}

	posts, total, err := c.postService.ListUserPosts(ctx, userID.(uint), query)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "获取日记列表失败")
		return
	}

	response.Success(ctx, gin.H{
		"list":  posts,
		"total": total,
		"page":  query.Page,
		"size":  query.PageSize,
	})
}

// GetMyPostByID 获取我的单篇日记
func (c *UserPostController) GetMyPostByID(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	id, err := parseID(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "无效的 ID")
		return
	}

	post, err := c.postService.GetUserPostByID(ctx, userID.(uint), id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "获取日记失败")
		return
	}
	if post == nil {
		response.Fail(ctx, http.StatusNotFound, apperror.CodeNotFound, "日记不存在")
		return
	}

	response.Success(ctx, post)
}

// CreateMyPost 创建日记
func (c *UserPostController) CreateMyPost(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	var req dto.SavePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "请求参数错误")
		return
	}

	id, err := c.postService.CreateUserPost(ctx, userID.(uint), req)
	if err != nil {
		if errors.Is(err, service.ErrSlugConflict) {
			response.Fail(ctx, http.StatusBadRequest, apperror.CodeConflict, "Slug 已存在")
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "创建日记失败")
		return
	}

	response.Success(ctx, gin.H{"id": id})
}

// UpdateMyPost 更新日记
func (c *UserPostController) UpdateMyPost(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	id, err := parseID(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "无效的 ID")
		return
	}

	var req dto.SavePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "请求参数错误")
		return
	}

	if err := c.postService.UpdateUserPost(ctx, userID.(uint), id, req); err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(ctx, http.StatusNotFound, apperror.CodeNotFound, "日记不存在")
			return
		}
		if errors.Is(err, service.ErrSlugConflict) {
			response.Fail(ctx, http.StatusBadRequest, apperror.CodeConflict, "Slug 已存在")
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "更新日记失败")
		return
	}

	response.Success(ctx, nil)
}

// DeleteMyPost 删除日记
func (c *UserPostController) DeleteMyPost(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	id, err := parseID(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "无效的 ID")
		return
	}

	if err := c.postService.DeleteUserPost(ctx, userID.(uint), id); err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			response.Fail(ctx, http.StatusNotFound, apperror.CodeNotFound, "日记不存在")
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "删除日记失败")
		return
	}

	response.Success(ctx, nil)
}

// GetMyStats 获取我的统计信息
func (c *UserPostController) GetMyStats(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	total, published, draft, err := c.postService.GetUserStats(ctx, userID.(uint))
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "获取统计信息失败")
		return
	}

	response.Success(ctx, gin.H{
		"total":     total,
		"published": published,
		"draft":     draft,
	})
}

func parseID(idStr string) (uint, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return 0, err
	}
	return id, nil
}
