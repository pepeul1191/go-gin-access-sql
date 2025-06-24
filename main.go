package main

import (
	"access/app/configs"
	"access/app/controllers"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	// Cargar .env al iniciar la app
	configs.LoadEnv()
	// home controller
	r.GET("/", configs.ViewAuthRequired(), controllers.HomeIndex)
	r.GET("/systems", configs.ViewAuthRequired(), controllers.HomeIndex)
	r.GET("/systems/:system_id/users", configs.ViewAuthRequired(), controllers.HomeIndex)
	r.GET("/users", configs.ViewAuthRequired(), controllers.HomeIndex)
	// login controller
	r.GET("/login", configs.ViewAuthGoToHome(), controllers.LoginIndex)
	r.POST("/login", controllers.LoginSignIn)
	r.GET("/sign-out", controllers.LoginSignOut)
	r.POST("/api/v1/sign-in", configs.ExtAPIAuthRequired(), controllers.LoginExtSignIn)
	// system controller
	r.GET("/api/v1/systems", configs.APIAuthRequired(), controllers.SystemFetchAll)
	r.POST("/api/v1/systems", configs.APIAuthRequired(), controllers.SystemCreate)
	r.PUT("/api/v1/systems", configs.APIAuthRequired(), controllers.SystemUpdate)
	r.DELETE("/api/v1/systems/:id", configs.APIAuthRequired(), controllers.SystemDelete)
	r.GET("/api/v1/systems/:id/roles", configs.APIAuthRequired(), controllers.SystemFetchRoles)
	r.GET("/api/v1/systems/:id/users", configs.APIAuthRequired(), controllers.SystemFetchUsers)
	r.POST("/api/v1/systems/:id/users", configs.APIAuthRequired(), controllers.SystemSaveUsers)
	r.POST("/api/v1/systems-permissions/:system_id/users/:user_id", configs.APIAuthRequired(), controllers.SystemSavePermissionsUsers)
	r.GET("/api/v1/systems-permissions/:system_id/users/:user_id/roles/:role_id", configs.APIAuthRequired(), controllers.SystemPermissionFetchUsers)
	// roles controller
	r.POST("/api/v1/roles/:system-id", configs.APIAuthRequired(), controllers.SaveRoles)
	r.GET("/api/v1/roles/:id/permissions", configs.APIAuthRequired(), controllers.RoleFetchPermissions)
	// permissions controller
	r.POST("/api/v1/permissions/:role-id", configs.APIAuthRequired(), controllers.SavePermissions)
	// user controller
	r.GET("/api/v1/users", configs.APIAuthRequired(), controllers.UserFetchAll)
	r.GET("/api/v1/users/:id", configs.APIAuthRequired(), controllers.UserFetchOne)
	r.POST("/api/v1/users", configs.APIAuthRequired(), controllers.UserCreate)
	r.PUT("/api/v1/users/:id", configs.APIAuthRequired(), controllers.UserUpdate)
	r.PUT("/api/v1/users/:id/password", configs.APIAuthRequired(), controllers.UserUpdatePassword)
	r.PUT("/api/v1/users/:id/activation-key", configs.APIAuthRequired(), controllers.UserUpdateActivationKey)
	r.PUT("/api/v1/users/:id/reset-key", configs.APIAuthRequired(), controllers.UserUpdateResetKey)
	r.PUT("/api/v1/users/:id/activated", configs.APIAuthRequired(), controllers.UserUpdateActivated)
	//r.DELETE("/api/v1/users/:id", configs.APIAuthRequired(), controllers.UserDelete) TODO
	// error controller
	r.NoRoute(controllers.Error404)
}

func main() {
	r := gin.Default()
	// load db
	if err := configs.ConnectToDB(); err != nil {
		log.Fatal("Error al iniciar DB:", err)
	}
	// settings
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./public")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	// load routes
	setupRoutes(r)
	r.Run(":8080")
}
