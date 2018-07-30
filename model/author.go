package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// Author represents a book author
type Author struct {
	ID        string `gorm:"primary_key"`
	Name      string `gorm:"not null;unique"`
	CreatedAt time.Time
}

func (a *Author) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", xid.New().String())
	return nil
}
