package common

import (
	"demo/model"
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() *gorm.DB {

	user := GbConfig.GetString("datasource.user")
	pass := GbConfig.GetString("datasource.pass")
	host := GbConfig.GetString("datasource.host")
	port := GbConfig.GetString("datasource.port")
	dbname := GbConfig.GetString("datasource.dbname")
	charset := GbConfig.GetString("datasource.charset")
	loc := GbConfig.GetString("datasource.loc")

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
