package database

import (
	"errors"
	"mygram/config"
	"mygram/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) (interface{}, error) {
	user := models.User{}
	ctx.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func InsertUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// func LoginUser(user *models.User) (interface{}, error) {
// 	var err error
// 	if err = config.DB.Where("email=? AND password=?", user.Email, user.Password).First(user).Error; err != nil {
// 		return nil, err
// 	}

// 	user.Token, err = middlewares.CreateToken(int(user.ID))
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := config.DB.Save(user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

func GetUserById(id int) (models.User, error) {
	var user models.User

	if rows := config.DB.Find(&user, id).RowsAffected; rows < 1 {
		err := errors.New("user not found")
		return user, err
	}

	return user, nil
}
