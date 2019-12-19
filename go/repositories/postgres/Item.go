package postgres

type Item struct {
	ID          int
	Name        string
	Description string
	Picture     string
	Fields      []ItemField
}
