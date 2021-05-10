package app

import (
	"backend/logs"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbClient() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Fatal(err.Error())
	}

	// pool settings:
	pool, err := client.DB()
	if err != nil {
		logs.Error(err.Error())
	}
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	return client
}
