package main

import (
	"access/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// set routes
	config.SetupRoutes(r)
	r.Run(":8080")
}
