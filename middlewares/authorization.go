package middlewares

import (
	"mygram/config"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		photoID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photos{}

		err = config.DB.Select("user_id").First(&Photo, uint(photoID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})
			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}

		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = config.DB.Select("user_id").First(&comment, uint(commentID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})
			return
		}

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}

		ctx.Next()
	}
}

func SosmedAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sosmedID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialmedia := models.Sosmed{}

		err = config.DB.Select("user_id").First(&socialmedia, uint(sosmedID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})
			return
		}

		if socialmedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}

		ctx.Next()
	}
}
