package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting products",
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing request body",
		})
		return
	}
	ID, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating product",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id":      ID,
		"message": "Created product",
	})
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")
	if id == "" {
		response := model.Response{
			Type:    model.MissingParams,
			Message: "Invalid product ID",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Type:    model.MissingParams,
			Message: "Invalid product ID",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductByID(int64(productID))
	if err != nil {
		response := model.Response{
			Type:    model.DatabaseError,
			Message: fmt.Sprintf("Error getting product ID: %s", id),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := model.Response{
			Type:    model.EmptyResponse,
			Message: fmt.Sprintf("Product ID: %s not found", id),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
