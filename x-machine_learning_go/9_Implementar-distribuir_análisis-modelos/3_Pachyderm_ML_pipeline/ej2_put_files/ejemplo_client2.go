package main

import (
	"log"
	"os"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Conéctarse a Pachyderm en localhost. Por defecto,
	// Pachyderm estará expuesto en el puerto 30650.
	c, err := client.NewFromAddress("192.168.99.100:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Iniciar commit en repositorio de datos de "attributes" en la rama "master".
	commit, err := c.StartCommit("attributes", "master")
	if err != nil {
		log.Fatal(err)
	}

	// Abrir uno de los archivos JSON de atributos.
	f, err := os.Open("1.json")
	if err != nil {
		log.Fatal(err)
	}

	// Colocar un archivo que contenga los atributos en el repositorio de datos.
	if _, err := c.PutFile("attributes", commit.ID, "1.json", f); err != nil {
		log.Fatal(err)
	}

	// Finalizar commit.
	if err := c.FinishCommit("attributes", commit.ID); err != nil {
		log.Fatal(err)
	}

	// Iniciar commit en repositorio de datos de "training" en la rama "master".
	commit, err = c.StartCommit("training", "master")
	if err != nil {
		log.Fatal(err)
	}

	// Abrir conjunto de datos de entrenamiento.
	f, err = os.Open("diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Colocar archivo que contenga dataset de entrenamiento en el repositorio de datos.
	if _, err := c.PutFile("training", commit.ID, "diabetes.csv", f); err != nil {
		log.Fatal(err)
	}

	// Finalizar commit.
	if err := c.FinishCommit("training", commit.ID); err != nil {
		log.Fatal(err)
	}

	// revisar en terminal: pachctl list repo
}
