package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE "

var refStringSlice = []string{
	" NOMBRE = 'Jesus' ",
	" NO_SEGURO = 123456789 ",
	" ACTIVO_DESDE = SYSDATE "}

type JoinFunc func(pieza string) string

func main() {

	jF := func(p string) string {
		if strings.Contains(p, "SEGURO") {
			return "OR"
		}

		return "AND"
	}
	result := JoinWithFunc(refStringSlice, jF)
	fmt.Println(selectBase + result)
}

func JoinWithFunc(refStringSlice []string, joinFunc JoinFunc) string {
	concatenar := refStringSlice[0]
	for _, val := range refStringSlice[1:] {
		concatenar = concatenar + joinFunc(val) + val
	}
	return concatenar
}
