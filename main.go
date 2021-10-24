package main

import (
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/models"
	"fog/db/repository"
	"fog/web"
	"net/http"

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

	directories := repository.NewDirectorySet(conn, logger)
	files := repository.NewFileSet(conn, logger)

	dirList, err := directories.List(100, 0)

	if err != nil {
		logger.Error(err)
	}
	someDir := dirList[0]

	someFile := models.File{Id: fmt.Sprintf("0x%x", [16]byte(uuid.New())), Path: "a.exe", ParentDirectory: someDir.Id}
	files.Add(someFile)

	fileList, err := files.List(100, 0)

	for _, file := range fileList {
		logger.Infof("%#v", file)
	}

	web.New()
	logger.Fatal(http.ListenAndServe(":8080", web.New()))

	logger.Info("Exiting Application...")
}
