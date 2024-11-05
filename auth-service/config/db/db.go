package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func InitDb() {
	dsn := os.Getenv("dsn")
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect db ")
	}
	err = DB.Ping()
	if err != nil {
		panic("Failed to ping the database")
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL
		);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Failed to create users table: %v", err))
	}
}
