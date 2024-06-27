package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dns string
var Database = func() (db *gorm.DB) {
	// CARGAR VARIABLES DE ENTORNO
	//dns := "root:password@/db_contacts"
	fmt.Println("Cargando variables de entorno")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	if db, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexion a la base de datos", err)
		panic(err)
	} else {
		fmt.Println("Conexion a la base de datos exitosa")
		return db
	}
}()
