package service

import (
	"github.com/jinzhu/gorm"
	"github.com/kivutar/chainz/model"
)

type BookService struct {
	db *gorm.DB
}

type BookServer interface {
	CreateBook(book model.Book) (model.Book, error)
	FindByTitle(title string) (model.Book, error)
	List() ([]model.Book, error)
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book model.Book) (model.Book, error) {
	err := s.db.Create(&book).Error
	return book, err
}

func (s *BookService) FindByTitle(title string) (model.Book, error) {
	book := model.Book{}
	err := s.db.First(&book, "title = ?", title).Error
	return book, err
}

func (s *BookService) List() ([]model.Book, error) {
	var books []model.Book
	err := s.db.Find(&books).Error
	return books, err
}

func (s *BookService) Count() (uint, error) {
	var count uint
	var books []model.Book
	err := s.db.Find(&books).Count(&count).Error
	return count, err
}
