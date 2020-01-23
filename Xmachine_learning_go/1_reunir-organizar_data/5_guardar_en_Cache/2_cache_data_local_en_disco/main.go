package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {

	// Abrir archivo de datos embedded.db en directorio actual.
	// Se creará si no existe.
	db, err := bolt.Open("dataembebida.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear un "depósito" en el archivo boltdb para los datos.
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Bucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Colocar las claves y los valores del mapa en el archivo BoltDB.
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Bucket"))
		err := b.Put([]byte("mikey"), []byte("mivalorZero"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	// Salida de claves y valores en archivo BoltDB incorporado a la salida estándar.
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Bucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
