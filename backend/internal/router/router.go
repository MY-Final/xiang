package router

import (
	"net/http"

	"xiang/backend/internal/config"
	"xiang/backend/internal/controller"
	"xiang/backend/internal/middleware"
	"xiang/backend/internal/repository"
	"xiang/backend/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.Static("/uploads", cfg.UploadDir)
	r.StaticFile("/swagger/doc.json", "./docs/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	adminRepo := repository.NewAdminRepository(db)
	postRepo := repository.NewPostRepository(db)
	tagRepo := repository.NewTagRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	mediaRepo := repository.NewMediaRepository(db)
	siteRepo := repository.NewSiteSettingRepository(db)

	authService := service.NewAuthService(adminRepo, cfg)
	postService := service.NewPostService(postRepo)
	tagService := service.NewTagService(tagRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	uploadService := service.NewUploadService(cfg.UploadDir, mediaRepo)
	mediaService := service.NewMediaService(mediaRepo, cfg.UploadDir)
	siteService := service.NewSiteSettingService(siteRepo)
	dashboardService := service.NewDashboardService(postRepo, tagRepo, categoryRepo, mediaRepo)

	authCtl := controller.NewAuthController(authService)
	postCtl := controller.NewPostController(postService)
	tagCtl := controller.NewTagController(tagService)
	categoryCtl := controller.NewCategoryController(categoryService)
	uploadCtl := controller.NewUploadController(uploadService)
	mediaCtl := controller.NewMediaController(mediaService)
	siteCtl := controller.NewSiteSettingController(siteService)
	dashboardCtl := controller.NewDashboardController(dashboardService)

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authCtl.Login)

			authProtected := auth.Group("")
			authProtected.Use(middleware.JWT(cfg.JWTSecret))
			{
				authProtected.GET("/me", authCtl.Me)
				authProtected.POST("/logout", authCtl.Logout)
			}
		}

		api.GET("/posts", postCtl.ListPublic)
		api.GET("/posts/:id", postCtl.GetPublicByID)
		api.GET("/posts/slug/:slug", postCtl.GetPublicBySlug)
		api.GET("/tags", tagCtl.List)
		api.GET("/categories", categoryCtl.List)
		api.GET("/timeline", postCtl.Timeline)
		api.GET("/site/settings", siteCtl.Get)

		admin := api.Group("/admin")
		admin.Use(middleware.JWT(cfg.JWTSecret))
		{
			admin.GET("/posts", postCtl.ListAdmin)
			admin.GET("/posts/:id", postCtl.GetAdminByID)
			admin.POST("/posts", postCtl.Create)
			admin.PUT("/posts/:id", postCtl.Update)
			admin.DELETE("/posts/:id", postCtl.Delete)
			admin.PATCH("/posts/:id/publish", postCtl.Publish)
			admin.PATCH("/posts/:id/unpublish", postCtl.Unpublish)
			admin.PATCH("/posts/:id/top", postCtl.SetTop)
			admin.GET("/posts/check-slug", postCtl.CheckSlug)

			admin.GET("/tags", tagCtl.List)
			admin.POST("/tags", tagCtl.Create)
			admin.PUT("/tags/:id", tagCtl.Update)
			admin.DELETE("/tags/:id", tagCtl.Delete)

			admin.GET("/categories", categoryCtl.List)
			admin.POST("/categories", categoryCtl.Create)
			admin.PUT("/categories/:id", categoryCtl.Update)
			admin.DELETE("/categories/:id", categoryCtl.Delete)

			admin.POST("/upload/image", uploadCtl.UploadImage)
			admin.POST("/upload/images", uploadCtl.UploadImages)

			admin.GET("/media", mediaCtl.List)
			admin.DELETE("/media/:id", mediaCtl.Delete)

			admin.GET("/site/settings", siteCtl.Get)
			admin.PUT("/site/settings", siteCtl.Update)

			admin.GET("/dashboard", dashboardCtl.Stats)
		}
	}

	return r
}
