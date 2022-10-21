package controllers

import (
	"mygram/constants"
	"mygram/helpers"
	"mygram/lib/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserController(ctx *gin.Context) {
	user, err := database.RegisterUser(ctx)
	contentType := helpers.GetContentType(ctx)
	_, _ = user, contentType

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}
