package common

import (
	"os"

	"github.com/spf13/viper"
)

var (
	// GbConfig global config
	GbConfig = viper.New()

	UploadPath string
)

// InitConfig initialize global config
func InitConfig() {
	workDir, _ := os.Getwd()

	GbConfig.SetConfigName("app")
	GbConfig.SetConfigType("yaml")
	GbConfig.AddConfigPath(workDir + "/config")

	err := GbConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dir := GbConfig.GetString("server.dir")
	if err = MkDirAll(dir); err != nil {
		panic(err)
	}
	UploadPath = dir
}
