package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type CommentRepository interface {
	CreateComment(*domain.Comment) (*domain.Comment, error)
	FindComment(uint) ([]domain.Comment, error)
	UpdateComment(*domain.UpdateComment) (*domain.Comment, error)
	DeleteComment(*domain.Comment) error
	FindUser(uint) (*domain.User, error)   //manggil user
	FindPhoto(uint) (*domain.Photo, error) //mangggil poto dari user
	FindComments(UserID uint) ([]*domain.Comment, error)
}
