package controllers

import (
	"mygram/lib/database"
	"mygram/lib/utils"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateSosMedController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if !utils.StringIsNotNumber(ctx.Param("id")) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id",
		})
	}

	var body models.Sosmed_post

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}

	sosmed, e := database.GetSosmedById(id)
	if e != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}

	sosmed.Name = utils.CompareStrings(sosmed.Name, body.Name)
	sosmed.SosmedURL = utils.CompareStrings(sosmed.SosmedURL, body.SosmedURL)

	_, err := database.CreateSosmed(sosmed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusOK, models.Sosmed_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success update photo",
		Data:    sosmed,
	})

}
