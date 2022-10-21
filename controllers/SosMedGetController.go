package controllers

import (
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSosMedController(ctx *gin.Context) {
	sosmed, err := database.GetSosmed()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": err.Error(),
		})
	}

	res := models.Sosmed_Resp{
		Code:    200,
		Status:  "Success",
		Message: "Success Get Photos",
		Data:    sosmed,
	}
	ctx.JSON(http.StatusOK, res)
	return
}
