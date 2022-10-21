package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	UserID   int    `gorm:"AUTO_INCREMENT" json:"userID"`
	Username string `gorm:"type:varchar(100)" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"type:varchar(45);not null" json:"email" form:"email" valid:"required"`
	Password string `gorm:"type:varchar(100);not null" json:"password" form:"password" valid:"required"`
	Age      int    `json:"age" form:"age" valid:"required"`
	Token    string `json:"token,omitempty" form:"token"`

	// Photos Photos `gorm:"foreignKey:ID"`
}

type User_response_single struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

var UserData = []User{}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
