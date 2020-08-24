package main

import "pooling-mysql/database"

func main() {
	if err := database.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
