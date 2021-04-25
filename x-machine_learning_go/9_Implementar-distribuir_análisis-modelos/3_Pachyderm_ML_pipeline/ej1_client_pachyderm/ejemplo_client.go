package main

import (
	"log"

	// go get -v -u
	"github.com/pachyderm/pachyderm/src/client"
)

// issue ref: https://github.com/kubernetes/client-go/issues/659
// need to be fix:
// k8s.io/client-go/discovery/discovery_client.go,
// k8s.io/client-go/transport/round_trippers.go,
// k8s.io/client-go/rest/request.go,

// github.com/pachyderm/pachyderm/src/client/pkg/config/config.go
// github.com/pachyderm/pachyderm/src/client/portforwarder.go -> golang.org/pkg/context/#TODO
func main() {

	// Conéctarse a Pachyderm utilizando la IP de nuestro clúster Kubernetes.
	// Aquí se usa localhost para imitar el escenario cuando se tenga k8 ejecutados
	// localmente y/o reenvíe el puerto Pachyderm a localhost.
	// Por defecto, Pachyderm estará expuesto en el puerto 30650.
	// "0.0.0.0:30650" <- context deadline exceeded
	c, err := client.NewFromAddress("192.168.99.100:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Crear un repositorio de datos llamado "training."
	if err := c.CreateRepo("training"); err != nil {
		log.Fatal(err)
	}

	// Crear un repositorio de datos llamado "attributes."
	if err := c.CreateRepo("attributes"); err != nil {
		log.Fatal(err)
	}

	// Listar todos los repositorios de datos actuales en el clúster Pachyderm
	// como verificación. Ahora debe haber dos repositorios de datos.
	repos, err := c.ListRepo() // nil
	if err != nil {
		log.Fatal(err)
	}

	// Comprobar que el número de repos es lo que esperamos.
	if len(repos) != 2 {
		log.Fatal("Unexpected number of data repositories")
	}

	// Verificar que el nombre de repositorios sea el que esperamos.
	if repos[0].Repo.Name != "attributes" || repos[1].Repo.Name != "training" {
		log.Fatal("Unexpected data repository name")
	}

	// revisar en terminal: pachctl list repo
}
