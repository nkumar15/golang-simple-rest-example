package location

import (
	"errors"
	"log"
	"time"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/sqlite"
)

var Server string

type Env struct {
	db *sqlbuilder.Database
}

type CommonFields struct {
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}

var sqliteSettings = sqlite.ConnectionURL{
	Database: `./db/migrations/database.sqlite`,
}

// ConnectionURL implements a PostgreSQL connection struct.
var pgSettings = postgresql.ConnectionURL{
	Host:     "localhost", // PostgreSQL server IP or name.
	Database: "test",      // Database name.
	User:     "postgres",  // Optional user name.
	Password: "abc123",    // Optional user password.
}

func init() {
	Server = "sqlite3"
}

// Connects to database and returns the database object if successful
// Otherwise err is also returned
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
