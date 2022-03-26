package models

type Product struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}
