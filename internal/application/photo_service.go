package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
	"gorm.io/gorm"
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

	result, err := s.Repo.Create(photo)
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

func (s *PhotoService) UpdatePhoto(photo *domain.UpdatePhoto) (*domain.Photo, error) {

	_, err := s.Repo.FindPhoto(photo.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("photo with id %d id not found", photo.ID)
	}
	if err != nil {
		return nil, errors.New("failed to get photo from database")
	}

	photos, err := s.Repo.Update(photo)
	if err != nil {
		return nil, errors.New("failed to update photo to database")
	}

	return photos, nil
}

func (s *PhotoService) DeletePhoto(photo *domain.Photo) error {

	_, err := s.Repo.FindPhoto(photo.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("photo with id %d id not found", photo.ID)
	}
	if err != nil {
		return errors.New("failed to get photo from database")
	}

	if err := s.Repo.Delete(photo); err != nil {
		return errors.New("we encountered an issue while trying to delete the photo")
	}

	return nil
}
