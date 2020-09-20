package db

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"webchat/utils"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", utils.Config.DBURL)
	if err != nil {
		log.Fatal(err)
	}
}

// crear un UUID aleatorio con RFC 4122
// adaptación de github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 es una variante reservada de RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Establecer los cuatro bits más significativos (bits 12 a 15) del
	// campo time_hi_and_version al número de versión de 4 bits.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])

	return uuid
}

// hash de texto plano con SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))

	return cryptext
}
