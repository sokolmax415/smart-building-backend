package postgres

import (
	"auth-service/internal/entity"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func newDBConnection(path string) (*sql.DB, error) {
	log.Printf("Openning postgres databse using connection string: %s", path)
	db, err := sql.Open("postgres", path)

	if err != nil {
		log.Printf("Error in newDBconnection: Failed to open DB: %v", err)
		return nil, entity.ErrOpenDb
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error in newDBconnection: Failed to ping DB: %v", err)
		return nil, entity.ErrConnectDb
	}

	log.Printf("Successfully openned and pinged to %s", path)
	return db, nil
}
