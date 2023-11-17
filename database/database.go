package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gs"
	password = "gs"
	dbname   = "gs"
)

var db *sql.DB

func GetDbConnection() *sql.DB {
	if db != nil {
		return db
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return db
}
