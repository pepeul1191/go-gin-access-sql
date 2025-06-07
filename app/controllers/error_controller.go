package controllers

import (
	"github.com/gin-gonic/gin"
)

func Error404(c *gin.Context) {
	c.HTML(404, "error.tmpl", gin.H{})
}
