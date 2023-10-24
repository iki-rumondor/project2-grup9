package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project2-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/middleware"
)

func StartServer(handler *customHTTP.UserHandler) *gin.Engine {
	router := gin.Default()

	users := router.Group("users")
	{
		users.POST("/register", middleware.ValidateRegister(), handler.Register)
		users.POST("/login", middleware.ValidateLogin(), handler.Login)
	}

	return router
}
