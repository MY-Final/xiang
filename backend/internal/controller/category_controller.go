package controller

import (
	"errors"

	"xiang/backend/internal/dto"
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// List 获取分类列表
// @Summary 获取分类列表
// @Tags Category
// @Produce json
// @Success 200 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/categories [get]
// @Router /api/v1/admin/categories [get]
func (ctl *CategoryController) List(c *gin.Context) {
	categories, err := ctl.service.List()
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "获取分类列表失败")
		return
	}
	response.Success(c, categories)
}

// Create 创建分类
// @Summary 创建分类
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body dto.SaveCategoryRequest true "分类参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/categories [post]
func (ctl *CategoryController) Create(c *gin.Context) {
	var req dto.SaveCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	category, err := ctl.service.Create(req)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "创建分类失败")
		return
	}
	response.Success(c, category)
}

// Update 更新分类
// @Summary 更新分类
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "分类ID"
// @Param payload body dto.SaveCategoryRequest true "分类参数"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/categories/{id} [put]
func (ctl *CategoryController) Update(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	var req dto.SaveCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Update(id, req)
	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "分类不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "更新分类失败")
		return
	}

	response.Success(c, nil)
}

// Delete 删除分类
// @Summary 删除分类
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Param id path int true "分类ID"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 404 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/categories/{id} [delete]
func (ctl *CategoryController) Delete(c *gin.Context) {
	id, err := parseUintID(c.Param("id"))
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	err = ctl.service.Delete(id)
	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			response.Fail(c, 404, apperror.CodeNotFound, "分类不存在")
			return
		}
		response.Fail(c, 500, apperror.CodeInternal, "删除分类失败")
		return
	}

	response.Success(c, nil)
}
