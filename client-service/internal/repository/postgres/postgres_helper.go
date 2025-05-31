package postgres

import (
	"client-service/internal/entity"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func newDBconnection(path string) (*sql.DB, error) {
	log.Printf("Openning postgres DB using connection string: %s", path)
	db, err := sql.Open("postgres", path)

	if err != nil {
		log.Printf("ERROR IN openning DB: %v", err)
		return nil, entity.ErrOpenDb
	}

	if err := db.Ping(); err != nil {
		log.Printf("ERROR in pinging DB: %v", err)
	}

	log.Printf("Successfully openned and pinged to %s", path)
	return db, nil
}
