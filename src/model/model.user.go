package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(100);not null;unique" json:"username" validate:"required,min=5,max=100"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=8,max=100"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
}

type UserInput struct {
	Username string `json:"username"  validate:"required,min=5,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Email    string `json:"email" validate:"required,email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
