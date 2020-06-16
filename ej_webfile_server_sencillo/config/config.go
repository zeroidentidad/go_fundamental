package config

import (
	"encoding/json"
	"fmt"
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

var confObj = GetConfiguration()

func PortHttp(decorator bool) string {
	if decorator {
		fmt.Println("------------------------------------------")
		fmt.Println("Valor config.json puerto: " + confObj.HttpPort)
		fmt.Println("------------------------------------------")
	}
	return fmt.Sprintf("%s", confObj.HttpPort)
}

func DirShared() string {
	return fmt.Sprintf("%s", confObj.DirShared)
}
