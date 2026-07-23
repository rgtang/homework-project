package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // 不在 JSON 中返回密码
	Email    string `gorm:"unique;not null" json:"email"`
}
