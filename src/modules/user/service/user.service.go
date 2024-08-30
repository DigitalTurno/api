package service

import (
	"apiturnos/src/modules/user/repository"
	"apiturnos/src/schema/model"
)

type UserService interface {
	CreateUser(user *model.UserInput) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	Update(id string, input *model.UserInput) (*model.User, error)
	UpdatePassword(id string, password string) (*model.User, error)
	DeleteUser(userId string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService() UserService {
	repo := repository.NewUserRepository()
	return &userService{repo}
}

func (s *userService) FindUserByUsername(username string) (*model.User, error) {
	return s.repo.FindUserByUsername(username)
}

func (s *userService) CreateUser(user *model.UserInput) (*model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) FindAll() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) FindById(id string) (*model.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) Update(id string, input *model.UserInput) (*model.User, error) {
	return s.repo.Update(id, input)
}

func (s *userService) DeleteUser(userId string) (*model.User, error) {
	return s.repo.Deleted(userId)
}

func (s *userService) UpdatePassword(id string, password string) (*model.User, error) {
	return s.repo.UpdatePassword(id, password)
}
