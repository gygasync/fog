package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open(location string) *sql.DB {
	db, err := sql.Open("sqlite3", location)

	if err != nil {
		fmt.Printf("Error opening DB: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging DB: %v \n", err)
	}

	fmt.Println("Connected to db")

	return db
}
