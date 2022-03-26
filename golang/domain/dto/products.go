package dto

type QueryParams struct {
	Category      *string
	PriceLessThan *int
}

type PriceResp struct {
	Original           int     `json:"original"`
	Final              float32 `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency" default:"EUR"`
}

type ProductResp struct {
	Sku      string    `json:"sku"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Price    PriceResp `json:"price"`
}
