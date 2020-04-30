package product

type Product struct {
	ID          int     `json:"id"`
	ProductCode string  `json:"productCode"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	StandarCost float64 `json:"standarCost"`
	ListPrice   float64 `json:"listPrice"`
	Category    string  `json:"category"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}
