package postgres

type ItemField struct {
	ID         int
	ItemID     int
	Item       Item
	FieldName  string
	FieldValue string
}
