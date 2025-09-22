package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Code      string
	Name      string
	Posts     []Post
	PostCount int
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment
	CommentStatus string
}

type Comment struct {
	gorm.Model
	Comment string
	PostID  uint
}

func Model(db *gorm.DB) {
	db.AutoMigrate(User{}, Post{}, Comment{})

	// user := User{
	// 	Code: "02",
	// 	Name: "李四",
	// 	Posts: []Post{
	// 		{
	// 			Title:   "文章3",
	// 			Content: "3333",
	// 			Comments: []Comment{
	// 				{
	// 					Comment: "哈哈3",
	// 				},
	// 				{
	// 					Comment: "哈哈哈3",
	// 				},
	// 				{
	// 					Comment: "哈3",
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// db.Create(&user)
}
