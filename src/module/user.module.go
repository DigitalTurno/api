package module

import (
	"github.com/diegofly91/apiturnos/src/model"
	"github.com/diegofly91/apiturnos/src/repository"
	"github.com/diegofly91/apiturnos/src/service"
	"gorm.io/gorm"
)

func InitializeUserModule(db *gorm.DB) service.UserService {
	db.AutoMigrate(&model.User{})
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	return userService

}
