package main

import (
	"fmt"

	"../orm"
)

func main() {
	orm.CreateConnection()

	//user := orm.NewUser("Jesus orm", "123", "testx@mail.com")
	//user.Save()

	users := orm.GetUsers()
	fmt.Println(users)

	fmt.Println("==================")

	user := orm.GetUser(2)
	fmt.Println(user)

	fmt.Println("==================")

	user.Username = "Mod orm"
	user.Password = "Pass orm"
	user.Email = "orm@mail.com"
	user.Save()

	fmt.Println(user)

	orm.CloseConnection()
}
