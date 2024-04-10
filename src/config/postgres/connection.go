package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB contiene la conexión a la base de datos
var DB *gorm.DB

// Connect permite realizar conexión a la base de datos de postgres
func Connect() {
	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var err error

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("database connected")
}
