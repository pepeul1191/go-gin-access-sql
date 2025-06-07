package controllers

import (
	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.HTML(200, "home.tmpl", gin.H{
		"title":   "Página de Inicio",
		"mensaje": "Hola desde una plantilla!",
	})
}
