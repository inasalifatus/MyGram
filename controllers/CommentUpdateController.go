package controllers

import (
	"mygram/lib/database"
	"mygram/lib/utils"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCommentController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if !utils.StringIsNotNumber(ctx.Param("id")) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id",
		})
	}

	var body models.Comment_update

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}

	comment, e := database.GetCommentById(id)
	if e != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
		return
	}

	comment.Message = utils.CompareStrings(comment.Message, body.Message)

	_, err := database.InsertComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusOK, models.Comment_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success create photo",
		Data:    comment,
	})

}
