package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("dsn")
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect db")
	}

	err = DB.Ping()
	if err != nil {
		panic("Failed to ping db")
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
		    id serial PRIMARY KEY,
		    username VARCHAR(255) NOT NULL,
		    email VARCHAR(255) NOT NULL UNIQUE,
		    avatar TEXT,
		    description TEXT
		);
	`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Failed to create a users table: %v", err))
	}
}
