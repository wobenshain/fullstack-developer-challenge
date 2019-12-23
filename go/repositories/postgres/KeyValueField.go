package postgres

type KeyValueField struct {
	ID     int64  `db:"id"`
	ItemID int    `db:"item_id"`
	Name   string `db:"name"`
	Value  string `db:"value"`
}
