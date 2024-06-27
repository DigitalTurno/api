package repository

import (
	"github.com/diegofly91/apiturnos/src/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindAll() []model.User
	FindById(id string) (model.User, error)
	FindUserByUsername(username string) (model.User, error)
	Update(user model.User) model.User
	Deleted(id string) model.User
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

func (r *userRepository) FindAll() []model.User {
	users := []model.User{}
	r.db.Find(&users)
	return users
}

func (r *userRepository) FindById(id string) (model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(user model.User) model.User {
	r.db.Save(&user)
	return user
}

func (r *userRepository) Deleted(id string) model.User {
	user := model.User{}
	r.db.First(&user, id)
	r.db.Delete(&user)
	return user
}
