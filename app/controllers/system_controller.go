// controllers/foot_controller.go
package controllers

import (
	"access/app/configs"
	"access/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
