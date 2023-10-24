package repository

import (
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

func (r *UserRepoImplementation) Save(user *domain.User) (*domain.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImplementation) FindByEmail(email string) (*domain.User, error){
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil{
		return nil, err
	}

	return &user, nil
}