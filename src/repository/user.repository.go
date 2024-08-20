package repository

import (
	"fmt"

	"github.com/diegofly91/apiturnos/src/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindUserByUsername(username string) (model.User, error)
	Update(user model.User) model.User
	Deleted(id string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindUserByUsername(username string) (model.User, error) {
	user := model.User{}
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	result := r.db.Find(&users)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) FindById(id string) (*model.User, error) {
	var user model.User
	result := r.db.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (r *userRepository) Update(user model.User) model.User {
	r.db.Save(&user)
	return user
}

func (r *userRepository) Deleted(id string) (model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
