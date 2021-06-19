package common

import (
	"demo/model"
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB global database
var DB *gorm.DB

// GetDB return global database
func GetDB() *gorm.DB {
	return DB
}

// InitDB initialize database
func InitDB() *gorm.DB {

	user := GbConfig.GetString("database.user")
	pass := GbConfig.GetString("database.pass")
	host := GbConfig.GetString("database.host")
	port := GbConfig.GetString("database.port")
	local := GbConfig.GetString("database.local")
	dbname := GbConfig.GetString("database.dbname")
	charset := GbConfig.GetString("database.charset")

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		user, pass, host, port, dbname, charset, url.QueryEscape(local),
	)

	// TODO connect to your db
	// hear you can change to your driver
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect db:" + err.Error())
	}

	// TODO add table.
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}
