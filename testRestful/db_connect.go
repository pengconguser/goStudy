package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	//打开数据库连接

	var err error

	db, err = gorm.Open("mysql", sqlConnection)

	if err != nil {
		panic("数据库连接失败")
	}

	db.AutoMigrate(&todoModel{})
}
