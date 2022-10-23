package controllers

import (
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSosMedController(ctx *gin.Context) {
	var body models.Sosmed_post

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": e.Error(),
		})
		return
	}

	var sosmed models.Sosmed
	sosmed.Name = body.Name
	sosmed.SosmedURL = body.SosmedURL

	_, err := database.CreateSosmed(sosmed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusCreated, models.Sosmed_response_single{
		Code:    201,
		Status:  "Success",
		Message: "Success create social media",
		Data:    sosmed,
	})

}
