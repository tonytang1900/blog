package model

import (
	"blog/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDB()  {
	db,err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName))
	if err != nil {
		fmt.Printf("数据库连接错误,%v", err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&Article{}, &User{}, &Category{})

	db.DB().SetConnMaxLifetime(time.Second * 10)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//db.Close()
}