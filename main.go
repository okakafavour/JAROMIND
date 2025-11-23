package main

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/jaromind-backend/config"
	"github.com/okakafavour/jaromind-backend/router"
)

func main() {

	config.InitDatabase()
	r := gin.Default()

	router.RegisterRoutes(r)
	r.Run(":8080")

}
