package service

import (
	"os"
	"time"

	"github.com/diegofly91/apiturnos/src/model"

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
	GetUserFromToken(tokenString string) (*model.UserPayload, error)
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
		return nil, err
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

// funcion para obtener el usuario del token
func (s *authService) GetUserFromToken(tokenString string) (*model.UserPayload, error) {
	token, err := parseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	user := model.UserPayload{
		Username: claims["sub"].(string),
		Id:       int64(claims["id"].(float64)),
		Email:    claims["email"].(string),
		Role:     model.Role(claims["role"].(string)),
		// Agrega más campos según sea necesario
	}
	return &user, nil
}
