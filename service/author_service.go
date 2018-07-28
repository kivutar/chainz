package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kivutar/chainz/model"
	"github.com/op/go-logging"
	"github.com/rs/xid"
)

type AuthorService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewAuthorService(db *sqlx.DB, log *logging.Logger) *AuthorService {
	return &AuthorService{db: db, log: log}
}

func (s *AuthorService) CreateAuthor(author *model.Author) (*model.Author, error) {
	authorID := xid.New()
	author.ID = authorID.String()
	authorSQL := `INSERT INTO authors (id, name)
			VALUES (:id, :name)`

	_, err := s.db.NamedExec(authorSQL, author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *AuthorService) FindByID(ID string) (*model.Author, error) {
	author := &model.Author{}

	authorSQL := `SELECT * FROM authors WHERE id = $1`
	udb := s.db.Unsafe()
	row := udb.QueryRowx(authorSQL, ID)
	err := row.StructScan(author)
	if err == sql.ErrNoRows {
		return author, nil
	}
	if err != nil {
		s.log.Errorf("Error in retrieving author : %v", err)
		return nil, err
	}
	return author, nil
}

func (r *AuthorService) FindByBookId(bookId string) (string, error) {
	var author *model.Author

	authorSQL := `SELECT a.*
	FROM authors a, books b
	WHERE b.author_id = $1 `
	err := r.db.Get(&author, authorSQL, bookId)
	if err == sql.ErrNoRows {
		return author.ID, nil
	}
	if err != nil {
		return "", err
	}
	return author.ID, nil
}
