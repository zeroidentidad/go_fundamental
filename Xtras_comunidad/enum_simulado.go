package main

import (
	"fmt"
)

type UserType int

const (
	Troll  UserType = 1
	Goblin UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {

	user0 := User{Username: "Nani", Type: Troll}
	user1 := User{Username: "Mazaka", Type: Goblin}

	if user0.Type == Troll {
		fmt.Println("El usuario", user0.Username, "es un Troll")
	}

	fmt.Println(user1)
}
