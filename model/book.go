package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// Book represents a published book
type Book struct {
	ID        string `gorm:"primary_key"`
	Title     string `gorm:"not null;unique"`
	NumPages  int32
	PubYear   int32
	AuthorID  string
	CreatedAt time.Time
}

func (b *Book) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", xid.New().String())
	return nil
}
