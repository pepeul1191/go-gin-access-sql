// controllers/foot_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Recurso no encontrado",
		"error":   "resource_not_found"})
	return
}
