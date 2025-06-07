package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error404(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusNotFound, "error.tmpl", gin.H{
			"title": "Página no encontrada",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Recurso no encontrado",
			"detail": c.Request.Method + "-" + c.Request.URL.Path,
		})
	}
}
