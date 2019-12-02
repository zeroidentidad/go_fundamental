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

	//models.CreateTables()

	//user := models.CreateUser("Jesus4", "xd4", "test2@mail.com")
	//fmt.Println(user)

	//userdel := models.User{ID: 1}
	//userdel.Delete()

	//user := models.GetUser(3)
	//fmt.Println(user)

	users := models.GetUsers()
	fmt.Println(users)

	models.CloseConnection()
}
