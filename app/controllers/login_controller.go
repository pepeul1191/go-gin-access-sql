package controllers

import (
	"access/app/configs"
	"access/app/models"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func LoginExtSignInByUsername(c *gin.Context) {
	// Leer JSON del cuerpo
	var req models.ExtSystemUsersUsernameInput
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

	// 🔐 Generar el JWT
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &models.Claims{
		Role:     "external",
		Username: existingUser.Username,
		Email:    existingUser.Email,
		UserID:   existingUser.ID,
		SystemID: req.SystemID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tu-app",
			Audience:  []string{"clientes"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(configs.JWTSecretKey) // Mejor guárdala en variables de entorno

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al generar el token", "error": err.Error()})
		return
	}

	// armar respuesta
	if existingUser.Activated {
		var userLogged models.ExtSystemUsersOutput
		userLogged.Email = existingUser.Email
		userLogged.Username = existingUser.Username
		userLogged.SystemID = req.SystemID
		userLogged.ID = existingUser.ID
		userLogged.Token = signedToken
		// devolver datos de usuario
		c.JSON(http.StatusOK, userLogged)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cuenta de usuario no activa", "error": "activated = false"})
	}
}

func LoginExtSignInByEmail(c *gin.Context) {
	// Leer JSON del cuerpo
	var req models.ExtSystemUsersEmailInput
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
	if err := configs.DB.Where("email = ? AND password = ? AND system_id = ?", req.Email, req.Password, req.SystemID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado", "error": err.Error()})
		return
	}

	// 🔐 Generar el JWT
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &models.Claims{
		Role:     "external",
		Username: existingUser.Username,
		Email:    existingUser.Email,
		UserID:   existingUser.ID,
		SystemID: req.SystemID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tu-app",
			Audience:  []string{"clientes"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(configs.JWTSecretKey) // Mejor guárdala en variables de entorno

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al generar el token", "error": err.Error()})
		return
	}

	// armar respuesta
	if existingUser.Activated {
		var userLogged models.ExtSystemUsersOutput
		userLogged.Email = existingUser.Email
		userLogged.Username = existingUser.Username
		userLogged.SystemID = req.SystemID
		userLogged.ID = existingUser.ID
		userLogged.Token = signedToken
		// devolver datos de usuario
		c.JSON(http.StatusOK, userLogged)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cuenta de usuario no activa", "error": "activated = false"})
	}
}

func AdminSignInByHeader(c *gin.Context) {
	// 🔐 Generar el JWT
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &models.AdminClaims{
		Role:     "admin",
		Username: "admin",
		Email:    "admin@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tu-app",
			Audience:  []string{"clientes"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(configs.JWTSecretKey) // Mejor guárdala en variables de entorno

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al generar el token", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, signedToken)
}

func LoginSignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("status")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
	return
}
