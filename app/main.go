package main

import (
	"errors"
	"log"
	"net/http"

	logs "github.com/astaxie/beego/logs"
	"github.com/nkumar15/location"
	"github.com/urfave/negroni"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/sqlite"
)

func connectDB() (sqlbuilder.Database, error) {

	var sqliteSettings = sqlite.ConnectionURL{
		Database: `D:\codes\database\location\database.sqlite`,
	}

	var pgSettings = postgresql.ConnectionURL{
		Host:     "localhost", // PostgreSQL server IP or name.
		Database: "test",      // Database name.
		User:     "postgres",  // Optional user name.
		Password: "abc123",    // Optional user password.
	}

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
		return db, errors.New("Couldn't connect database")
	}

	err = db.Ping()
	if err != nil {
		return db, errors.New("Couldn't connect database")
	}
	return db, nil
}

func serveWeb() {
	env := location.Env{}
	db, err := connectDB()

	if err != nil {
		log.Fatal("Not able to connect database.", err)
	}
	env.Database.DB = db

	logger := logs.NewLogger()
	var config location.LogConfig
	env.SetupLogger(logger, config)

	router := env.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":5000", n))
}

func main() {
	serveWeb()
}
