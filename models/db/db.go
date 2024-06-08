package db

import (
	"database/sql"
	"log"
	"fmt"
	_"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "api.db") \
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
        jan INTEGER DEFAULT 0,
        feb INTEGER DEFAULT 0,
        march INTEGER DEFAULT 0,
        april INTEGER DEFAULT 0,
        may INTEGER DEFAULT 0,
        june INTEGER DEFAULT 0,
        july INTEGER DEFAULT 0,
        aug INTEGER DEFAULT 0,
        sept INTEGER DEFAULT 0,
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



