package main

import (
	"fmt"

	"pkg-errors-envolver-error/errorwrap"
)

func main() {
	errorwrap.Wrap()
	fmt.Println()
	errorwrap.Unwrap()
	fmt.Println()
	errorwrap.StackTrace()
}
