package controllers

import (
	"access/app/configs"
	"access/app/models"
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

func LoginExtSignIn(c *gin.Context) {
	// Leer JSON del cuerpo
	var req models.ExtSystemUsersInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No se pueden parsear los datos enviados", "error": err.Error()})
		return
	}

	// Conexión a la base de datos
	if configs.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión DB no inicializada"})
		return
	}

	// Verificar si el usuario esta registrado en sistema
	var existingUser models.SystemUserView
	if err := configs.DB.Where("username = ? AND password = ? AND system_id = ?", req.Username, req.Password, req.SystemID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado", "error": err.Error()})
		return
	}
	// armar respuesta
	if existingUser.Activated {
		var userLogged models.ExtSystemUsersOutput
		userLogged.Email = existingUser.Email
		userLogged.Username = existingUser.Username
		userLogged.SystemID = req.SystemID
		userLogged.ID = existingUser.ID
		// devolver datos de usuario
		c.JSON(http.StatusOK, userLogged)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cuenta de usuario no activa", "error": "activated = false"})
	}
}

func LoginSignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("status")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
	return
}
