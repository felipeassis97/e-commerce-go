package controllers

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting product",
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		response := model.Response{
			Type:    model.InvalidBody,
			Message: "Error parsing request body",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	createdProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		response := model.Response{
			Type:    model.InternalError,
			Message: "Error while creating product",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
		"product": createdProduct,
	})
}

func (p *ProductController) GetProductByID(ctx *gin.Context) {
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
			Type:    model.InternalError,
			Message: fmt.Sprintf("Error getting product ID: %s", id),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := model.Response{
			Type:    model.EmptyResponse,
			Message: fmt.Sprintf("Name ID: %s not found", id),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
