package route

import (
	"demo/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter initizlize frontend routes.
func InitRouter(r *gin.Engine) *gin.Engine {

	userRoutes := r.Group("/auth")
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)

	return r
}
