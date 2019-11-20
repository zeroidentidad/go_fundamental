package main

import (
	"../models"
)

func main() {
	models.CreateConnection()
	models.CloseConnection()
}
