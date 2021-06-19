package route

import (
	"demo/common"
	"demo/controller"
	"demo/frontend"
	"demo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	if common.GbConfig.GetBool("server.assets") {
		frontend.InitGinRouter(r)
	}
	r.Use(middleware.CorsMiddleWare(), middleware.RecoverMiddleware())

	userRoutes := r.Group("auth")
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)
	userRoutes.GET("/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	CategoryContoller := controller.NewCategoryController()
	categoryRoutes.POST("", CategoryContoller.Create)
	categoryRoutes.PUT("/:id", CategoryContoller.Update)
	categoryRoutes.DELETE("/:id", CategoryContoller.Delete)
	categoryRoutes.GET("/:id", CategoryContoller.Show)

	postRoutes := r.Group("/post")
	postRoutes.Use(middleware.AuthMiddleware())
	postContoller := controller.NewPostController()
	postRoutes.POST("", postContoller.Create)
	postRoutes.PUT("/:id", postContoller.Update)
	postRoutes.DELETE("/:id", postContoller.Delete)
	postRoutes.GET("/:id", postContoller.Show)
	postRoutes.GET("page/list", postContoller.PageList)

	fileRoutes := r.Group("/file")
	fileRoutes.POST("/upload", controller.Upload)

	return r
}
