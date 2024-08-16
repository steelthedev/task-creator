package services

import (
	"github.com/steelthedev/task-handler/data"
	"github.com/steelthedev/task-handler/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

// Add user service
// Function adds user
func (us *UserService) AddUser(user *data.User) (*data.User, error) {
	user, err := us.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUserByID(id uint) (*data.User, error) {
	user, err := us.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUserByEmail(email string) (*data.User, error) {
	user, err := us.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
