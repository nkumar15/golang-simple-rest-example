package location

import (
	"errors"
	"log"
	"time"

	"upper.io/db.v2/lib/sqlbuilder"
	"upper.io/db.v2/postgresql"
	"upper.io/db.v2/sqlite"
)

//Server ... Set this value to sqlite3 or pg to change underlying server
var Server string

func init() {
	Server = "sqlite3"
}

type lcDatabase struct {
	DB sqlbuilder.Database
}

//Env ... Global environment
type Env struct {
	Database lcDatabase
}

// CommonFields ... Used in models
type CommonFields struct {
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}

var sqliteSettings = sqlite.ConnectionURL{
	Database: `D:\programming\database\location\database.sqlite`,
}

var pgSettings = postgresql.ConnectionURL{
	Host:     "localhost", // PostgreSQL server IP or name.
	Database: "test",      // Database name.
	User:     "postgres",  // Optional user name.
	Password: "abc123",    // Optional user password.
}

// ConnectDB ... Connects to database and returns the database object if successful. Otherwise err is also returned
func ConnectDB() (sqlbuilder.Database, error) {
	var db sqlbuilder.Database
	var err error

	if Server == "sqlite3" {
		db, err = sqlite.Open(sqliteSettings)
	} else if Server == "pg" {
		db, err = postgresql.Open(pgSettings)
	} else {
		log.Fatal("Invalid database")
		return nil, errors.New("Invalid database name")
	}

	if err != nil {
		log.Fatal(Server, ".open: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Not able to ping database.", err)
		return nil, err
	}
	return db, nil
}
