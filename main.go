package main

import (
	"access/app/configs"
	"access/app/controllers"
	"log"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	// Cargar .env al iniciar la app
	configs.LoadEnv()
	// home controller
	r.GET("/", configs.ViewAuthRequired(), controllers.HomeIndex)
	// login controller
	r.POST("/api/v1/sign-in/by-username", configs.ExtAPIAuthRequired(), controllers.LoginExtSignInByUsername)
	r.POST("/api/v1/sign-in/by-email", configs.ExtAPIAuthRequired(), controllers.LoginExtSignInByEmail)
	r.POST("/api/v1/sign-in/admin", configs.AdminAPIAuthRequired(), controllers.AdminSignInByHeader)
	// system controller
	r.GET("/api/v1/systems", configs.RequireAdminJWT(), controllers.SystemFetchAll)
	r.POST("/api/v1/systems", configs.RequireAdminJWT(), controllers.SystemCreate)
	r.PUT("/api/v1/systems", configs.RequireAdminJWT(), controllers.SystemUpdate)
	r.DELETE("/api/v1/systems/:id", configs.RequireAdminJWT(), controllers.SystemDelete)
	r.GET("/api/v1/systems/:id/roles", configs.RequireAdminJWT(), controllers.SystemFetchRoles)
	r.GET("/api/v1/systems/:id/users", configs.RequireAdminJWT(), controllers.SystemFetchUsers)
	r.POST("/api/v1/systems/:id/users", configs.RequireAdminJWT(), controllers.SystemSaveUsers)
	r.POST("/api/v1/systems-permissions/:system_id/users/:user_id", configs.RequireAdminJWT(), controllers.SystemSavePermissionsUsers)
	r.GET("/api/v1/systems-permissions/:system_id/users/:user_id/roles/:role_id", configs.RequireAdminJWT(), controllers.SystemPermissionFetchUsers)
	// roles controller
	r.POST("/api/v1/roles/:system-id", configs.RequireAdminJWT(), controllers.SaveRoles)
	r.GET("/api/v1/roles/:id/permissions", configs.RequireAdminJWT(), controllers.RoleFetchPermissions)
	// permissions controller
	r.POST("/api/v1/permissions/:role-id", configs.RequireAdminJWT(), controllers.SavePermissions)
	// user controller
	r.GET("/api/v1/users", configs.RequireAdminJWT(), controllers.UserFetchAll)
	r.GET("/api/v1/users/:id", configs.RequireAdminJWT(), controllers.UserFetchOne)
	r.POST("/api/v1/users", configs.RequireAdminJWT(), controllers.UserCreate)
	r.PUT("/api/v1/users/:id", configs.RequireAdminJWT(), controllers.UserUpdate)
	r.PUT("/api/v1/users/:id/password", configs.RequireAdminJWT(), controllers.UserUpdatePassword)
	r.PUT("/api/v1/users/:id/activation-key", configs.RequireAdminJWT(), controllers.UserUpdateActivationKey)
	r.PUT("/api/v1/users/:id/reset-key", configs.RequireAdminJWT(), controllers.UserUpdateResetKey)
	r.PUT("/api/v1/users/:id/activated", configs.RequireAdminJWT(), controllers.UserUpdateActivated)
	//r.DELETE("/api/v1/users/:id", configs.RequireAdminJWT(), controllers.UserDelete) TODO
	// error controller
	r.NoRoute(controllers.Error404)
}

func main() {
	r := gin.Default()
	// before all
	r.Use(func(c *gin.Context) {
		c.Header("X-Powered-By", "Gin")
		c.Header("Server", "Ubuntu")
		c.Next()
	})
	// cors
	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "https://tudominio.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// load db
	if err := configs.ConnectToDB(); err != nil {
		log.Fatal("Error al iniciar DB:", err)
	}
	// load routes
	setupRoutes(r)
	r.Run(":8080")
}
