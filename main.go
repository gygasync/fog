package main

import (
	"fog/common"
	"fog/db"
)

func main() {
	logger := common.NewStdFileLogger()
	logger.Info("Starting Application")

	props, err := common.LoadProperties("dev")

	if err != nil {
		logger.Fatal("Could not load props.")
		return
	}

	conn, err := db.Open(props["sqliteDbLocation"], logger)
	if err != nil {
		logger.Fatal("Could not establish connection to database, Exiting...")
		return
	}

	db.Up(conn, logger)

	logger.Info("Exiting Application...")
}
