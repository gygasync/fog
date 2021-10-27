package main

import (
	"fog/common"
	"fog/db"
	"fog/db/genericmodels"
	"fog/db/repository"
	"fog/services"
	"fog/web"
	"fog/web/routes"
	"net/http"
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

	// directories := repository.NewDirectorySet(conn, logger)
	// directories.Add(models.Directory{Id: fmt.Sprintf("0x%x", [16]byte(uuid.New())), Path: "/usr/test/"})
	// files := repository.NewFileSet(conn, logger)

	// dirList, err := directories.List(100, 0)

	// if err != nil {
	// 	logger.Error(err)
	// }
	// someDir := dirList[0]

	// someFile := models.File{Id: fmt.Sprintf("0x%x", [16]byte(uuid.New())), Path: "a.exe", ParentDirectory: someDir.Id}
	// files.Add(someFile)

	// fileList, err := files.List(100, 0)

	// for _, file := range fileList {
	// 	logger.Infof("%#v", file)
	// }

	tplEngine := routes.NewTplEngine(logger)

	tplEngine.RegisterTemplate("./web/static/templates/main.template.html", "main")
	tplEngine.RegisterTemplate("./web/static/templates/body.template.html", "body")
	tplEngine.RegisterTemplate("./web/static/templates/header.template.html", "header")

	router := web.NewTplRouter(logger, tplEngine)

	indexRoute := routes.NewIndexRoute(logger, tplEngine)

	fileRepository := repository.NewFileRepository(logger, conn)
	fileService := services.NewFileService(logger, fileRepository)

	directoryRepository := repository.NewDirectorySet(conn, logger)
	directoryService := services.NewDirectoryService(logger, directoryRepository, fileService)

	tagRepository := repository.NewTagRepository(logger, conn)
	tagService := services.NewTagService(logger, tagRepository)

	dirRoute := routes.NewDirRoute(logger, tplEngine, directoryService)
	dirFiles := routes.NewFilesRoute(logger, tplEngine, fileService, directoryService)
	tagRoute := routes.NewTagRoute(logger, tplEngine, tagService)

	router.RegisterRoute("/", web.GET, indexRoute)
	router.RegisterRoute("/dir", web.GET, dirRoute)
	router.RegisterRoute("/dir", web.POST, dirRoute)
	router.RegisterRoute("/files/:id", web.GET, dirFiles)
	router.RegisterRoute("/files", web.POST, dirFiles)
	router.RegisterRoute("/tags", web.GET, tagRoute)
	router.RegisterRoute("/tags", web.POST, tagRoute)

	genericDirRepo := repository.NewRepository(logger, conn, &genericmodels.Directory{})
	dir, _ := genericDirRepo.FindOne("Id", "0x4b859d08ddb0442da48c30c038f20df3")
	logger.Infof("%+v\n", dir)

	logger.Fatal(http.ListenAndServe(":8080", router.Router()))

	logger.Info("Exiting Application...")
}
