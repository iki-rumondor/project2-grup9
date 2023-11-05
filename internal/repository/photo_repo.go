package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type PhotoRepository interface {
	FindUser(uint) (*domain.User, error)
	Create(*domain.Photo) (*domain.Photo, error)
	Delete(*domain.Photo) error
	Update(*domain.Photo) (*domain.Photo, error)
	FindPhotos(uint) (*[]domain.Photo, error)
	FindAllUserPhotos() (*[]domain.Photo, error)
	FindPhoto(uint) (*domain.Photo, error)
}
