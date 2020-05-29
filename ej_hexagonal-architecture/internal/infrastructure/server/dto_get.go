package server

import "github.com/zeroidentidad/hex_arq/internal/core/domain"

type ResponsePersonGet domain.Person

func BuildResponsePersonGet(person domain.Person) ResponsePersonGet {
	return ResponsePersonGet(person)
}

type ResponseUserGet domain.User

func BuildResponseUserGet(user domain.User) ResponseUserGet {
	return ResponseUserGet(user)
}
