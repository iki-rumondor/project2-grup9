package application

import (
	"errors"

	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
	"github.com/iki-rumondor/project2-grup9/internal/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	// save user into database
	user, err := s.Repo.Save(user)
	if err != nil{
		return nil, errors.New("failed to save user into database")
	}

	return user, nil
}

func (s *UserService) VerifyUser(user *domain.User) (string, error) {
	// find user by email from database
	result, err := s.Repo.FindByEmail(user.Email)
	if err != nil{
		return "", errors.New("sorry, the provided email is not registered in our system")
	}
	
	// verify user password
	if err := utils.ComparePassword(result.Password, user.Password); err != nil{
		return "", errors.New("whoops! password mismatch")
	}

	data := map[string]interface{}{
		"id": result.ID,
	}

	// create jwt token
	jwt, err := utils.GenerateToken(data)
	if err != nil{
		return "", err
	}

	return jwt, nil
}


