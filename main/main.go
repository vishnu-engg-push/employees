package main

import (
	"employees/config"
	"employees/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDB()
	config.ConnectToRedis()
	router := gin.Default()
	routes.Routes(router)
	router.Run(":8080")
}
