package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photos struct {
	gorm.Model
	// ID       int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url"`
	// UserID   uint   `json:"user_id"`
	User   *User
	UserID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`

	// User []User `json:"user,omitempty" gorm:"foreignKey:ID"`
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

func (p *Photos) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photos) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
