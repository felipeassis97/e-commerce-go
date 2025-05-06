package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

// NewProductUsecase returns a new productUsecase
func NewProductUsecase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

// GetProducts returns a list of products from the Repository
func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (int64, error) {
	return pu.repository.CreateProduct(product)
}

func (pu *ProductUseCase) GetProductByID(ID int64) (*model.Product, error) {
	return pu.repository.GetProductByID(ID)
}
