package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	transactions := r.Group("/transactions")
	{
		transactions.GET("/:id", middlewares.AuthMiddleware(), controllers.GetTransactionByID)
		transactions.GET("/user/:user_id", middlewares.AuthMiddleware(), controllers.GetTransactionsByUserID)
		transactions.POST("/", middlewares.AuthMiddleware(), controllers.CreateTransaction)
		transactions.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteTransaction)
	}
}
