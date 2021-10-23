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

	conn := db.NewDbConn(props["sqliteDbLocation"], logger)
	if conn == nil {
		logger.Fatal("Could not establish connection to database, Exiting...")
		return
	}

	conn.Up()

	logger.Info("Exiting Application...")
}
