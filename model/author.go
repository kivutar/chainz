package model

// Author represents a book author
type Author struct {
	ID        string
	Name      string
	CreatedAt string `db:"created_at"`
}
