package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserRefreshToken struct {
	ID           int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int64           `gorm:"not null;index" json:"userId"` // Configurar como clave foránea y agregar índice
	RefreshToken string          `json:"refresh_token"`
	UserAgent    string          `json:"user_agent"`
	Expiration   time.Time       `json:"expires_at"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	Deleted      *gorm.DeletedAt `gorm:"index" json:"deleted" ignore:"true"`
}

// CustomClaims define las reclamaciones personalizadas que incluyen los campos adicionales del modelo User
type UserPayload struct {
	Username   string    `json:"username"`
	Id         string    `json:"id"`
	Role       Role      `json:"role"`
	Expiration time.Time `json:"expiration"`
	jwt.RegisteredClaims
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
