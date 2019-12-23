package postgres

type Item struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	Picture        string `db:"picture"`
	Description    string `db:"description"`
}
