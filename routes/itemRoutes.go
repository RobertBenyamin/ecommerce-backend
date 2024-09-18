package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(r *gin.Engine) {
	items := r.Group("/items")
	{
		items.GET("/", controllers.GetAllItems)
		items.GET("/:id", controllers.GetItemByID)

		items.POST("/", middlewares.AuthMiddleware(), controllers.CreateItem)
		items.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateItem)
		items.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteItem)
	}
}
