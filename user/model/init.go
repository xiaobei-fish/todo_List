package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connStr string) {
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	// 开启数据库日志
	db.LogMode(true)
	// 如果是上线模式，关闭日志避免logger等不同造成的异常
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)

	//配置连接池
	db.DB().SetMaxIdleConns(20) // 设置最大空闲连接数

	db.DB().SetMaxOpenConns(100) // 设置最大连接数

	DB = db

	migration()
}
