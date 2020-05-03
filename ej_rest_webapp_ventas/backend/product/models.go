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

type ProductTop struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	Vendidos    float64 `json:"vendidos"`
}

type ProductTopList struct {
	Data        []*ProductTop `json:"data"`
	TotalVentas float64       `json:"totalVentas"`
}
