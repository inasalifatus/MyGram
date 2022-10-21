package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID      int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	UserID  User   `json:"user_id" form:"user_id"`
	PhotoID Photos `json:"photo_id" form:"photo_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Photos  Photos `gorm:"foreignKey:PhotoID"`
	Message string `json:"message" form:"message"`
}

type Comment_post struct {
	Message string `json:"message"`
	PhotoID Photos `json:"photo_id"`
}

type Comment_response_single struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    Comment `json:"data"`
}

type Comment_Resp struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Comment `json:"data"`
}

type Comment_update struct {
	Message string `json:"message"`
}
