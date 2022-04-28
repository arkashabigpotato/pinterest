package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx postgres database driver
	"log"
)

// PostgresConnection create connect to postgres and check it
func PostgresConnection() *sql.DB {
	dsn := fmt.Sprintf("host=localhost port=5432 dbname=postgres user=agboriskin")

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
