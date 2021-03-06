package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conexion, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	//variable que ignores errores
	done := make(chan struct{})

	// funcion anonima entrada y salida paquete "io"
	go func() {
		io.Copy(os.Stdout, conexion)
		log.Println("Finalizado")
		// avisar a goroutine principal
		done <- struct{}{}
	}()

	mustCopy(conexion, os.Stdin)
	conexion.Close()
	// esperar que goroutine de conexion al servidor termine
	<-done

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
