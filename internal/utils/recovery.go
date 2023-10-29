package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
)

func Recovery(c *gin.Context) {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: "opps... something went wrong, please check your request!!",
		})
	}
}
