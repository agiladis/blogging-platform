package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"not null;unique;type:varchar(20)"`
	Email     string `json:"email" gorm:"not null;unique;type:varchar(40)"`
	Password  string `json:"password" gorm:"not null;type:varchar(80)"`
	BlogPosts []BlogPost
}
