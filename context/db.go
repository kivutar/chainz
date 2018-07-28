package context

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// Side effect import of pq
	_ "github.com/lib/pq"
)

// OpenDB creates the connection to the database
func OpenDB(config *Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
	))

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
