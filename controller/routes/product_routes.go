package routes

import (
	"go-api/controller/controllers"

	"github.com/gin-gonic/gin"
)

func registerProductRoutes(server *gin.Engine, controller controllers.ProductController) {
	products := server.Group("/product")
	{
		products.GET("/", controller.GetProducts)
		products.POST("/create", controller.CreateProduct)
		products.GET("/:productID", controller.GetProductByID)
	}
}
