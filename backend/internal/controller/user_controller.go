package controller

import (
	"errors"
	"net/http"
	"xiang/backend/internal/dto"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "注册信息"
// @Success 200 {object} dto.UserProfile
// @Router /api/v1/user/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "请求参数错误")
		return
	}

	profile, err := c.userService.Register(ctx, req)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			response.Fail(ctx, http.StatusBadRequest, apperror.CodeConflict, "用户名已存在")
			return
		}
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			response.Fail(ctx, http.StatusBadRequest, apperror.CodeConflict, "邮箱已被使用")
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "注册失败")
		return
	}

	response.Success(ctx, profile)
}

// Login 用户登录
// @Summary 用户登录
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.UserLoginRequest true "登录信息"
// @Success 200 {object} dto.UserLoginResponse
// @Router /api/v1/user/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "请求参数错误")
		return
	}

	resp, err := c.userService.Login(ctx, req)
	if err != nil {
		if errors.Is(err, service.ErrUserInvalidCredential) {
			response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "用户名或密码错误")
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "登录失败")
		return
	}

	response.Success(ctx, resp)
}

// GetMe 获取当前用户信息
// @Summary 获取当前用户信息
// @Tags User
// @Produce json
// @Success 200 {object} dto.UserProfile
// @Router /api/v1/user/me [get]
func (c *UserController) GetMe(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	profile, err := c.userService.GetProfile(ctx, userID.(uint))
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "获取用户信息失败")
		return
	}
	if profile == nil {
		response.Fail(ctx, http.StatusNotFound, apperror.CodeNotFound, "用户不存在")
		return
	}

	response.Success(ctx, profile)
}

// UpdateProfile 更新用户资料
// @Summary 更新用户资料
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "用户资料"
// @Success 200 {object} dto.UserProfile
// @Router /api/v1/user/profile [put]
func (c *UserController) UpdateProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, apperror.CodeUnauthorized, "未登录")
		return
	}

	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, apperror.CodeBadRequest, "请求参数错误")
		return
	}

	profile, err := c.userService.UpdateProfile(ctx, userID.(uint), req)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			response.Fail(ctx, http.StatusBadRequest, apperror.CodeConflict, "邮箱已被使用")
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, apperror.CodeInternal, "更新失败")
		return
	}

	response.Success(ctx, profile)
}
