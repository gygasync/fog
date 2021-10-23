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

	db.Open(props["sqliteDbLocation"])
	db.Up(props["sqliteDbLocation"])

	logger.Info("Exiting Application...")
}
