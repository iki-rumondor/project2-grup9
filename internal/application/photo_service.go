package application

import (
	"errors"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
)

type PhotoService struct {
	Repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) *PhotoService {
	return &PhotoService{
		Repo: repo,
	}
}

func (s *PhotoService) CreatePhoto(photo *domain.Photo) (*domain.Photo, error) {

	result, err := s.Repo.Save(photo)
	if err != nil {
		return nil, errors.New("failed to save photo into database")
	}

	return result, nil
}

func (s *PhotoService) GetPhotos(userID uint) (*[]domain.Photo, error) {

	photos, err := s.Repo.FindPhotos(userID)
	if err != nil {
		return nil, errors.New("failed to get user photos from database")
	}

	return photos, nil
}
