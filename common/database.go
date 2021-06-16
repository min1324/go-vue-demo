package common

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"demo/model"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() *gorm.DB {

	user := viper.GetString("datasource.user")
	pass := viper.GetString("datasource.pass")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		user, pass, host, port, dbname, charset, url.QueryEscape(loc),
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect db:" + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}
