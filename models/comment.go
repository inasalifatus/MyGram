package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	// ID      int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Message string `gorm:"not null" json:"message" form:"message"`
	UserID  uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	User    *User
	PhotoID uint `json:"photo_id"`
	// Photo   *Photos

	// Photo   *Photos
	// PhotoID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos_id"`
	// User    User   `gorm:"foreignKey:ID"`
	// Photos  Photos `gorm:"foreignKey:ID"`

	// PhotoID Photos `json:"photo_id" form:"photo_id"`
	// User []User `json:"user" gorm:"foreignKey:ID"`
	// UserID  User   `json:"user_id" form:"user_id"`

}

type Comment_post struct {
	Message string `json:"message"`
	Photo   *Photos
	PhotoID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos_id"`
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
