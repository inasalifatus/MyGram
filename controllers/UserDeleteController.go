package controllers

import (
	"fmt"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUserController(ctx *gin.Context) {
	ID := ctx.Param("ID")
	intID, _ := strconv.Atoi(ID)
	condition := false
	var userIndex int
	// var user models.User
	var UserData = []models.User{}

	for i, user := range UserData {
		if intID == user.ID {
			condition = true
			userIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintf("user with id %v not found", intID),
		})
		return
	}

	copy(UserData[userIndex:], UserData[userIndex+1:])
	UserData[len(UserData)-1] = models.User{}
	UserData = UserData[:len(UserData)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("user with id %v has been successfully updated", intID),
	})
}
