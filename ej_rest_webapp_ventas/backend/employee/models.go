package employee

type Employee struct {
	ID            int    `json:"id"`
	LastName      string `json:"lastName"`
	FirstName     string `json:"firstName"`
	Company       string `json:"company"`
	EmailAddress  string `json:"emailAddress"`
	JobTitle      string `json:"jobTitle"`
	BusinessPhone string `json:"businessPhone"`
	HomePhone     string `json:"homePhone"`
	MobilePhone   string `json:"mobilePhone"`
	FaxNumber     string `json:"faxNumber"`
	Address       string `json:"address"`
}

type EmployeeList struct {
	Data         []*Employee `json:"data"`
	TotalRecords int64       `json:"totalRecords"`
}
