package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx postgres database driver
	"log"
	"os"
)

// PostgresConnection create connect to postgres and check it
func PostgresConnection() *sql.DB {
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s port=5432 dbname=postgres user=annyka password=pass", host)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalln("cant parse postgres config", err)
	}

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(20)

	return db
}
