package middleware

import (
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
)

func AllUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.AllUserData
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

		c.Set("userData", body)
		c.Next()
	}
}

func UserWithEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.UserWithEmail
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

		c.Set("userWithEmail", body)
		c.Next()
	}
}

func ValidateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
				Message: "Bearer token is not valid",
			})
			return
		}

		stringToken := strings.Split(headerToken, " ")[1]

		c.Set("jwt", stringToken)
		c.Next()
	}
}
