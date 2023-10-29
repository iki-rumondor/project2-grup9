package customHTTP

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project2-grup9/internal/application"
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/utils"
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

	body, ok := c.Get("userData")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: "something went wrong at user service",
		})
		return
	}

	var register = body.(request.AllUserData)

	user := domain.User{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
		Age:      register.Age,
	}

	result, err := h.Service.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.User{
		Age:       result.Age,
		Email:     result.Email,
		ID:        result.ID,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	body, ok := c.Get("userWithEmail")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: "something went wrong at user service",
		})
		return
	}

	var login = body.(request.UserWithEmail)

	user := domain.User{
		Email:    login.Email,
		Password: login.Password,
	}

	jwt, err := h.Service.VerifyUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.JWT{
		Token: jwt,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	mapClaims, err := utils.VerifyToken(c.GetString("jwt"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Message{
			Message: err.Error(),
		})
		return
	}

	body, ok := c.Get("userWithEmail")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: "something went wrong at user service",
		})
		return
	}

	var user = body.(request.UserWithEmail)

	userID := uint(mapClaims["id"].(float64))

	req := domain.User{
		ID:       userID,
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := h.Service.UpdateUser(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdatedUser{
		ID:        result.ID,
		Email:     result.Email,
		Username:  result.Username,
		Age:       result.Age,
		UpdatedAt: result.UpdatedAt,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	mapClaims, err := utils.VerifyToken(c.GetString("jwt"))
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
				Message: "something went wrong, please check your credentials",
			})
			return
		}
	}()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Message{
			Message: err.Error(),
		})
		return
	}

	userID := uint(mapClaims["id"].(float64))

	user := domain.User{
		ID: userID,
	}

	if err := h.Service.DeleteUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Your account has been successfully deleted",
	})
}
