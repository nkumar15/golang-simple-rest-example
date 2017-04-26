package location

import (
	"time"

	"upper.io/db.v2/lib/sqlbuilder"
)

type lcDatabase struct {
	DB sqlbuilder.Database
}

//Env ... Global environment
type Env struct {
	DBInstance string
	Database   lcDatabase
}

// CommonFields ... Used in models
type CommonFields struct {
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}
