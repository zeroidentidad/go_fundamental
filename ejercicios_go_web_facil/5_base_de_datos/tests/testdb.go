package main

import (
	"../models"
)

func main() {
	models.CreateConnection()
	//result := models.ExistsTable("users")
	//fmt.Println(result)
	//models.Ping()
	//models.CreateTables()

	//user := models.CreateUser("Jesus4", "xd4", "test2@mail.com")
	//fmt.Println(user)

	userdel := models.User{ID: 1}
	userdel.Delete()

	models.CloseConnection()
}
