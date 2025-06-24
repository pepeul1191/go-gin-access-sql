package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Variables globales
var (
	JWTSecretKey string
)

// LoadEnv carga las variables de entorno desde .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("La variable JWT_SECRET no está definida en .env")
	}

	JWTSecretKey = jwtSecret
}
