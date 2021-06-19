package common

import (
	"os"

	"github.com/spf13/viper"
)

var GbConfig = viper.New()

func InitConfig() {
	workDir, _ := os.Getwd()

	GbConfig.SetConfigName("app")
	GbConfig.SetConfigType("yaml")
	GbConfig.AddConfigPath(workDir + "/config")

	err := GbConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
