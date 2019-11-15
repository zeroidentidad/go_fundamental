package models

import (
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{ID: 1, Username: "Zero2", Password: "xd2"}

	users[user.ID] = user
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("Error: usuario no existe.")
}

func GetUsers() Users {
	list := Users{}

	for _, user := range users {
		list = append(list, user)
	}

	return list
}
