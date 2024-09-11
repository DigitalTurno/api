package service

import (
	"apiturnos/src/modules/profile/repository"
	"apiturnos/src/schema/model"
)

type ProfileService interface {
	CreateProfile(userId string, input model.ProfileInput) (*model.Profile, error)
	UpdateProfile(userId string, input model.ProfileInput) (*model.Profile, error)
	GetProfileByUserId(userId string) (*model.Profile, error)
}

type profileService struct {
	repo repository.ProfileRepository
}

func NewProfileService() ProfileService {
	repo := repository.NewProfileRepository()
	return &profileService{repo}
}

func (s *profileService) CreateProfile(userId string, input model.ProfileInput) (*model.Profile, error) {
	return s.repo.Create(userId, input)
}

func (s *profileService) UpdateProfile(userId string, input model.ProfileInput) (*model.Profile, error) {
	return s.repo.Update(userId, input)
}

func (s *profileService) GetProfileByUserId(userId string) (*model.Profile, error) {
	return s.repo.GetProfileByUserId(userId)
}
