package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:secret@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	// Model(db)
	// QueryAll(db)
	// QueryTopCommentPost(db)
	// db.Create(&Post{Title: "空文章", Content: "", UserID: 1})
	comments := []Comment{}
	db.Find(&comments)
	// 1、必须是物理删除才会触发钩子
	// 2、批量删除如果传入的是空属性，则钩子无法接收到删除的参数
	db.Unscoped().Delete(&comments)
}
