package configs

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ViewAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("status") != "activate" {
			c.Redirect(302, "/login?error=Debe de estar logueado")
			c.Abort()
			return
		}
		c.Next()
	}
}

func ViewAuthGoToHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("status") == "activate" {
			c.Redirect(302, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}

func APIAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("status") != "activate" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Debe de estar logueado",
				"message": "session status not activate",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ExtAPIAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		c.Next()
	}
}

func AdminAPIAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el valor del header X-Auth-Admin-Trigger
		incomingAuth := c.GetHeader("X-Auth-Admin-Trigger")

		// Obtener el valor esperado del entorno
		expectedAuth := os.Getenv("HTTP_X_AUTH_ADMIN_TRIGGER")

		// Validar que ambos valores coincidan
		if incomingAuth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Missing X-Auth-Admin-Trigger header",
			})
			c.Abort()
			return
		}

		if incomingAuth != expectedAuth {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid X-Auth-Admin-Trigger value",
			})
			c.Abort()
			return
		}

		// Si la validación es exitosa, continuar con la siguiente función
		c.Next()
	}
}

func RequireAdminJWT() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse con validación y claims
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validar el método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No se pudieron leer los claims"})
			return
		}

		// Validar que el rol sea "admin"
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Acceso restringido a administradores"})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
