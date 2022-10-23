package database

import (
	"errors"
	"mygram/config"
	"mygram/models"
)

// func RegisterUser(ctx *gin.Context) (interface{}, error) {
// 	user := models.User{}
// 	ctx.Bind(&user)
// 	if err := config.DB.Save(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

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

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	if rows := config.DB.Find(&user, id).RowsAffected; rows < 1 {
		err := errors.New("user not found")
		return user, err
	}

	return user, nil
}

func DeleteUserById(id int) error {

	rows := config.DB.Delete(&models.User{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("user to be deleted is not found")
		return err
	}
	return nil

}
