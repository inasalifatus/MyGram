package controllers

import (
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCommentController(ctx *gin.Context) {
	comment, err := database.GetComment()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": err.Error(),
		})
	}

	res := models.Comment_Resp{
		Code:    200,
		Status:  "Success",
		Message: "Success Get Photos",
		Data:    comment,
	}
	ctx.JSON(http.StatusOK, res)
	return
}
