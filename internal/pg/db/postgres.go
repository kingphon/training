package dbpg

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	db  *sql.DB
	err error
)

// ConnectDB connects to the PostgreSQL database.
func ConnectDB() {
	// Replace "your_database_url" with your actual PostgreSQL database URL.
	db, err = sql.Open("pgx", "postgres://admin:123456@localhost:5432/training?sslmode=disable")

	log.Println("Connected to the database")
}

// GetDB ...
func GetDB() *sql.DB {
	return db
}
