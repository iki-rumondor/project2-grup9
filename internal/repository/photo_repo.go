package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type PhotoRepository interface {
	FindUser(uint) (*domain.User, error)
	Save(*domain.Photo) (*domain.Photo, error)
	FindPhotos(uint) (*[]domain.Photo, error)
}
