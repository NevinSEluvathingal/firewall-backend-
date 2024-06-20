package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not connect to database")
	}
	DB.SetMaxOpenConns(10)

	createTables(DB)
}

func createTables(db *sql.DB) {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS users (
        username TEXT NOT NULL,
        mail TEXT NOT NULL,
        password TEXT NOT NULL,
        usertype TEXT DEFAULT "user",
        ip TEXT,
        jan INTEGER DEFAULT 0,
        feb INTEGER DEFAULT 0,
        mar INTEGER DEFAULT 0,
        apr INTEGER DEFAULT 0,
        may INTEGER DEFAULT 0,
        jun INTEGER DEFAULT 0,
        jul INTEGER DEFAULT 0,
        aug INTEGER DEFAULT 0,
        sep INTEGER DEFAULT 0,
        oct INTEGER DEFAULT 0,
        nov INTEGER DEFAULT 0,
        dec INTEGER DEFAULT 0
    )`
	_, err := db.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	} else {
		fmt.Println("Users table created successfully!")
	}
}
