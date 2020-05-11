package customer

type Customer struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Address       string `json:"address"`
	BusinessPhone string `json:"businessphone"`
	City          string `json:"city"`
	Company       string `json:"company"`
}

type CustomerList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int64       `json:"totalRecords"`
}
