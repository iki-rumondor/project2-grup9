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

type CommentHandler struct {
	Service *application.CommentService
}

func NewCommentHandler(service *application.CommentService) *CommentHandler {
	return &CommentHandler{
		Service: service,
	}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

	var body request.Comment
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

	comment := domain.Comment{
		Message: body.Message,
		PhotoID: body.PhotoID,
		UserID:  userID,
	}

	result, err := h.Service.CreateComment(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.CreateComment{
		ID:        result.ID,
		Message:   result.Message,
		PhotoID:   result.PhotoID,
		UserID:    result.UserID,
		CreatedAt: result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CommentHandler) GetComments(c *gin.Context) {

	UserID := c.GetUint("user_id")
	defer utils.Recovery(c)

	result, err := h.Service.GetComments(UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var comments response.Comments

	for _, comment := range result {
		comments.Comments = append(comments.Comments, &response.Comment{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: response.UserProfile{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: response.PhotoProfile{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	c.JSON(http.StatusOK, comments.Comments)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	var body request.CommentID
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
	UserID := c.GetUint("user_id")

	comment := domain.Comment{
		ID:      uint(id),
		Message: body.Message,
		UserID:  UserID,
	}

	result, err := h.Service.UpdateComment(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &response.UpdateComment{
		ID:        result.ID,
		Title:     result.Photo.Title,
		Caption:   result.Photo.Caption,
		PhotoUrl:  result.Photo.PhotoUrl,
		UserID:    result.UserID,
		UpdatedAt: result.UpdatedAt,
	})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	urlParam := c.Param("id")
	commentID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	comment := domain.Comment{
		ID: uint(commentID),
	}

	if err := h.Service.DeleteComment(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Your comment has been successfully deleted",
	})
}
