package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type SocialMediaRepository interface {
	CreateSocialmedia(*domain.SocialMedia) (*domain.SocialMedia, error)
	FindSocialmedia(uint) (*domain.SocialMedia, error)
	UpdateSocialmedia(*domain.SocialMedia) (*domain.SocialMedia, error)
	DeleteSocialmedia(*domain.SocialMedia) error
	FindUser(uint) (*domain.User, error)
	FindSocialmedias(UserID uint) (*[]domain.SocialMedia, error)
	FindAllUserSocialmedias() (*[]domain.SocialMedia, error)
}
