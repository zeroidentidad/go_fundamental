package main

import (
	"../models"
)

func main() {
	models.CreateConnection()
	//result := models.ExistsTable("users")
	//fmt.Println(result)
	//models.Ping()
	models.CreateTables()
	models.CloseConnection()
}
