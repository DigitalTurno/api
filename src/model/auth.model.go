package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims define las reclamaciones personalizadas que incluyen los campos adicionales del modelo User
type CustomClaims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
