package controller

import (
	"xiang/backend/internal/dto"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type SiteSettingController struct {
	service *service.SiteSettingService
}

func NewSiteSettingController(service *service.SiteSettingService) *SiteSettingController {
	return &SiteSettingController{service: service}
}

// Get 获取站点设置
// @Summary 获取站点设置
// @Tags Site
// @Produce json
// @Success 200 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/site/settings [get]
// @Router /api/v1/admin/site/settings [get]
func (ctl *SiteSettingController) Get(c *gin.Context) {
	data, err := ctl.service.Get()
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取站点设置失败")
		return
	}
	response.Success(c, data)
}

// Update 更新站点设置
// @Summary 更新站点设置
// @Tags Site
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body dto.SaveSiteSettingRequest true "站点设置"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/site/settings [put]
func (ctl *SiteSettingController) Update(c *gin.Context) {
	var req dto.SaveSiteSettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	if err := ctl.service.Update(req); err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "更新站点设置失败")
		return
	}

	response.Success(c, nil)
}
