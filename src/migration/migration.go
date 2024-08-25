package migration

import (
	db "github.com/diegofly91/apiturnos/src/config"
	"github.com/diegofly91/apiturnos/src/model"
)

func MigrateTable() {
	db := db.Database
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{})
}
