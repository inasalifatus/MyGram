package controllers

import (
	"mygram/lib/database"
	"mygram/lib/utils"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUserController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if !utils.StringIsNotNumber(ctx.Param("id")) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id user",
		})
		return
	}

	var body models.User_update

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id Uer",
		})
		return
	}

	user, e := database.GetUserById(id)
	if e != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}

	user.Username = utils.CompareStrings(user.Username, body.Username)
	user.Email = utils.CompareStrings(user.Email, body.Email)

	_, err := database.InsertUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.User_response_single{
		Code:    200,
		Status:  "success",
		Message: "succces update user",
		Data:    user,
	})

}
