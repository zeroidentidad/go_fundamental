package main

import (
	"fmt"

	"../models"
)

func main() {
	models.CreateConnection()
	//result := models.ExistsTable("users")
	//fmt.Println(result)
	//models.Ping()
	models.CreateTables()

	user := models.NewUser("Jesus", "xd")
	fmt.Println(user)
	user.Save()

	models.CloseConnection()
}
