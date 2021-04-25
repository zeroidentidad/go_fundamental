package main

import (
	"fmt"
	"time"
)

// definir un struct
type Employee struct {
	id      int
	name    string
	country string
	created time.Time
}

func main() {

	// declarar variables
	var emp Employee
	newEmp := new(Employee)

	// establecer valores
	emp.id = 2
	emp.name = "Jesus 2"
	emp.country = "MX"
	emp.created = time.Now()

	newEmp.id = 5
	newEmp.name = "Joshua 5"
	newEmp.country = "USA"
	newEmp.created = time.Now()

	// mostrar
	fmt.Println("=====================")
	fmt.Println(emp.id)
	fmt.Println(emp.name)
	fmt.Println(emp.country)
	fmt.Println(emp.created)

	fmt.Println("=====================")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.name)
	fmt.Println(newEmp.country)
	fmt.Println(newEmp.created)
}
