package location

type Branch struct {
	ID   int64  `db:"Id,omitempty"`
	Code string `db:"Code"`
	Name string `db:"Name"`
	CommonFields
}
