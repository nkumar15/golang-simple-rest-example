package location

import (
	"time"
)

type Country struct {
	ID        int        `db:"Id,omitempty"`
	Code      string     `db:"Code,omitempty"`
	Name      string     `db:"Name,omitempty"`
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}
