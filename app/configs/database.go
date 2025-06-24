package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("La variable DATABASE_URL no está definida en .env")
	}

	var err2 error
	DB, err2 = gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err2 != nil {
		return err2
	}
	return nil
}
