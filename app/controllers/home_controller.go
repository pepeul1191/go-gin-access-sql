// controllers/foot_controller.go
package controllers

import (
	"access/app/configs"
	"access/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.HTML(200, "home.tmpl", gin.H{
		"title":   "Página de Inicio",
		"mensaje": "Hola desde una plantilla!",
	})
}

func FetchFoots(c *gin.Context) {
	var foots []models.Foot
	// try to connect to DB
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudo conectar a la base de datos",
			"message": err.Error(),
		})
		return
	}
	// try to execute query
	if err := configs.DB.Find(&foots).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Error al consultar foots",
				"message": err.Error(),
			},
		)
		return
	}
	// response
	c.JSON(http.StatusOK, foots)
}
