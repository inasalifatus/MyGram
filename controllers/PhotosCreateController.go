package controllers

import (
	"mygram/constants"
	"mygram/helpers"
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhotosController(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Photos := models.Photos{}
	userID := uint(userData["id"].(float64))
	Photos.UserID = userID

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON((&Photos))
	} else {
		ctx.ShouldBind(&Photos)
	}

	_, err := database.InsertPhoto(Photos)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusCreated, models.Photos_response_single{
		Code:    201,
		Status:  "Success",
		Message: "Success create photo",
		Data:    Photos,
	})
}
