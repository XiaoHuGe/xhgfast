package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"xhgfast/utils/setting"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err := gorm.Open(setting.AppSetting.Database.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.AppSetting.Database.User,
		setting.AppSetting.Database.Password,
		setting.AppSetting.Database.Host,
		setting.AppSetting.Database.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.AppSetting.Database.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{})
	DB = db
	return db
}
