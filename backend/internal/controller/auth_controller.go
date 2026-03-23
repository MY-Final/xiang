package controller

import (
	"errors"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/middleware"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{service: service}
}

// Login 管理员登录
// @Summary 管理员登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body dto.LoginRequest true "登录参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/auth/login [post]
func (ctl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	data, err := ctl.service.Login(req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			response.Fail(c, 401, apperror.CodeUnauthorized, "账号或密码错误")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "登录失败")
		return
	}

	response.Success(c, data)
}

// Me 获取当前管理员信息
// @Summary 获取当前管理员信息
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/auth/me [get]
func (ctl *AuthController) Me(c *gin.Context) {
	adminID, ok := c.Get(middleware.CtxAdminIDKey)
	if !ok {
		response.Fail(c, 401, apperror.CodeUnauthorized, "未授权")
		return
	}

	profile, err := ctl.service.Me(adminID.(uint))
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取用户信息失败")
		return
	}
	if profile == nil {
		response.Fail(c, 404, apperror.CodeNotFound, "用户不存在")
		return
	}

	response.Success(c, profile)
}

// Logout 退出登录
// @Summary 退出登录
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Body
// @Failure 401 {object} response.Body
// @Router /api/v1/auth/logout [post]
func (ctl *AuthController) Logout(c *gin.Context) {
	response.Success(c, nil)
}
