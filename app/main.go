package main

import (
	"log"
	"net/http"

	"github.com/nkumar15/location"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/sqlite"
)

var sqliteSettings = sqlite.ConnectionURL{
	Database: `D:\programming\database\location\database.sqlite`,
}

var pgSettings = postgresql.ConnectionURL{
	Host:     "localhost", // PostgreSQL server IP or name.
	Database: "test",      // Database name.
	User:     "postgres",  // Optional user name.
	Password: "abc123",    // Optional user password.
}

func serveWeb() {
	var db sqlbuilder.Database
	var err error

	var Server = "sqlite3"
	if Server == "sqlite3" {
		db, err = sqlite.Open(sqliteSettings)
	} else if Server == "pg" {
		db, err = postgresql.Open(pgSettings)
	} else {
		log.Fatal("Invalid database")
	}

	if err != nil {
		log.Fatal(Server, ".open: %s", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Not able to ping database.", err)
		return
	}

	env := location.Env{}
	env.Database.DB = db

	router := env.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	serveWeb()
}
