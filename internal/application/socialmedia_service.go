package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
	"gorm.io/gorm"
)

type SocialMediaService struct {
	Repo repository.SocialMediaRepository
}

func NewSocialMediaService(repo repository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		Repo: repo,
	}
}

func (s *SocialMediaService) CreateSocialmedia(sosmed *domain.SocialMedia) (*domain.SocialMedia, error) {

	result, err := s.Repo.CreateSocialmedia(sosmed)
	if err != nil {
		return nil, errors.New("failed to save comment into database")
	}

	return result, nil
}

func (s *SocialMediaService) GetSocialMedia(UserID uint) (*[]domain.SocialMedia, error) {

	sosmed, err := s.Repo.FindSocialmedias(UserID)
	if err != nil {
		return nil, errors.New("failed to get user comment from database")
	}

	return sosmed, nil
}

func (s *SocialMediaService) UpdateSocialmedia(sosmed *domain.SocialMedia) (*domain.SocialMedia, error) {

	_, err := s.Repo.FindSocialmedia(sosmed.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("comment with ID %d not found", sosmed.ID)
	}

	if err != nil {
		return nil, errors.New("failed to get comment from the database")
	}

	updatedsosmed, err := s.Repo.UpdateSocialmedia(sosmed)
	if err != nil {
		return nil, errors.New("failed to update comment in the database")
	}

	return updatedsosmed, nil
}
