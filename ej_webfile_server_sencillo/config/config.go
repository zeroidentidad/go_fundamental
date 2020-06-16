package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/zeroidentidad/servfiles/models"
)

func GetConfiguration() *models.Config {
	var c models.Config
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}
