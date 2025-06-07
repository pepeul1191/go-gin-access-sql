package main

import (
	config "access/app/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// settings
	config.SetupRoutes(r)
	r.LoadHTMLGlob("views/*")
	r.Run(":8080")
}
