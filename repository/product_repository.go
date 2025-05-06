package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
)

// ProductRepository represents a product repository
type ProductRepository struct {
	connection *sql.DB
}

// NewProductRepository returns a new product repository
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Printf("Error GetProducts: %s", err)
		return []model.Product{}, err
	}

	var products []model.Product
	var productObj model.Product
	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Product, &productObj.Price)
		if err != nil {
			fmt.Printf("Error Scan: %s", err)
		}
		products = append(products, productObj)
	}
	err = rows.Close()
	if err != nil {
		fmt.Printf("Error Close iteration: %s", err)
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int64, error) {
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"
	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Printf("Error Prepare: %s", err)
		return 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(product.Product, product.Price).Scan(&id)
	if err != nil {
		fmt.Printf("Error while executing query: %s", err)
		return 0, err
	}
	return id, nil
}

func (pr *ProductRepository) GetProductByID(ID int64) (*model.Product, error) {
	query := "SELECT id, product_name, price FROM product WHERE id = $1"
	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Printf("Error Prepare: %s", err)
		return nil, err
	}
	defer stmt.Close()
	var productObj model.Product
	err = stmt.QueryRow(ID).Scan(&productObj.ID, &productObj.Product, &productObj.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Printf("Error while executing query: %s", err)
		return nil, err
	}

	return &productObj, nil
}
