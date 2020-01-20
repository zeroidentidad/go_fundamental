package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// citiBikeURL proporciona los estados de las estaciones de bicicletas compartidas CitiBike.
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData se utiliza para desarmar el documento JSON devuelto desde citiBikeURL
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station se utiliza para desarmar cada uno de los documentos de la estación en stationData.
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {

	// Obtener la respuesta JSON de la URL.
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Cerrando el cuerpo de respuesta en diferido.
	defer response.Body.Close()

	// Leer cuerpo de la respuesta en []byte.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Declarar una variable de tipo stationData.
	var sd stationData

	// Descomprimir los datos JSON en la variable.
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	// Print primera estación.
	fmt.Printf("%+v\n\n", sd.Data.Stations[0])
}
