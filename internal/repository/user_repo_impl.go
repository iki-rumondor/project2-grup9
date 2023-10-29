package repository

import (
	"errors"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"gorm.io/gorm"
)

type UserRepoImplementation struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &UserRepoImplementation{
		db: db,
	}
}

func (r *UserRepoImplementation) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoImplementation) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoImplementation) IsUniqueEmail(user *domain.User) bool {
	var result domain.User
	if err := r.db.Where("email = ? AND id != ?", user.Email, user.ID).First(&result).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func (r *UserRepoImplementation) Save(user *domain.User) (*domain.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImplementation) Update(user *domain.User) (*domain.User, error) {
	var updatedUser domain.User
	if err := r.db.Updates(user).First(&updatedUser).Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
