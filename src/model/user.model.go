package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string
type Status string

const (
	Active    Status = "ACTIVE"
	Inactive  Status = "INACTIVE"
	Preactive Status = "PREACTIVE"
)
const (
	Admin     Role = "ADMIN"
	SuperUser Role = "SUPERUSER"
	Advisor   Role = "ADVISER"
	Guest     Role = "GUEST"
)

type User struct {
	ID        int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string          `gorm:"type:varchar(100);not null;unique" json:"username" validate:"required,min=5,max=100"`
	Password  string          `gorm:"type:varchar(100);not null" json:"-" validate:"required,min=8,max=100"`
	Email     string          `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	Role      Role            `gorm:"type:enum('ADMIN', 'SUPERUSER', 'ADVISER', 'GUEST');default:'GUEST'" json:"role"`
	Status    Status          `gorm:"type:enum('ACTIVE', 'INACTIVE', 'PREACTIVE');default:'PREACTIVE'" json:"status"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Deleted   *gorm.DeletedAt `gorm:"index" json:"deleted" ignore:"true"`
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
