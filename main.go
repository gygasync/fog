package main

import (
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/genericmodels"
	"fog/db/repository"
	"fog/services"
	"fog/tasks"
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
	tplEngine.RegisterTemplate("./web/static/templates/navigation.template.html", "navigation")

	router := web.NewTplRouter(logger, tplEngine)

	indexRoute := routes.NewIndexRoute(logger, tplEngine)

	orchestartor := tasks.NewOrchestrator("amqp://guest:guest@localhost:5672/", logger)

	// fileRepository := repository.NewFileRepository(logger, conn)
	fileRepository := repository.NewRepository(logger, conn.GetDB(), &genericmodels.File{})
	fileService := services.NewFileService(logger, fileRepository, orchestartor)

	// directoryRepository := repository.NewDirectorySet(conn, logger)
	directoryRepository := repository.NewRepository(logger, conn.GetDB(), &genericmodels.Directory{})
	directoryService := services.NewDirectoryService(logger, directoryRepository, fileService, orchestartor)

	// tagRepository := repository.NewTagRepository(logger, conn)
	// tagService := services.NewTagService(logger, tagRepository)

	logf := func(b []byte) { logger.Infof("%s", b) }

	exifWorkers := tasks.NewWorkerGroup("exif", "", orchestartor.GetConnection(), logger)
	orchestartor.RegisterGroup(exifWorkers)
	worker := tasks.NewWorker(orchestartor.GetConnection(), "exif", logf, logger)
	go worker.Start()

	orchestratorRoute := routes.NewOrchestratorRoute(logger, tplEngine, orchestartor)

	dirRoute := routes.NewDirRoute(logger, tplEngine, directoryService)
	dirFiles := routes.NewFilesRoute(logger, tplEngine, fileService, directoryService, exifWorkers)
	// tagRoute := routes.NewTagRoute(logger, tplEngine, tagService)

	router.RegisterRoute("/", web.GET, indexRoute)
	router.RegisterRoute("/dir", web.GET, dirRoute)
	router.RegisterRoute("/dir", web.POST, dirRoute)
	router.RegisterRoute("/files/:id", web.GET, dirFiles)
	router.RegisterRoute("/files", web.POST, dirFiles)
	router.RegisterRoute("/orchestrator", web.GET, orchestratorRoute)
	router.RegisterRoute("/orchestrator", web.POST, orchestratorRoute)
	// router.RegisterRoute("/tags", web.GET, tagRoute)
	// router.RegisterRoute("/tags", web.POST, tagRoute)

	genericFileRepo := repository.NewRepository(logger, conn.GetDB(), &genericmodels.File{})
	lis, _ := genericFileRepo.FindMany("ParentDirectory", "0xe2a69ec6e12b488eae3ea2cf1084965e")

	for _, item := range lis {
		fmt.Printf("%+v\n", item)
	}

	// dir, _ := genericDirRepo.FindOne("Id", "0x4b859d08ddb0442da48c30c038f20df3")
	// logger.Info(repository.GetModelFields(&genericmodels.Directory{}))

	// metadataRepo := repository.NewRepository(logger, conn.GetDB(), &genericmodels.Metadata{})
	// metadataTypeRepo := repository.NewRepository(logger, conn.GetDB(), &genericmodels.MetadataType{})

	// metadataTypeService := services.NewMetadataTypeService(logger, metadataTypeRepo)
	// metadataService := services.NewMetadataService(logger, metadataRepo, metadataTypeService, "EXIF")

	// exifSch := routines.NewExifScheduler(logger, metadataService, *fileService)
	// worker := exifSch.Schedule([]string{"0xd10af1c9bde24195825134e3949288bf"})

	logger.Fatal(http.ListenAndServe(":8080", router.Router()))

	logger.Info("Exiting Application...")
}
