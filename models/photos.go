package models

import "gorm.io/gorm"

type Photos struct {
	gorm.Model
	ID       int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Title    string `json:"title" form:"title" valid:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
	// UserID   User   `json:"user_id" form:"user_id"`
	User []User `json:"user" gorm:"foreignKey:UserID"`
}

type Photos_Resp struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []Photos `json:"data"`
}

type Photos_response_single struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Photos `json:"data"`
}

type Photos_post struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}

type Photos_update struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}
