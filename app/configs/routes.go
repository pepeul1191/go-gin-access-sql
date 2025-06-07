package config

import (
	"access/app/controllers" // ajusta si tu módulo tiene otro nombre

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomeIndex)
}
