package application

import (
	"errors"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
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
		return nil, err
	}

	return result, nil
}

func (s *CommentService) GetComments(UserID uint) ([]*domain.Comment, error) {

	comment, err := s.Repo.FindComments(UserID)
	if err != nil {
		return nil, errors.New("failed to get user comment from database")
	}

	return comment, nil
}

func (s *CommentService) UpdateComment(comment *domain.Comment) (*domain.Comment, error) {

	comment, err := s.Repo.UpdateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentService) DeleteComment(comment *domain.Comment) error {

	if err := s.Repo.DeleteComment(comment); err != nil {
		return err
	}

	return nil
}
