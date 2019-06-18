package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

const (
	DbUser     = "unotifica"
	DbPassword = "x"
	DbName     = "regs_notificacion"
	DbSsl      = "disable"
)

func waitNotification(l *pq.Listener) {
	for {
		select {

		case n := <-l.Notify:
			var ja bytes.Buffer
			fmt.Println("Datos recibidos del:", n.Channel)
			if err := json.Indent(&ja, []byte(n.Extra), "", "\t"); err != nil {
				fmt.Println("Error procesando json: ", err)
				return
			}
			fmt.Println(string(ja.Bytes()))
			return

		case <-time.After(60 * time.Second):
			fmt.Println("No se ha recibido informacion, revisando conexión")

			go func() {
				l.Ping()
			}()
			return
		}
	}
}

func main() {
	conexion := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s", DbName, DbUser, DbPassword, DbSsl)

	if _, err := sql.Open("postgres", conexion); err != nil {
		panic(err)
	}

	reporte := func(let pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	listener := pq.NewListener(conexion, 10*time.Second, time.Minute, reporte)

	if err := listener.Listen("canal"); err != nil {
		panic(err)
	}

	fmt.Println("Inicia monitoreo")
	for {
		waitNotification(listener)
	}
}
