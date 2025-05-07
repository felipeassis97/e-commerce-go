package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
	"log"
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
	query := "SELECT id, product_name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Printf("Error GetProducts: %s", err)
		return []model.Product{}, err
	}

	var products []model.Product
	var productObj model.Product
	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
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

func (pr *ProductRepository) CreateProduct(product model.Product) (*model.Product, error) {
	query := "INSERT INTO products (product_name, price) VALUES ($1, $2) RETURNING id, product_name, price"
	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Printf("Error Prepare: %s", err)
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	var productObj model.Product
	err = stmt.QueryRow(product.Name, product.Price).Scan(&productObj.ID, &productObj.Name, &productObj.Price)
	if err != nil {
		fmt.Printf("Error while executing query: %s", err)
		return nil, err
	}
	return &productObj, nil
}

func (pr *ProductRepository) GetProductByID(ID int64) (*model.Product, error) {
	query := "SELECT id, product_name, price FROM products WHERE id = $1"
	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Printf("Error Prepare: %s", err)
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	var productObj model.Product
	err = stmt.QueryRow(ID).Scan(&productObj.ID, &productObj.Name, &productObj.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Printf("Error while executing query: %s", err)
		return nil, err
	}

	return &productObj, nil
}
