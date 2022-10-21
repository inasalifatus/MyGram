package routers

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	users.POST("/register", controllers.RegisterUserController)
	users.POST("/login", controllers.LoginUserController)
	users.PUT("/update/:id", controllers.UpdateUserController)
	users.DELETE("/delete/:id", controllers.DeleteUserController)

	photos := router.Group("/photos")
	photos.POST("/create", controllers.CreatePhotosController)
	photos.GET("/get", controllers.GetPhotosController)
	photos.PUT("/update/:id", controllers.UpdatePhotosController)
	photos.DELETE("/delete/:id", controllers.DeletePhotoController)

	// comments := router.Group("/comments")
	// comments.POST("/create", controllers.CreateCommentContoller)
	// comments.GET("/get", controllers.GetCommentController)
	// comments.PUT("/update/:id", controllers.UpdateCommentController)
	// comments.DELETE("/delete", controllers.DeleteCommentController)

	// socialMedias := router.Group("/socialMedia")
	// socialMedias.POST("/create", controllers.CreateSosMedController)
	// socialMedias.GET("/get", controllers.GetSosMedController)
	// socialMedias.PUT("/update/:id", controllers.UpdateSosMedController)
	// socialMedias.DELETE("/delete/:id", controllers.DeleteSosMedController)

	return router
}
