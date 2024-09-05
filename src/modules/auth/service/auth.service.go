package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"apiturnos/src/constants"
	userService "apiturnos/src/modules/user/service"
	"apiturnos/src/schema/model"
	"apiturnos/src/utils"

	"github.com/golang-jwt/jwt/v5"
)

type Secrets struct {
	JWTSecret                  string
	JWTExpirationSecret        string
	JWTSecretRefresh           string
	JWTExpirationSecretRefresh string
	JWTSecretEmail             string
	JWTExpirationSecretEmail   string
}

var secrets = getJwtSecret()

func getJwtSecret() Secrets {
	secrets := Secrets{
		JWTSecret:                  os.Getenv("JWT_SECRET"),
		JWTExpirationSecret:        os.Getenv("JWT_EXPIRATION_SECRET"),
		JWTSecretRefresh:           os.Getenv("JWT_REFRESH_SECRET"),
		JWTExpirationSecretRefresh: os.Getenv("JWT_EXPIRATION_SECRET_REFRESH"),
		JWTSecretEmail:             os.Getenv("JWT_SECRET_EMAIL"),
		JWTExpirationSecretEmail:   os.Getenv("JWT_EXPIRATION_SECRET_EMAIL"),
	}
	return secrets
}

type AuthService interface {
	Login(username, password string) (*model.Token, error)
	JwtValidate(ctx context.Context, token string, typeToken constants.TypeToken) (*jwt.Token, error)
	GenerateJWT(user *model.UserPayload) (string, error)
}

type authService struct {
	user userService.UserService
}

func NewAuthService() AuthService {
	user := userService.NewUserService()
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
	AccessToken, err := generateJWT(user, secrets.JWTExpirationSecret, secrets.JWTSecret)

	if err != nil {
		return nil, err
	}
	RefreshToken, err := generateJWT(user, secrets.JWTExpirationSecretRefresh, secrets.JWTSecretRefresh)
	if err != nil {
		return nil, err
	}

	return &model.Token{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
	}, nil
}

func generateJWT(user *model.User, expiration string, secret string) (string, error) {
	// Define the token expiration time
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		return "", fmt.Errorf("error al parsear duración: %v", err)
	}
	expirationTime := time.Now().Add(duration)
	// Create the JWT claims, which includes the username and expiry time
	claims := &model.UserPayload{
		Id:         strconv.FormatInt(user.ID, 10),
		Username:   user.Username,
		Role:       user.Role,
		Expiration: expirationTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), //
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenString, nil
}

func (s *authService) JwtValidate(ctx context.Context, token string, typeToken constants.TypeToken) (*jwt.Token, error) {
	secret := ""
	switch typeToken {
	case constants.TypeToken("TOKEN"):
		secret = secrets.JWTSecret
	case constants.TypeToken("REFRESH"):
		secret = secrets.JWTSecretRefresh
	case constants.TypeToken("EMAIL_PASSWORD"):
		secret = secrets.JWTSecretEmail
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &model.UserPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenClaims, nil
}

func (s *authService) GenerateJWT(userPayload *model.UserPayload) (string, error) {
	duration, err := time.ParseDuration(secrets.JWTExpirationSecret)
	if err != nil {
		return "", fmt.Errorf("error al parsear duración: %v", err)
	}
	expirationTime := time.Now().Add(duration)

	claims := &model.UserPayload{
		Username: userPayload.Username,
		Id:       userPayload.Id,
		Role:     userPayload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secrets.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
