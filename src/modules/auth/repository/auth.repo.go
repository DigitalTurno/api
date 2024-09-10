package repository

import (
	db "apiturnos/src/config"
	"apiturnos/src/schema/model"
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(userId int64, token string, userAgent string, expiration time.Time) (bool, error)
	GetTokenForUserAndUserAgent(userId int64, userAgent string) (*model.UserRefreshToken, error)
	DeletedToken(userId int64, userAgend string) (bool, error)
	DeleteAllTokenByUserId(userId int64) (bool, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	return &authRepository{db.Database}
}

func (r *authRepository) Create(userId int64, token string, userAgent string, expiration time.Time) (bool, error) {
	userRefreshToken := model.UserRefreshToken{
		UserID:       userId,
		RefreshToken: token,
		UserAgent:    userAgent,
		Expiration:   expiration,
	}
	if err := r.db.Create(&userRefreshToken).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *authRepository) GetTokenForUserAndUserAgent(userId int64, userAgent string) (*model.UserRefreshToken, error) {
	userRefreshToken := model.UserRefreshToken{}
	result := r.db.Where("user_id = ? AND user_agent = ?", userId, userAgent).First(&userRefreshToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userRefreshToken, nil
}

func (r *authRepository) DeletedToken(userId int64, userAgend string) (bool, error) {
	userRefreshToken, err := r.GetTokenForUserAndUserAgent(userId, userAgend)
	if err != nil {
		return false, err
	}
	if err := r.db.Delete(&userRefreshToken).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *authRepository) DeleteAllTokenByUserId(userId int64) (bool, error) {
	if err := r.db.Where("user_id = ?", userId).Delete(&model.UserRefreshToken{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
