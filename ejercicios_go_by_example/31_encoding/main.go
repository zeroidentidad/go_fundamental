package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {

	message := "hello,go (*w3hu%#"
	demoBase64(message)
	demoHex(message)
	demoJson()
	demoXml()
	demoCsv()
}

func demoBase64(message string) {

	fmt.Println("--------Demo encoding base64--------")

	fmt.Println("plaintext:")
	fmt.Println(message)

	encoding := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("base64 msg:")
	fmt.Println(encoding)

	decoding, _ := base64.StdEncoding.DecodeString(encoding)
	fmt.Println("decoding base64 msg:")
	fmt.Println(string(decoding))

}

func demoHex(message string) {

	fmt.Println("--------Demo encoding Hex--------")

	fmt.Println("plaintext:")
	fmt.Println(message)

	encoding := hex.EncodeToString([]byte(message))
	fmt.Println("Hex msg:")
	fmt.Println(encoding)

	decoding, _ := hex.DecodeString(encoding)
	fmt.Println("decoding Hex msg:")
	fmt.Println(string(decoding))
}

func demoJson() {

	fmt.Println("--------Demo encoding json--------")
	type Employee struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// struct to json
	fmt.Println(">>>>struct to json....")
	emp := &Employee{Id: "12345", Name: "Michael", Email: "michael@email.com"}
	b, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// json string to struct
	fmt.Println(">>>>json to struct....")
	var newEmp Employee
	str := `{"Id":"4566","Name":"Brown","Email":"brown@email.com"}`
	json.Unmarshal([]byte(str), &newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)
}

func demoXml() {

	fmt.Println("--------Demo encoding xml--------")
	type EmployeeCountry struct {
		CountryCode string `xml:"code,attr"` // XML attribute: code
		CountryName string `xml:",chardata"` // XML value
	}
	type Employee struct {
		XMLName xml.Name        `xml:"employee"`
		Id      string          `xml:"id"`
		Name    string          `xml:"name"`
		Email   string          `xml:"email"`
		Country EmployeeCountry `xml:"country"`
	}

	// struct to xml
	fmt.Println(">>>>struct to xml....")
	emp := &Employee{Id: "12345", Name: "Michael", Email: "michael@email.com",
		Country: EmployeeCountry{CountryCode: "DE", CountryName: "Germany"}}
	b, err := xml.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// xml string to struct
	fmt.Println(">>>>xml to struct....")
	var newEmp Employee
	str := `<employee><id>555</id><name>Lee</name><email>lee@email.com</email><country code="CN">China</country></employee>`
	xml.Unmarshal([]byte(str), &newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)
	fmt.Printf("CountryCode: %s\n", newEmp.Country.CountryCode)
	fmt.Printf("CountryName: %s\n", newEmp.Country.CountryName)
}

func demoCsv() {

	fmt.Println("--------Demo encoding csv--------")
	type Employee struct {
		Name    string
		Email   string
		Country string
	}

	// read csv file to a array of struct
	fmt.Println(">>>>read a csv file and load to array of struct....")
	file, err := os.Open("./demo.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, item := range csvData {
		emp.Name = item[0]
		emp.Email = item[1]
		emp.Country = item[2]
		employees = append(employees, emp)
		fmt.Printf("name: %s email: %s  country: %s\n", item[0], item[1], item[2])
	}
	fmt.Println(">>>>show all employee structs....")
	fmt.Println(employees)

	// write data to csv file
	fmt.Println(">>>>write data to a csv file ....")
	csvFile, err := os.Create("./employee.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	for _, itemEmp := range employees {
		records := []string{itemEmp.Name, itemEmp.Email, itemEmp.Country}
		err := writer.Write(records)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
	fmt.Println(">>>>Done")

}
