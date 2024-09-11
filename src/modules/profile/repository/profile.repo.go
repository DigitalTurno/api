package repository

import (
	db "apiturnos/src/config"
	"apiturnos/src/schema/model"
	"strconv"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(userId string, input model.ProfileInput) (*model.Profile, error)
	Update(userId string, input model.ProfileInput) (*model.Profile, error)
	GetProfileByUserId(userId string) (*model.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository() ProfileRepository {
	return &profileRepository{db.Database}
}

func (r *profileRepository) Create(userId string, input model.ProfileInput) (*model.Profile, error) {
	id, _ := strconv.ParseInt(userId, 10, 64)

	profile := model.Profile{
		UserID:    id,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
	}
	if err := r.db.Create(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) Update(userId string, input model.ProfileInput) (*model.Profile, error) {

	profile := model.Profile{}
	result := r.db.Where("user_id = ?", userId).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}
	id, _ := strconv.ParseInt(userId, 10, 64)

	profile.Firstname = input.Firstname
	profile.Lastname = input.Lastname
	profile.Email = input.Email
	profile.UserID = id
	if err := r.db.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) GetProfileByUserId(userId string) (*model.Profile, error) {
	profile := model.Profile{}
	result := r.db.Where("user_id = ?", userId).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}
	return &profile, nil
}
