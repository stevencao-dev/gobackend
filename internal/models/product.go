package models

// json: tells go the encoding to use when sending to the frontend
type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}

