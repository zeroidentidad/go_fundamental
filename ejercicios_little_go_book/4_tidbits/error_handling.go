package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type error interface {
	Error() string
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}

	input()
}

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}

	return nil
}

func input() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input!")
	}
}
