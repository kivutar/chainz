package model

// Book represents a published book
type Book struct {
	ID        string
	Title     string
	NumPages  int32  `db:"num_pages"`
	PubYear   int32  `db:"pub_year"`
	AuthorID  string `db:"author_id"`
	CreatedAt string `db:"created_at"`
}
