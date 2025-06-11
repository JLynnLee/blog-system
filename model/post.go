package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `form:"title"`
	Content string `form:"content"`
	UserID  uint
	User    User
}
