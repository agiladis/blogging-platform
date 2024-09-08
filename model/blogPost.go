package model

import (
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title string `json:"title" gorm:"not null;type:varchar(30)"`
	Content string `json:"content" gorm:"not null;type:text"`
	UserID uint `json:"user_id" gorm:"not null"`
}