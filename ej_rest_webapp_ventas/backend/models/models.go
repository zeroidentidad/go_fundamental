package models

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}
