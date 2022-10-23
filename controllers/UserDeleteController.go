package controllers

import (
	"mygram/lib/database"
	"mygram/lib/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUserController(ctx *gin.Context) {
	//
	id, _ := strconv.Atoi(ctx.Param("id"))

	if !utils.StringIsNotNumber(ctx.Param("id")) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})

		return
	}

	err := database.DeleteUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "Your account has been succesfully deleted",
	})
	return
}
