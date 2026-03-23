package controller

import (
	"errors"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	service *service.MediaService
}

func NewMediaController(service *service.MediaService) *MediaController {
	return &MediaController{service: service}
}

// List 获取媒体列表
// @Summary 获取媒体列表
// @Tags Media
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/media [get]
func (ctl *MediaController) List(c *gin.Context) {
	var query dto.MediaListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	data, err := ctl.service.List(query)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取媒体列表失败")
		return
	}

	response.Success(c, data)
}

// Delete 删除媒体资源
// @Summary 删除媒体资源
// @Tags Media
// @Produce json
// @Security BearerAuth
// @Param id path int true "媒体ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/media/{id} [delete]
func (ctl *MediaController) Delete(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Delete(id)
	if err != nil {
		if errors.Is(err, service.ErrMediaNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "资源不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "删除媒体失败")
		return
	}

	response.Success(c, nil)
}
