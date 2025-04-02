package postgresql

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgresql() *sqlx.DB {
	dsn := "user=admin password=password dbname=bank_db sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	return db
}

func IsSQLReallyError(err error) bool {
	return err != nil && !errors.Is(err, sql.ErrNoRows)
}
