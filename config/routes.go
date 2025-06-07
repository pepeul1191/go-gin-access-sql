package config

import (
	"access/app/controllers" // ajusta si tu módulo tiene otro nombre

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Rutas de Home
	r.GET("/", controllers.HomeIndex)

	// Puedes seguir registrando más rutas aquí
	// r.GET("/users", controllers.UserIndex)
}
