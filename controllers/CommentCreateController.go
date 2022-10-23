package controllers

import (
	"mygram/constants"
	"mygram/helpers"
	"mygram/lib/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateCommentContoller(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	userID := uint(userData["id"].(float64))
	var comment models.Comment

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBindJSON(&comment)
	}
	comment.UserID = userID

	// if e := ctx.Bind(&comment); e != nil {
	// 	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"code":    400,
	// 		"status":  "error",
	// 		"message": e.Error(),
	// 	})
	// 	return
	// }
	// var comments models.Comment
	// comments.Message = comment.Message
	// comments.PhotoID = comment.PhotoID

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
		Message: "Success update photo",
		Data:    comment,
	})
}
