package main

import (
	"fmt"

	"../models"
)

func main() {
	models.CreateConnection()
	result := models.ExistsTable("users")
	fmt.Println(result)
	models.Ping()
	models.CloseConnection()
}
