package controllers

import (
	"fmt"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUserController(ctx *gin.Context) {
	ID := ctx.Param("id")
	intID, _ := strconv.Atoi(ID)
	condition := false
	var updatedUser models.User
	var UserData = []models.User{}
	// var dataUser models.us

	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, user := range UserData {
		if intID == user.ID {
			condition = true
			UserData[i] = updatedUser
			UserData[i].ID = intID
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("user with id %v has been successfully updated", intID),
	})

}
