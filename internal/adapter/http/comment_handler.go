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
	// userID := c.GetUint("user_id")
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

	for _, comment := range *result {
		comments.Comments = append(comments.Comments, &response.Comment{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: response.UserProfile{
				ID:       comment.UserProfile.ID,
				Email:    comment.UserProfile.Email,
				Username: comment.UserProfile.Username,
			},
			Photo: response.PhotoProfile{
				ID:       comment.PhotoProfile.ID,
				Title:    comment.PhotoProfile.Title,
				Caption:  comment.PhotoProfile.Caption,
				PhotoUrl: comment.PhotoProfile.PhotoUrl,
				UserID:   comment.PhotoProfile.UserID,
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
	_, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	comment := domain.UpdateComment{
		Message: body.Message,
	}

	result, err := h.Service.UpdateComment(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateComment{
		ID:        result.ID,
		Tittle:    result.Tittle,
		PhotoID:   result.PhotoID,
		Message:   result.Message,
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
