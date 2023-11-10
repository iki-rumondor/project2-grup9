package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type CommentRepository interface {
	CreateComment(*domain.Comment) (*domain.Comment, error)
	UpdateComment(*domain.Comment) (*domain.Comment, error)
	DeleteComment(*domain.Comment) error
	FindComments(uint) ([]*domain.Comment, error)
}
