package db

import (
	"database/sql"
	"fog/common"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func up(db *sql.DB, logger *common.StdLogger) {
	l := *logger
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		l.Errorf("Instance err: %v \n", err)
	}

	fileSource, err := (&file.File{}).Open("file://db/migrations")
	if err != nil {
		l.Errorf("opening file error: %v \n", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, "fog.sqlite", dbDriver)
	if err != nil {
		l.Errorf("migrate error: %v \n", err)
	}

	if err = m.Up(); err != nil {
		l.Errorf("migrate up error: %v \n", err)
	}

	l.Info("Migrate up success.")
}
