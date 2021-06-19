package main

import (
	"demo/common"
	"demo/route"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitConfig()
	common.InitDB()

	r := gin.Default()
	r = route.CollectRoute(r)

	ip := common.GbConfig.GetString("server.ip")
	if ip != "" {
		panic(r.Run(ip))
	}
	r.Run()
}
