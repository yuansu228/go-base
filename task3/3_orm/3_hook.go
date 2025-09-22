package main

import "gorm.io/gorm"

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var postCount int
	tx.Debug().Model(&Post{}).Select("count(1) as post_count").Where("user_id = ?", p.UserID).Scan(&postCount)
	tx.Debug().Model(&User{}).Where("id = ?", p.UserID).Update("post_count", postCount)
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	commentCount := 0
	tx.Debug().Model(&Comment{}).Select("count(1) as comment_count").Where("post_id = ?", c.PostID).Scan(&commentCount)
	if commentCount == 0 {
		tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
	}
	return nil
}
