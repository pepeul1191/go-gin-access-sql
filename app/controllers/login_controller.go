package controllers

import (
	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{
		"title":   "Página de Inicio",
		"mensaje": "Hola desde una plantilla!",
	})
}
