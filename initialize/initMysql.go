package initialize

import (
	"gin-admin-back/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DEFAULTDB *gorm.DB

func InitMysql(admin config.Admin) {
	if db, err := gorm.Open("mysql", admin.UserName+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		log.Printf("DEFAULTDB err: %S", err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(10)
		DEFAULTDB.DB().SetMaxIdleConns(100)
		DEFAULTDB.AutoMigrate()
	}
}
