package repository

import (
	"fmt"

	db "apiturnos/src/config"
	"apiturnos/src/schema/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.UserInput) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	UpdatePassword(id string, password string) (*model.User, error)
	Update(id string, input *model.UserInput) (*model.User, error)
	Deleted(id string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db.Database}
}

func (r *userRepository) FindUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	result := r.db.Where("username = ?", username).Unscoped().First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found by username")
	}
	return &user, nil
}

func (r *userRepository) CreateUser(inputUser *model.UserInput) (*model.User, error) {
	_, err := r.FindUserByUsername(inputUser.Username)
	if err == nil {
		return nil, fmt.Errorf("username already exists")
	}

	var user model.User = model.User{
		Username: inputUser.Username,
		Password: inputUser.Password,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	result := r.db.Omit("Password").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) FindById(id string) (*model.User, error) {
	var user model.User
	result := r.db.Omit("Password").Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (r *userRepository) UpdatePassword(id string, password string) (*model.User, error) {
	user, err := r.FindById(id)
	if err != nil {
		return nil, err
	}
	user.Password = password
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(id string, input *model.UserInput) (*model.User, error) {
	user, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	user.Username = input.Username
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Deleted(id string) (*model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
