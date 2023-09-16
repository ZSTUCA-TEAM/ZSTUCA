package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
	"zstuCA/server"
	"zstuCA/server/database"
)

func main() {
	// 数据库连接与初始化
	db, err := gorm.Open("sqlite3", "D:/Go/code/zstuCA/database.sqlite")
	//db, err := gorm.Open("mysql", "zstuCA:3pM8DiWZNWRyswcy@(127.0.0.1:3306)/zstuca?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Failed to close database:", err)
		}
	}(db)
	database.DbInit(db)

	// iris服务器框架初始化与路由绑定
	app := iris.Default()
	server.Handle(app)

	err = app.Listen(":80")
	if err != nil {
		println("Failed to start server:", err)
	}
}
