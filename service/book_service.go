package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kivutar/chainz/model"
	"github.com/op/go-logging"
	"github.com/rs/xid"
)

const (
	defaultListFetchSize = 10
)

type BookService struct {
	db            *sqlx.DB
	authorService *AuthorService
	log           *logging.Logger
}

func NewBookService(db *sqlx.DB, authorService *AuthorService, log *logging.Logger) *BookService {
	return &BookService{db: db, authorService: authorService, log: log}
}

func (s *BookService) FindByTitle(title string) (*model.Book, error) {
	book := &model.Book{}

	bookSQL := `SELECT * FROM books WHERE title = $1`
	udb := s.db.Unsafe()
	row := udb.QueryRowx(bookSQL, title)
	err := row.StructScan(book)
	if err == sql.ErrNoRows {
		return book, nil
	}
	if err != nil {
		s.log.Errorf("Error in retrieving book : %v", err)
		return nil, err
	}
	return book, nil
}

func (s *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	bookID := xid.New()
	book.ID = bookID.String()
	bookSQL := `INSERT INTO books (id, title, pub_year, num_pages, author_id)
			VALUES (:id, :title, :pub_year, :num_pages, :author_id)`

	_, err := s.db.NamedExec(bookSQL, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) List() ([]*model.Book, error) {
	books := make([]*model.Book, 0)

	bookSQL := `SELECT * FROM books ORDER BY created_at DESC;`
	err := s.db.Select(&books, bookSQL)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (u *BookService) Count() (int, error) {
	var count int
	bookSQL := `SELECT count(*) FROM books`
	err := u.db.Get(&count, bookSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}
