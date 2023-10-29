package repository

import "github.com/iki-rumondor/project2-grup9/internal/domain"

type UserRepository interface {
	Save(*domain.User) (*domain.User, error)
	Update(*domain.User) (*domain.User, error)
	Delete(*domain.User) error
	FindByEmail(string) (*domain.User, error)
	FindByID(uint) (*domain.User, error)

	IsUniqueEmail(*domain.User) bool
}
