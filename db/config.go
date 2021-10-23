package db

import (
	"database/sql"
	"fmt"
	"fog/common"

	_ "github.com/mattn/go-sqlite3"
)

func Open(connection string, logger common.Logger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connection)

	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("Error pinging DB: %v \n", err))
		return nil, err
	}

	logger.Info("Connected to SQLite db!")

	return db, nil
}
