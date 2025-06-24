package configs

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
