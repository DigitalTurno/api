package module

import (
	db "github.com/diegofly91/apiturnos/src/config"
	"github.com/diegofly91/apiturnos/src/model"
	"github.com/diegofly91/apiturnos/src/repository"
	"github.com/diegofly91/apiturnos/src/service"
)

func InitializeUserModule() service.UserService {
	db := db.Database
	db.AutoMigrate(&model.User{})
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	return userService

}
