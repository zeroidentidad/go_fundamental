package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	testConnection()
	insertData()

}

func testConnection() {
	// connect to mysql 127.0.0.1
	// user: researcher
	// password: passx
	// database: researcherdb
	db, err := sql.Open("mysql", string("researcher:passx@tcp(127.0.0.1:3306)/researchdb"))
	if err != nil {
		panic(err)
	}
	err = db.Ping() // test connection
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected")
	defer db.Close()

}

func insertData() {

	db, err := sql.Open("mysql", string("researcher:passx@tcp(127.0.0.1:3306)/researchdb"))
	if err != nil {
		panic(err)
	}
	err = db.Ping() // test connection
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected")

	// prepare development
	stmt, err := db.Prepare("INSERT INTO sensordevice(deviceid,temperature,humidity,light_intensity,pressure) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	fmt.Println("inserting data...")
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(2, 0.2*float64(i), 0.6*float64(i), 0.5*float64(i), 0.3*float64(i))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("done")
	defer db.Close()

}
