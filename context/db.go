package context

import (
	"log"

	"github.com/jmoiron/sqlx"
	// Side effect import of pq
	_ "github.com/lib/pq"
)

// OpenDB creates the connection to the database
func OpenDB(config *Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("postgres", config.DBURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database is connected ")
	return db, nil
}
