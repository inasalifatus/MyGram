package controllers

import (
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCommentContoller(ctx *gin.Context) {
	var body models.Comment_post

	if e := ctx.Bind(&body); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": e.Error(),
		})
		return
	}
	var comments models.Comment
	comments.Message = body.Message
	comments.PhotoID = body.PhotoID

	_, err := database.InsertComment(comments)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusCreated, models.Comment_response_single{
		Code:    201,
		Status:  "Success",
		Message: "Success create photo",
		Data:    comments,
	})
}
