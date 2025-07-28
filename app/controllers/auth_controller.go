package controllers

import (
	"access/app/configs"
	"access/app/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

func ViewToken(c *gin.Context) {
	tokenString := c.GetHeader("token") // o "Authorization" si usas Bearer
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
		return
	}

	// Parsear y validar el token
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron leer los claims"})
		return
	}

	c.JSON(http.StatusOK, claims)
}
