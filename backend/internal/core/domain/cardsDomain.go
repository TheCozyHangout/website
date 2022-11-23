package domain

type Card struct {
	Name        string `db:"name"`
	Identifier  string `db:"identifier"` // This is the indexed row in the database, use this for searching
	Description string `db:"description"`
	Damage      int8   `db:"damage"`
	Health      int8   `db:"health"`
	Cost        int8   `db:"cost"`
}
