package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project2-grup9/internal/adapter/http"
)

func StartServer(handler *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/")

	api := router.Group("api/v1")
	{
		api.GET("/comments", handler.CommentHandler.GetComments)
		api.POST("/comments", handler.CommentHandler.CreateComment)
		api.PUT("/comments/:id", handler.CommentHandler.UpdateComment)
		api.DELETE("/comments/:id", handler.CommentHandler.DeleteComment)
	}

	return router
}
