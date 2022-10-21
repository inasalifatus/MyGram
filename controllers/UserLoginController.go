package controllers

import (
	"mygram/config"
	"mygram/constants"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUserController(ctx *gin.Context) {
	user := models.User{}
	contentType := helpers.GetContentType(ctx)
	_, _ = user, contentType
	ctx.Bind(&user)
	password := ""

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	password = user.Password

	// users, err := database.LoginUser(&user)
	err := config.DB.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unautorized",
			"message": "Invalid Email/Password",
		})
		return
	}
	token := helpers.GenerateToken(user.ID, user.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}
