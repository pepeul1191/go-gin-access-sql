package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{
		"title":   "Página de Inicio",
		"mensaje": "Hola desde una plantilla!",
	})
}

func LoginSignIn(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "admin" && password == "sistema123" {
		// Autenticación exitosa: establecer sesión
		session := sessions.Default(c)
		session.Set("status", "activate")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	// Autenticación fallida
	c.Redirect(http.StatusSeeOther, "/login?error=Usuario+y/o+contraseña+incorrectos")
}

func LoginSignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("status")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
	return
}
