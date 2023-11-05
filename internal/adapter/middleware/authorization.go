package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project2-grup9/internal/utils"
)

func GetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		mapClaims, err := utils.VerifyToken(c.GetString("jwt"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Message{
				Message: err.Error(),
			})
			return
		}

		userID := uint(mapClaims["id"].(float64))
		c.Set("user_id", userID)
		c.Next()
	}
}
