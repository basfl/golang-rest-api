package main

import (
	_ "example.com/rest-api/db"
	"example.com/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
