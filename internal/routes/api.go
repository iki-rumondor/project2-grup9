package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project2-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/middleware"
)

func StartServer(handler *customHTTP.Handlers) *gin.Engine {
	
	router := gin.Default()

	public := router.Group("")
	{
		public.POST("users/register", middleware.AllUserData(), handler.UserHandler.Register)
		public.POST("users/login", middleware.UserWithEmail(), handler.UserHandler.Login)
	}

	users := router.Group("users").Use(middleware.ValidateHeader())
	{
		users.PUT("/", middleware.UserWithEmail(), handler.UserHandler.UpdateUser)
		users.DELETE("/", middleware.GetUserID(), handler.UserHandler.DeleteUser)
	}

	photos := router.Group("photos").Use(middleware.ValidateHeader(), middleware.GetUserID())
	{
		photos.POST("/", handler.PhotoHandler.CreatePhoto)
		photos.GET("/", handler.PhotoHandler.GetPhotos)
		photos.PUT("/:id", handler.PhotoHandler.UpdatePhoto)
		photos.DELETE("/:id", handler.PhotoHandler.DeletePhoto)
	}
  
  	api.GET("/comments", handler.CommentHandler.GetComments)
	api.POST("/comments", handler.CommentHandler.CreateComment)
	api.PUT("/comments/:id", handler.CommentHandler.UpdateComment)
	api.DELETE("/comments/:id", handler.CommentHandler.DeleteComment)

	return router
}
