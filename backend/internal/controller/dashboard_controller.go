package controller

import (
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	service *service.DashboardService
}

func NewDashboardController(service *service.DashboardService) *DashboardController {
	return &DashboardController{service: service}
}

// Stats 获取仪表盘统计
// @Summary 获取仪表盘统计
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/dashboard [get]
func (ctl *DashboardController) Stats(c *gin.Context) {
	data, err := ctl.service.Stats()
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取仪表盘统计失败")
		return
	}
	response.Success(c, data)
}
