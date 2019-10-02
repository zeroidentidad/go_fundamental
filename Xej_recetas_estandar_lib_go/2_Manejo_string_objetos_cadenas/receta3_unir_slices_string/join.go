package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE %s "

var refStringSlice = []string{
	" NOMBRE = 'Jesus' ",
	" NO_SEGURO = 123456789 ",
	" ACTIVO_DESDE = SYSDATE "}

func main() {

	sentencia := strings.Join(refStringSlice, "AND")
	fmt.Printf(selectBase+"\n", sentencia)

}
