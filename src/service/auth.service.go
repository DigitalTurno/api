package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/diegofly91/apiturnos/src/model"
	"github.com/diegofly91/apiturnos/src/utils"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}

type AuthService interface {
	Login(username, password string) (*model.Token, error)
	JwtValidate(ctx context.Context, token string) (*jwt.Token, error)
}

type authService struct {
	user UserService
}

func NewAuthService() AuthService {
	user := NewUserService()
	return &authService{user: user}
}

func (s *authService) Login(username, password string) (*model.Token, error) {
	user, err := s.user.FindUserByUsername(username)

	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	comparedPassword := utils.ComparePassword(user.Password, password)
	if comparedPassword != nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	// Generate JWT
	token, err := generateJWT(user)
	if err != nil {
		return nil, err
	}
	return &model.Token{AccessToken: token}, nil
}

func (s *authService) ValidateToken(tokenString string) (bool, error) {
	token, err := parseJWT(tokenString)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func generateJWT(user *model.User) (string, error) {
	// Define the token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &model.UserPayload{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Token válido por 24 horas
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenString, nil
}

func parseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authService) JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &model.UserPayload{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
