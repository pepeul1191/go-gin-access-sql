package main

import (
	"access/app/configs"
	"access/app/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	// home controller
	r.GET("/", configs.ViewAuthRequired(), controllers.HomeIndex)
	r.GET("/systems", configs.ViewAuthRequired(), controllers.HomeIndex)
	r.GET("/users", configs.ViewAuthRequired(), controllers.HomeIndex)
	// login controller
	r.GET("/login", configs.ViewAuthGoToHome(), controllers.LoginIndex)
	r.POST("/login", controllers.LoginSignIn)
	r.GET("/sign-out", controllers.LoginSignOut)
	// system controller
	r.GET("/apis/v1/systems", configs.APIAuthRequired(), controllers.SystemFetchAll)
	r.POST("/apis/v1/systems", configs.APIAuthRequired(), controllers.SystemCreate)
	r.PUT("/apis/v1/systems", configs.APIAuthRequired(), controllers.SystemUpdate)
	// error controller
	r.NoRoute(controllers.Error404)
}

func main() {
	r := gin.Default()
	// settings
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./public")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	// load routes
	setupRoutes(r)
	r.Run(":8080")
}
