package repository

import (
	"github.com/steelthedev/task-handler/conn"
	"github.com/steelthedev/task-handler/data"
)

type userInterface interface {
	Create(user *data.User) (*data.User, error)
	Get(id uint) (*data.User, error)
	GetByEmail(email string) (*data.User, error)
}

type UserRepository struct {
}

func (u *UserRepository) Create(user *data.User) (*data.User, error) {
	if err := conn.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Get(id uint) (*data.User, error) {
	var user data.User
	if err := conn.DB.Where("ID=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetByEmail(email string) (*data.User, error) {
	var user data.User
	if err := conn.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
