package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crea una nueva instancia del router
	router := gin.Default()

	// Ruta GET simple
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola desde Gin!",
		})
	})

	// Ruta con parámetro
	router.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.JSON(http.StatusOK, gin.H{
			"saludo": "Hola, " + nombre,
		})
	})

	// Ruta POST con JSON
	router.POST("/datos", func(c *gin.Context) {
		var datos map[string]interface{}
		if err := c.BindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"recibido": datos,
		})
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")
}
