package main

import (
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/models"
	"fog/db/repository"

	"github.com/google/uuid"
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

	directories := repository.NewDirectorySet(conn.GetDB(), logger)
	testDirectory := models.Directory{Id: fmt.Sprintf("0x%x", [16]byte(uuid.New())), Path: "/var/test"}

	err = directories.Add(testDirectory)

	if err == nil {
		logger.Info("Added into directory")
	} else {
		logger.Error("Could not add into directory", err)
	}

	dirList, err := directories.List(100, 0)

	if err != nil {
		logger.Error(err)
	}

	for _, dir := range dirList {
		logger.Infof("%#v", dir)
	}

	directories.Delete(testDirectory.Id)
	logger.Info("#### PERFORM DELETE ####")
	dirList, err = directories.List(100, 0)

	for _, dir := range dirList {
		logger.Infof("%#v", dir)
	}

	logger.Info("Exiting Application...")
}
