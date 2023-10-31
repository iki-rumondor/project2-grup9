package repository

import (
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"gorm.io/gorm"
)

type SocialMediaRepoImplementation struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepoImplementation{
		db: db,
	}
}

func (r *SocialMediaRepoImplementation) CreateSocialmedia(socialmedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	if err := r.db.Create(&socialmedia).Error; err != nil {
		return nil, err
	}
	return socialmedia, nil
}

func (r *SocialMediaRepoImplementation) FindSocialmedia(socialmediaID uint) (*domain.SocialMedia, error) {
	var sosmed domain.SocialMedia

	if err := r.db.First(&sosmed, "id = ?", socialmediaID).Error; err != nil {
		return nil, err
	}

	return &sosmed, nil
}

func (r *SocialMediaRepoImplementation) UpdateSocialmedia(socialmedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	var result domain.SocialMedia
	if err := r.db.Model(&domain.SocialMedia{}).Where("id = ?", socialmedia.ID).Updates(&socialmedia).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *SocialMediaRepoImplementation) DeleteSocialmedia(socialmedia *domain.SocialMedia) error {
	if err := r.db.Delete(&socialmedia).Error; err != nil {
		return err
	}
	return nil
}

func (r *SocialMediaRepoImplementation) FindUser(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SocialMediaRepoImplementation) FindSocialmedias(UserID uint) (*[]domain.SocialMedia, error) {
	var sosmed []domain.SocialMedia
	if err := r.db.Preload("UserProfile").Find(&sosmed, "user_id = ?", UserID).Error; err != nil {
		return nil, err
	}

	return &sosmed, nil
}
