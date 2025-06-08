// controllers/foot_controller.go
package controllers

import (
	"access/app/configs"
	"access/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET: apis/v1/systems
func SystemFetchAll(c *gin.Context) {
	var systems []models.System
	// try to connect to DB
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudo conectar a la base de datos",
			"message": err.Error(),
		})
		return
	}
	// try to execute query
	if err := configs.DB.Find(&systems).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Error al consultar systems",
				"message": err.Error(),
			},
		)
		return
	}
	// response
	c.JSON(http.StatusOK, systems)
}

// POST: /apis/v1/systems
func SystemCreate(c *gin.Context) {
	var system models.System
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&system); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
		return
	}
	// Establecer fechas (si no se hace en frontend)
	system.Created = time.Now()
	system.Updated = time.Now()
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión fallida", "message": err.Error()})
		return
	}
	if err := configs.DB.Create(&system).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el sistema", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, system)
}

// PUT: /apis/v1/systems/:id
func SystemUpdate(c *gin.Context) {
	id := c.Param("id")
	var system models.System
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión fallida", "message": err.Error()})
		return
	}
	// Buscar el sistema existente
	if err := configs.DB.First(&system, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sistema no encontrado"})
		return
	}
	// Bind JSON sobre el sistema existente
	if err := c.ShouldBindJSON(&system); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
		return
	}
	system.Updated = time.Now()
	if err := configs.DB.Model(&system).Select("name", "description", "repository", "updated").Updates(system).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, system)
}
