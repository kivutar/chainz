package context

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kivutar/chainz/model"
	// Side effect import of postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDB creates the connection to the database
func OpenDB(config *Config) (*gorm.DB, error) {
	log.Println("Database is connecting... ")
	db, err := gorm.Open("postgres", config.DBURL)

	if err != nil {
		return nil, err
	}

	db.DropTableIfExists(&model.Author{}, &model.Book{})
	db.Debug().AutoMigrate(&model.Author{}, &model.Book{})

	log.Println("Database is connected ")
	return db, nil
}
