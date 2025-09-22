package main

import (
	"fmt"

	"gorm.io/gorm"
)

func QueryAll(db *gorm.DB) {
	var user User
	db.Debug().Preload("Posts.Comments").First(&user, 1)
	fmt.Println(user)
}

func QueryTopCommentPost(db *gorm.DB) {
	var result struct {
		PostID       uint
		CommentCount int
	}
	db.Model(&Comment{}).Select("post_id,count(1) as comment_count").Group("post_id").Order("comment_count desc").Scan(&result)

	post := Post{}
	db.First(&post, result.PostID)
	fmt.Println("评论最多的文章信息：", post)
}
