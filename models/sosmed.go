package models

import "gorm.io/gorm"

type Sosmed struct {
	gorm.Model
	// ID        int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Name      string `gorm:"not null" json:"name" form:"name"`
	SosmedURL string `gorm:"not null" json:"sosmed_url" form:"sosmed_url"`
	// UserID    User   `json:"user_id" form:"user_id"`
	// User      User   `gorm:"foreignKey:ID"`
	UserID uint `json:"user_id"`
	User   *User
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
