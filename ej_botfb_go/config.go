package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/gommon/log"
)

// c variable que tiene toda la información del archivo de configuración
var c Config

// Config estructura de los datos variables de la app
type Config struct {
	Port    string `json:"port"`
	CertPem string `json:"cert_pem"`
	KeyPem  string `json:"key_pem"`
	MyToken string `json:"my_token"`
	FbToken string `json:"fb_token"`
	FbURL   string `json:"fb_url"`
}

// getConfig del archivo config.json
func getConfig() {
	log.Print("Leyendo conf.json...")
	b, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		log.Fatalf("Se encontró error al leer archivo conf.json: %v", err)
	}

	err = json.Unmarshal(b, &c) //apuntando a var c Config
	if err != nil {
		log.Fatalf("Error al convertir el json a estructura: %v", err)
	}

	log.Print("Lectura de configuración finalizada...")
}
