package controllers

import (
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db = newDB()
)

func newDB() *gorm.DB {
	user := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	scheme := beego.AppConfig.String("mysqldb")
	sqlDebug := beego.AppConfig.String("mysqldebug")
	Logger := logger.Default
	if sqlDebug == "true" {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(user+":"+password+"@tcp("+url+")/"+scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	return db
}
