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
	r.GET("/systems/:system_id/users", configs.ViewAuthRequired(), controllers.HomeIndex)
	r.GET("/users", configs.ViewAuthRequired(), controllers.HomeIndex)
	// login controller
	r.GET("/login", configs.ViewAuthGoToHome(), controllers.LoginIndex)
	r.POST("/login", controllers.LoginSignIn)
	r.GET("/sign-out", controllers.LoginSignOut)
	// system controller
	r.GET("/apis/v1/systems", configs.APIAuthRequired(), controllers.SystemFetchAll)
	r.POST("/apis/v1/systems", configs.APIAuthRequired(), controllers.SystemCreate)
	r.PUT("/apis/v1/systems", configs.APIAuthRequired(), controllers.SystemUpdate)
	r.DELETE("/apis/v1/systems/:id", configs.APIAuthRequired(), controllers.SystemDelete)
	r.GET("/apis/v1/systems/:id/roles", configs.APIAuthRequired(), controllers.SystemFetchRoles)
	// roles controller
	r.POST("/apis/v1/roles/:system-id", configs.APIAuthRequired(), controllers.SaveRoles)
	r.GET("/apis/v1/roles/:id/permissions", configs.APIAuthRequired(), controllers.RoleFetchPermissions)
	// permissions controller
	r.POST("/apis/v1/permissions/:role-id", configs.APIAuthRequired(), controllers.SavePermissions)
	// user controller
	r.GET("/apis/v1/users", configs.APIAuthRequired(), controllers.UserFetchAll)
	r.GET("/apis/v1/users/:id", configs.APIAuthRequired(), controllers.UserFetchOne)
	//r.GET("/apis/v1/users/:id", configs.APIAuthRequired(), controllers.UserFetchOne)
	r.POST("/apis/v1/users", configs.APIAuthRequired(), controllers.UserCreate)
	//r.PUT("/apis/v1/users/:id/email", configs.APIAuthRequired(), controllers.UserUpdateEmail)
	r.PUT("/apis/v1/users/:id/password", configs.APIAuthRequired(), controllers.UserUpdatePassword)
	r.PUT("/apis/v1/users/:id/activation-key", configs.APIAuthRequired(), controllers.UserUpdateActivationKey)
	r.PUT("/apis/v1/users/:id/reset-key", configs.APIAuthRequired(), controllers.UserUpdateResetKey)
	r.PUT("/apis/v1/users/:id/activated", configs.APIAuthRequired(), controllers.UserUpdateActivated)
	//r.PUT("/apis/v1/users/:id/activated", configs.APIAuthRequired(), controllers.UserUpdateActivated)
	//r.DELETE("/apis/v1/users/:id", configs.APIAuthRequired(), controllers.UserDelete)
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
