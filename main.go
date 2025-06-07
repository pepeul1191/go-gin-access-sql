package main

import (
	"access/app/controllers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomeIndex)
	r.GET("/login", controllers.LoginIndex)
	r.GET("/apis/v1/foots", controllers.FetchFoots)
	r.NoRoute(controllers.Error404)
}

func main() {
	r := gin.Default()
	// settings
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./public")
	// load routes
	setupRoutes(r)
	r.Run(":8080")
}
