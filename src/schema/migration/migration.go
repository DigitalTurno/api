package migration

import (
	db "apiturnos/src/config"
	"apiturnos/src/schema/model"
)

func MigrateTable() {
	db := db.Database
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{})
}
