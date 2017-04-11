package location

// Country Model
type Country struct {
	ID           int    `db:"Id,omitempty"`
	Code         string `db:"Code,omitempty"`
	Name         string `db:"Name,omitempty"`
	CommonFields `db:",inline"`
}
