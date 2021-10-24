package db

import (
	"fog/common"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func up(dbConfig DbConfig, logger common.Logger) {
	dbDriver, err := sqlite3.WithInstance(dbConfig.GetDB(), &sqlite3.Config{})
	if err != nil {
		logger.Errorf("Instance err: %v \n", err)
	}

	fileSource, err := (&file.File{}).Open("file://db/migrations")
	if err != nil {
		logger.Errorf("opening file error: %v \n", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, "fog.sqlite", dbDriver)
	if err != nil {
		logger.Errorf("migrate error: %v \n", err)
	}

	if err = m.Up(); err != nil {
		logger.Errorf("migrate up error: %v \n", err)
	}

	logger.Info("Migrate up success.")
}
