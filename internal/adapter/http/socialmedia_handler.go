package customHTTP

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project2-grup9/internal/application"
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"gorm.io/gorm"
)

type SocialMediaHandler struct {
	Service *application.SocialMediaService
}

func NewSocialMediaHandler(service *application.SocialMediaService) *SocialMediaHandler {
	return &SocialMediaHandler{
		Service: service,
	}
}

func (h *SocialMediaHandler) CreateSocialmedia(c *gin.Context) {
	userID := c.GetUint("user_id")

	var body request.SocialMedia
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	sosmed := domain.SocialMedia{
		Name:           body.Name,
		SocialMediaURL: body.SocialMediaURL,
		UserID:         userID,
	}

	result, err := h.Service.CreateSocialmedia(&sosmed)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
			return
		}
		
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.CreateSocialmedia{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaURL: result.SocialMediaURL,
		UserID:         result.UserID,
		CreatedAt:      result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *SocialMediaHandler) GetSocialmedia(c *gin.Context) {
	UserID := c.GetUint("user_id")

	results, err := h.Service.GetSocialMedia(UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var socialmedia = response.Sosmeds{}

	for _, sosmed := range *results {
		socialmedia.Sosmeds = append(socialmedia.Sosmeds, &response.Sosmed{
			ID:             sosmed.ID,
			Name:           sosmed.Name,
			SocialMediaURL: sosmed.SocialMediaURL,
			UserID:         sosmed.UserID,
			CreatedAt:      sosmed.CreatedAt,
			UpdatedAt:      sosmed.UpdatedAt,
			User: &response.UserSosmed{
				ID:       sosmed.User.ID,
				Username: sosmed.User.Username,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": socialmedia.Sosmeds,
	})
}

func (h *SocialMediaHandler) UpdateSocialmedia(c *gin.Context) {
	var body request.SocialMedia
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	urlParam := c.Param("id")
	id, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	userID := c.GetUint("user_id")

	sosmed := domain.SocialMedia{
		ID:             uint(id),
		Name:           body.Name,
		SocialMediaURL: body.SocialMediaURL,
		UserID:         userID,
	}

	result, err := h.Service.UpdateSocialmedia(&sosmed)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateSosmed{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaURL: result.SocialMediaURL,
		UserID:         result.UserID,
		UpdatedAt:      result.UpdatedAt,
	})
}

func (h *SocialMediaHandler) DeleteSocialmedia(c *gin.Context) {
	urlParam := c.Param("id")
	sosmedID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	userID := c.GetUint("user_id")
	sosmed := domain.SocialMedia{
		ID:     uint(sosmedID),
		UserID: userID,
	}

	if err := h.Service.DeleteSocialMedia(&sosmed); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Your social media has been successfully deleted",
	})
}
