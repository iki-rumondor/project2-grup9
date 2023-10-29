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

type PhotoHandler struct {
	Service *application.PhotoService
}

func NewPhotoHandler(service *application.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		Service: service,
	}
}

func (h *PhotoHandler) CreatePhoto(c *gin.Context) {

	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

	var body request.Photo
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

	photo := domain.Photo{
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoUrl: body.PhotoUrl,
		UserID:   userID,
	}

	result, err := h.Service.CreatePhoto(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.CreatePhoto{
		ID:        result.ID,
		Title:     result.Title,
		Caption:   result.Caption,
		PhotoUrl:  result.PhotoUrl,
		UserID:    result.UserID,
		CreatedAt: result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *PhotoHandler) GetPhotos(c *gin.Context) {

	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

	result, err := h.Service.GetPhotos(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var photos response.Photos

	for _, photo := range *result {
		photos.Photos = append(photos.Photos, &response.Photo{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: response.UserProfile{
				Email:    photo.UserProfile.Email,
				Username: photo.UserProfile.Username,
			},
		})
	}

	c.JSON(http.StatusOK, photos.Photos)
}

func (h *PhotoHandler) UpdatePhoto(c *gin.Context) {

	var body request.Photo
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
	photoID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	photo := domain.UpdatePhoto{
		ID:       uint(photoID),
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoUrl: body.PhotoUrl,
	}

	result, err := h.Service.UpdatePhoto(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdatePhoto{
		ID:        result.ID,
		Title:     result.Title,
		Caption:   result.Caption,
		PhotoUrl:  result.PhotoUrl,
		UserID:    result.UserID,
		UpdatedAt: result.UpdatedAt,
	})
}

func (h *PhotoHandler) DeletePhoto(c *gin.Context) {

	urlParam := c.Param("id")
	photoID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	photo := domain.Photo{
		ID: uint(photoID),
	}

	if err := h.Service.DeletePhoto(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Your photo has been successfully deleted",
	})
}
