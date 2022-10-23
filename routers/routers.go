package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	users.POST("/register", controllers.RegisterUserController)
	users.POST("/login", controllers.LoginUserController)
	users.PUT("/update/:id", middlewares.Authentication(), controllers.UpdateUserController)
	users.DELETE("/delete/:id", middlewares.Authentication(), controllers.DeleteUserController)

	photos := router.Group("/photos")
	photos.Use(middlewares.Authentication())
	photos.POST("/create", controllers.CreatePhotosController)
	photos.GET("/get", controllers.GetPhotosController)
	photos.PUT("/update/:id", middlewares.PhotoAuthorization(), controllers.UpdatePhotosController)
	photos.DELETE("/delete/:id", middlewares.PhotoAuthorization(), controllers.DeletePhotoController)

	comments := router.Group("/comments")
	comments.Use(middlewares.Authentication())
	comments.POST("/create", controllers.CreateCommentContoller)
	comments.GET("/get", controllers.GetCommentController)
	comments.PUT("/update/:id", middlewares.CommentAuthorization(), controllers.UpdateCommentController)
	comments.DELETE("/delete/:id", middlewares.CommentAuthorization(), controllers.DeleteCommentController)

	socialMedias := router.Group("/socialmedia")
	socialMedias.Use(middlewares.Authentication())
	socialMedias.POST("/create", controllers.CreateSosMedController)
	socialMedias.GET("/get", controllers.GetSosMedController)
	socialMedias.PUT("/update/:id", middlewares.SosmedAuthorization(), controllers.UpdateSosMedController)
	socialMedias.DELETE("/delete/:id", middlewares.SosmedAuthorization(), controllers.DeleteSosMedController)

	return router
}
