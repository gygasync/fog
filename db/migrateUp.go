package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func Up(dbPath string) {
	db := Open(dbPath)
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Printf("Instance err: %v \n", err)
	}

	fileSource, err := (&file.File{}).Open("file://db/migrations")
	if err != nil {
		fmt.Printf("opening file error: %v \n", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, dbPath, dbDriver)
	if err != nil {
		fmt.Printf("migrate error: %v \n", err)
	}

	if err = m.Up(); err != nil {
		fmt.Printf("migrate up error: %v \n", err)
	}

	fmt.Println("Migrate up success.")
}
