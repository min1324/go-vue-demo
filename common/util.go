package common

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	GbConfig   = viper.New()
	UploadPath string
)

func InitConfig() {
	workDir, _ := os.Getwd()

	GbConfig.SetConfigName("app")
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

func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func MkDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
	}
	return err
}

func MkDirAll(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}
