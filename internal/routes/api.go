package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project2-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/middleware"
)

func StartServer(handler *customHTTP.UserHandler) *gin.Engine {
	router := gin.Default()

	public_users := router.Group("users")
	users := router.Group("users").Use(middleware.ValidateHeader())
	{
		public_users.POST("/register", middleware.AllUserData(), handler.Register)
		public_users.POST("/login", middleware.UserWithEmail(), handler.Login)
		users.PUT("/", middleware.UserWithEmail(), handler.UpdateUser)
		users.DELETE("/", handler.DeleteUser)
	}

	return router
}
