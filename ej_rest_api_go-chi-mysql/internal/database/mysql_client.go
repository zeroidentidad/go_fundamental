package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(datasource string) *MySqlClient {
	db, err := sql.Open("mysql", datasource)

	if err != nil {
		fmt.Errorf("Could not create tenant db: %s", err.Error())
		panic(err)
	}

	//err = db.Ping()

	return &MySqlClient{db}
}

func (c *MySqlClient) ViewStats() sql.DBStats {
	return c.Stats()
}
