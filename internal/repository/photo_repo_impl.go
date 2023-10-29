package repository

import (
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"gorm.io/gorm"
)

type PhotoRepoImplementation struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepoImplementation{
		db: db,
	}
}

func (r *PhotoRepoImplementation) Save(photo *domain.Photo) (*domain.Photo, error) {
	if err := r.db.Save(photo).Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *PhotoRepoImplementation) FindUser(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PhotoRepoImplementation) FindPhotos(userID uint) (*[]domain.Photo, error) {
	var photos []domain.Photo
	if err := r.db.Preload("UserProfile").Find(&photos, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &photos, nil
}
