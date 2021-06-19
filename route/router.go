package route

import (
	"demo/controller"
	"demo/frontend"

	"github.com/gin-gonic/gin"
)

// InitRouter initizlize frontend routes.
func InitRouter(r *gin.Engine) *gin.Engine {

	// add frontend assets,router
	frontend.InitRouter(r)

	// add user router group
	userRoutes := r.Group("/auth")
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)

	// TODO add other router

	return r
}
