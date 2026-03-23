package controller

import (
	"xiang/backend/internal/pkg/apperror"
	"xiang/backend/internal/pkg/response"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	service *service.UploadService
}

func NewUploadController(service *service.UploadService) *UploadController {
	return &UploadController{service: service}
}

// UploadImage 上传单图
// @Summary 上传单图
// @Tags Upload
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param file formData file true "图片文件"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/upload/image [post]
func (ctl *UploadController) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "请上传图片文件")
		return
	}

	media, err := ctl.service.SaveImage(file)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "上传失败")
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "上传成功",
		"data":    media,
	})
}

// UploadImages 上传多图
// @Summary 上传多图
// @Tags Upload
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param files formData file true "图片文件(可多选)"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Failure 401 {object} response.Body
// @Failure 500 {object} response.Body
// @Router /api/v1/admin/upload/images [post]
func (ctl *UploadController) UploadImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.Fail(c, 400, apperror.CodeBadRequest, "参数错误")
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		response.Fail(c, 400, apperror.CodeBadRequest, "请上传图片文件")
		return
	}

	items, err := ctl.service.SaveImages(files)
	if err != nil {
		response.Fail(c, 500, apperror.CodeInternal, "上传失败")
		return
	}

	response.Success(c, items)
}
