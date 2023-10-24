package middleware

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
)

func ValidateRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.Register
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
				Message: err.Error(),
			})
			return
		}

		if _, err := govalidator.ValidateStruct(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
				Message: err.Error(),
			})
			return
		}

		c.Set("register", body)
		c.Next()
	}
}

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.Login
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
				Message: err.Error(),
			})
			return
		}

		if _, err := govalidator.ValidateStruct(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
				Message: err.Error(),
			})
			return
		}

		c.Set("login", body)
		c.Next()
	}
}
