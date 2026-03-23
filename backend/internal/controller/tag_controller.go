package controller

import (
	"errors"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	service *service.TagService
}

func NewTagController(service *service.TagService) *TagController {
	return &TagController{service: service}
}

// List 获取标签列表
// @Summary 获取标签列表
// @Tags Tag
// @Produce json
// @Success 200 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/tags [get]
// @Router /api/v1/admin/tags [get]
func (ctl *TagController) List(c *gin.Context) {
	tags, err := ctl.service.List()
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取标签列表失败")
		return
	}
	response.Success(c, tags)
}

// Create 创建标签
// @Summary 创建标签
// @Tags Tag
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body dto.SaveTagRequest true "标签参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/tags [post]
func (ctl *TagController) Create(c *gin.Context) {
	var req dto.SaveTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	tag, err := ctl.service.Create(req)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "创建标签失败")
		return
	}
	response.Success(c, tag)
}

// Update 更新标签
// @Summary 更新标签
// @Tags Tag
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "标签ID"
// @Param payload body dto.SaveTagRequest true "标签参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/tags/{id} [put]
func (ctl *TagController) Update(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	var req dto.SaveTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Update(id, req)
	if err != nil {
		if errors.Is(err, service.ErrTagNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "标签不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "更新标签失败")
		return
	}

	response.Success(c, nil)
}

// Delete 删除标签
// @Summary 删除标签
// @Tags Tag
// @Produce json
// @Security BearerAuth
// @Param id path int true "标签ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/tags/{id} [delete]
func (ctl *TagController) Delete(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Delete(id)
	if err != nil {
		if errors.Is(err, service.ErrTagNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "标签不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "删除标签失败")
		return
	}

	response.Success(c, nil)
}
