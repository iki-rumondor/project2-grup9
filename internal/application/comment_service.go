package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
	"gorm.io/gorm"
)

type CommentService struct {
	Repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{
		Repo: repo,
	}
}

func (s *CommentService) CreateComment(comment *domain.Comment) (*domain.Comment, error) {

	result, err := s.Repo.CreateComment(comment)
	if err != nil {
		return nil, errors.New("failed to save comment into database")
	}

	return result, nil
}

func (s *CommentService) GetComments(UserID uint) (*[]domain.Comment, error) {

	comment, err := s.Repo.FindComments(UserID)
	if err != nil {
		return nil, errors.New("failed to get user comment from database")
	}

	return comment, nil
}

func (s *CommentService) UpdateComment(comment *domain.UpdateComment) (*domain.Comment, error) {
	_, err := s.Repo.FindComment(comment.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("comment with ID %d not found", comment.ID)
	}

	if err != nil {
		return nil, errors.New("failed to get comment from the database")
	}

	updatedComment, err := s.Repo.UpdateComment(comment)
	if err != nil {
		return nil, errors.New("failed to update comment in the database")
	}

	return updatedComment, nil
}

func (s *CommentService) DeleteComment(comment *domain.Comment) error {

	_, err := s.Repo.FindComment(comment.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("comment with id %d id not found", comment.ID)
	}
	if err != nil {
		return errors.New("failed to get comment from database")
	}

	if err := s.Repo.DeleteComment(comment); err != nil {
		return errors.New("we encountered an issue while trying to delete the comment")
	}

	return nil
}
