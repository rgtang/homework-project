package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	User    *User  `gorm:"foreignKey:UserID" json:"author,omitempty"`
	PostID  uint   `json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostID" json:"-"`
}
