package main

import (
	"demo/common"
	"demo/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDB()

	r := gin.Default()
	r = route.CollectRoute(r)

	//r.Run("192.168.188.220:8080") // 监听并在 0.0.0.0:8080 上启动服务

	ip := viper.GetString("server.ip")
	if ip != "" {
		r.Run(ip)
	}
	r.Run()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
