package main

import (
	"fog/common"
	"fog/db"
	"fog/learning"
)

func main() {
	logger := learning.Getlogger(false)
	defer logger.Sync()

	learning.Log("Application started")
	learning.Log("Connecting to db")

	props, err := common.LoadProperties("dev")

	if err != nil {
		learning.Log("Could not load props")
		return
	}

	db.Open(props["sqliteDbLocation"])
	db.Up(props["sqliteDbLocation"])

	learning.Log("Exiting from application")
}
