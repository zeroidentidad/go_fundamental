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

	user := models.NewUser("Jesus3", "xd3", "test@mail.com")
	user.Save()
	fmt.Println(user)

	models.CloseConnection()
}
