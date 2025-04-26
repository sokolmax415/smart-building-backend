package postgres

import (
	"auth-service/internal/entity"
	"database/sql"

	_ "github.com/lib/pq"
)

func newDBConnection(path string) (*sql.DB, error) {
	db, err := sql.Open("postgres", path)

	if err != nil {
		return nil, entity.ErrOpenDb
	}

	if err := db.Ping(); err != nil {
		return nil, entity.ErrConnectDb
	}

	return db, nil
}
