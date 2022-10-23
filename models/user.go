package models

import (
	"errors"
	"fmt"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID int `gorm:"primarykey;AUTO_INCREMENT" json:"ID" form:"id"`
	// UserID   int    `gorm:"AUTO_INCREMENT" json:"userID"`
	Username string `gorm:"type:varchar(100);not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"type:varchar(45);not null;uniqueIndex" json:"email" form:"email" valid:"required"`
	Password string `gorm:"type:varchar(100);not null" json:"password" form:"password" valid:"required"`
	Age      int    `json:"age" form:"age" valid:"required"`
	Token    string `json:"token,omitempty" form:"token"`
	// Photos   []Photos `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	// Sosmed  []Sosmed `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sosmed"`

	// Photos Photos `gorm:"foreignKey:ID"`
}

// var UserData = []User{}

type User_response_single struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type User_update struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	fmt.Println("User Before Create()")

	if u.Age < 8 {
		err = errors.New("Age user is too young")
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
