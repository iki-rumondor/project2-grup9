package application

import (
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
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
		return nil, err
	}

	return result, nil
}

func (s *SocialMediaService) GetSocialMedia(UserID uint) (*[]domain.SocialMedia, error) {

	sosmed, err := s.Repo.FindSocialmedias(UserID)
	if err != nil {
		return nil, err
	}

	return sosmed, nil
}

func (s *SocialMediaService) UpdateSocialmedia(sosmed *domain.SocialMedia) (*domain.SocialMedia, error) {

	updatedsosmed, err := s.Repo.UpdateSocialmedia(sosmed)
	if err != nil {
		return nil, err
	}

	return updatedsosmed, nil
}

func (s *SocialMediaService) DeleteSocialMedia(sosmed *domain.SocialMedia) error {

	if err := s.Repo.DeleteSocialmedia(sosmed); err != nil {
		return err
	}

	return nil
}
