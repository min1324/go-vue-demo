package route

import (
	"demo/controller"
	"demo/frontend"
	"demo/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter initizlize frontend routes.
func InitRouter(r *gin.Engine) *gin.Engine {

	// add frontend assets,router
	frontend.InitRouter(r)

	r.Use(middleware.RecoverMiddleware())
	r.Use(middleware.CorsMiddleWare())

	// add user router group
	userRoutes := r.Group("/auth")
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)
	userRoutes.GET("/info", middleware.AuthMiddleware(), controller.Info)

	// TODO add other router
	fileRoutes := r.Group("/file")

	// add authorize middleware to file upload
	fileRoutes.Use(middleware.AuthMiddleware())

	fileRoutes.POST("upload", controller.Upload)

	return r
}
