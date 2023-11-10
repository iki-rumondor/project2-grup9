package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type PhotoRepository interface {
	Create(*domain.Photo) (*domain.Photo, error)
	Delete(*domain.Photo) error
	Update(*domain.Photo) (*domain.Photo, error)
	FindPhotos(uint) (*[]domain.Photo, error)
}
