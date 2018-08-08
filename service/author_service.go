package service

import (
	"github.com/jinzhu/gorm"
	"github.com/kivutar/chainz/model"
)

type AuthorService struct {
	db *gorm.DB
}

type AuthorServer interface {
	CreateAuthor(author model.Author) (model.Author, error)
	FindByID(ID string) (model.Author, error)
	FindByBookID(bookID string) (string, error)
}

func NewAuthorService(db *gorm.DB) *AuthorService {
	return &AuthorService{db: db}
}

func (s *AuthorService) CreateAuthor(author model.Author) (model.Author, error) {
	err := s.db.Create(&author).Error
	return author, err
}

func (s *AuthorService) FindByID(ID string) (model.Author, error) {
	author := model.Author{}
	err := s.db.First(&author, "id = ?", ID).Error
	return author, err
}

func (s *AuthorService) FindByBookID(bookID string) (string, error) {
	var author model.Author
	err := s.db.Raw(`SELECT a.*
	FROM authors a, books b
	WHERE b.author_id = a.id AND b.id = ?`, bookID).First(&author).Error
	return author.ID, err
}
