package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	conn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ProductRepository := repository.NewProductRepository(conn)
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.POST("/product", productController.CreateProduct)

	server.GET("/product/:productID", productController.GetProductByID)

	server.Run(":8000")
}
