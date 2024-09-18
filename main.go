package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.UserRoutes(r)
	routes.ItemRoutes(r)
	routes.TransactionRoutes(r)

	r.Run(":8080")
}
