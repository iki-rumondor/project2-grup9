package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project2-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project2-grup9/internal/application"
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/utils"
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
	// userID := c.GetUint("user_id")
	defer utils.Recovery(c)

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
	}

	result, err := h.Service.CreateSocialmedia(&sosmed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.Socialmedia{
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
	defer utils.Recovery(c)

	result, err := h.Service.GetSocialMedia(UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var socialmedia response.Sosmed

	for _, sosmed := range *result {
		socialmedia.SocialMedia = append(socialmedia.SocialMedia, &response.Socialmedia{
			ID:             sosmed.ID,
			Name:           sosmed.Name,
			SocialMediaURL: sosmed.SocialMediaURL,
			UserID:         sosmed.UserID,
			CreatedAt:      sosmed.CreatedAt,
			UpdatedAt:      sosmed.UpdatedAt,
			User: response.UserProfiles{
				ID:              sosmed.UserProfiles.ID,
				Email:           sosmed.UserProfiles.Email,
				ProfileImageURL: sosmed.UserProfiles.ProfileImageURL,
			},
		})
	}

	c.JSON(http.StatusOK, socialmedia.SocialMedia)
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
	_, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	sosmed := domain.SocialMedia{
		Name:           body.Name,
		SocialMediaURL: body.SocialMediaURL,
	}

	result, err := h.Service.UpdateSocialmedia(&sosmed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Socialmedia{
		ID:             result.ID,
		Name:           result.Name,
		SocialMediaURL: result.SocialMediaURL,
		UserID:         result.UserID,
		CreatedAt:      result.CreatedAt,
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

	sosmed := domain.SocialMedia{
		ID: uint(sosmedID),
	}

	if err := h.Service.DeleteSocialmedia(&sosmed); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Your social media has been successfully deleted",
	})
}
