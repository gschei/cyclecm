package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var host string
var port int
var user string
var password string
var dbname string

var db *sql.DB

func GetDbConnection() *sql.DB {
	if db != nil {
		return db
	}
	readConfig()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return db
}

func readConfig() {

	host = os.Getenv("PG_HOST")
	portStr := os.Getenv("PG_PORT")
	user = os.Getenv("PG_USER")
	password = os.Getenv("PG_PASSWORD")
	dbname = os.Getenv("PG_DBNAME")

	var err error
	port, err = strconv.Atoi(portStr)
	if err != nil {
		panic("cannot convert port to int: " + err.Error())
	}

}
