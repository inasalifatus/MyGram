package controllers

import (
	"mygram/lib/database"
	"mygram/lib/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteSosMedController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if !utils.StringIsNotNumber(ctx.Param("id")) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Failed",
			"message": "Invalid Id",
		})
		return
	}

	err := database.DeleteSosmedById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "sosmed successfully deleted",
	})

}
