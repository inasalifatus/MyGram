package controllers

import (
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePhotosController(ctx *gin.Context) {
	var body models.Photos_post

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": e.Error(),
		})
		return
	}

	var photos models.Photos
	photos.Title = body.Title
	photos.Caption = body.Caption
	photos.PhotoUrl = body.PhotoUrl

	_, err := database.InsertPhoto(photos)
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
		Data:    photos,
	})
}
