package model

// Product representation of a product
type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
}
