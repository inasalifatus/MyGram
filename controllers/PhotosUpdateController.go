package controllers

import (
	"mygram/lib/database"
	"mygram/lib/utils"

	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePhotosController(ctx *gin.Context) {
	photoID, _ := strconv.Atoi(ctx.Param("id"))

	if !utils.StringIsNotNumber(ctx.Param("id")) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id",
		})
	}

	var body models.Photos_update

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}

	photo, e := database.GetPhotoById(photoID)
	if e != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}
	photo.Title = utils.CompareStrings(photo.Title, body.Title)
	photo.Caption = utils.CompareStrings(photo.Caption, body.Caption)
	photo.PhotoUrl = utils.CompareStrings(photo.PhotoUrl, body.PhotoUrl)

	_, err := database.InsertPhoto(photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusOK, models.Photos_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success create photo",
		Data:    photo,
	})

}
