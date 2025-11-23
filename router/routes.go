package router

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/jaromind-backend/controllers"
	"github.com/okakafavour/jaromind-backend/middleware"
)

func RegisterRoutes(router *gin.Engine) {

	//--- Auth routes ---
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// ----- protected routes (JWT required) -----
	protected := router.Group("/user")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", controllers.GetProfile)
	}
}
