package customHTTP

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project2-grup9/internal/application"
	"github.com/iki-rumondor/project2-grup9/internal/domain"
)

type UserHandler struct {
	Service *application.UserService
}

func NewHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	body, ok := c.Get("register")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Message: "something went wrong at user service",
		})
		return
	}

	var register = body.(request.Register)

	user := domain.User{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
		Age:      register.Age,
	}

	result, err := h.Service.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	response := response.User{
		Age:      result.Age,
		Email:    result.Email,
		ID:       result.ID,
		Username: result.Username,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	body, ok := c.Get("login")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Message: "something went wrong at user service",
		})
		return
	}

	var login = body.(request.Login)

	user := domain.User{
		Email:    login.Email,
		Password: login.Password,
	}

	jwt, err := h.Service.VerifyUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.JWT{
		Token: jwt,
	})
}
