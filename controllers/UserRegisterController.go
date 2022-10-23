package controllers

import (
	"mygram/constants"
	"mygram/helpers"
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserController(ctx *gin.Context) {
	// db := config.GetDB()
	contentType := helpers.GetContentType(ctx)
	// _, _ = db, contentType
	User := models.User{}

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	user, err := database.InsertUser(User)
	// err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}
