package controllers

import (
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPhotosController(ctx *gin.Context) {
	photos, err := database.GetPhotos()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": err.Error(),
		})
	}
	res := models.Photos_Resp{
		Code:    200,
		Status:  "Success",
		Message: "Success Get Photos",
		Data:    photos,
	}
	ctx.JSON(http.StatusOK, res)
	return
}
