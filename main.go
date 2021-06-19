package main

import (
	"demo/common"
	"demo/middleware"
	"demo/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// Frist init config,then use config init db.
	// InitConfig must before InitDB.
	common.InitConfig()
	common.InitDB()

	// init router
	r := gin.Default()
	r.Use(middleware.RecoverMiddleware())

	route.InitRouter(r)

	addr := common.GbConfig.GetString("server.dns")
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}
