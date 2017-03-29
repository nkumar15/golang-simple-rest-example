package location

import (
	"database/sql"
	"log"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

type Env struct {
	DB *sql.DB
}

type CommonFields struct {
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}

var settings = sqlite.ConnectionURL{
	Database: `./db/migrations/database.sqlite`,
}

func ConnectDB() sqlbuilder.Database {
	db, err := sqlite.Open(settings)

	if err != nil {
		log.Fatal("sqlite.open: %s", err)
		panic(err)
	}

	if err1 := db.Ping(); err1 != nil {
		log.Fatal("Not able to ping database.", err1)
		panic(err)
	}

	return db
}
