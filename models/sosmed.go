package models

import "gorm.io/gorm"

type Sosmed struct {
	gorm.Model
	ID        int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	UserID    User   `json:"user_id" form:"user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	Name      string `json:"name" form:"name"`
	SosmedURL string `json:"sosmed_url" form:"sosmed_url"`
}

type Sosmed_post struct {
	Name      string `json:"name" validate:"required"`
	SosmedURL string `json:"sosmed_url" validate:"required"`
}

type Sosmed_response_single struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Sosmed `json:"data"`
}

type Sosmed_Resp struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []Sosmed `json:"data"`
}
