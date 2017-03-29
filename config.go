package location

import (
	"log"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

type Env struct {
	db *sqlbuilder.Database
}

type CommonFields struct {
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}

var settings = sqlite.ConnectionURL{
	Database: `./db/migrations/database.sqlite`,
}

// Connects to database and returns the database object if successful
// Otherwise err is also returned
func ConnectDB() (sqlbuilder.Database, error) {

	db, err := sqlite.Open(settings)
	if err != nil {
		log.Fatal("sqlite.open: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Not able to ping database.", err)
		return nil, err
	}
	return db, nil
}
