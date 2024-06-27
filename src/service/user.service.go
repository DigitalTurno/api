package service

import (
	"github.com/diegofly91/apiturnos/src/model"
	"github.com/diegofly91/apiturnos/src/repository"
)

type UserService interface {
	CreateUser(user model.User) (model.User, error)
	FindAll() []model.User
	FindById(id string) (model.User, error)
	FindUserByUsername(username string) (model.User, error)
	Update(user model.User) model.User
	Delete(userId string) model.User
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) FindUserByUsername(username string) (model.User, error) {
	return s.repo.FindUserByUsername(username)
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) FindAll() []model.User {
	return s.repo.FindAll()
}

func (s *userService) FindById(id string) (model.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) Update(user model.User) model.User {
	return s.repo.Update(user)
}

func (s *userService) Delete(userId string) model.User {
	return s.repo.Deleted(userId)
}
