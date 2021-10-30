package routines

import (
	"encoding/json"
	"fog/common"
	"fog/db/genericmodels"
	"fog/services"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

type ExifScheduler struct {
	logger      common.Logger
	metaService services.IMetadataService
	fileService services.FileService
}

func NewExifScheduler(logger common.Logger, metaService services.IMetadataService, fileService services.FileService) *ExifScheduler {
	exif.RegisterParsers(mknote.All...)
	return &ExifScheduler{logger: logger, metaService: metaService, fileService: fileService}
}

func (e *ExifScheduler) Schedule(workItems []string) IWorker {
	ex := NewExifWorker(workItems, e.metaService, &e.fileService)
	return ex
}

func NewExifWorker(workQueue []string, metaService services.IMetadataService, fileService services.IFileService) *ExifWorker {
	return &ExifWorker{workQueue: workQueue, metaService: metaService, fileService: fileService}
}

type ExifWorker struct {
	metaService services.IMetadataService
	fileService services.IFileService
	workQueue   []string
}

func (e *ExifWorker) Work() {
	for _, s := range e.workQueue {
		file, err := e.fileService.Find(s)
		if err != nil {
			continue
		}

		f, err := os.Open(file.Path)
		if err != nil {
			continue
		}

		exifData, err := exif.Decode(f)
		if err != nil {
			continue
		}

		exifJson, err := json.Marshal(exifData)
		if err != nil {
			continue
		}

		e.metaService.Add(&genericmodels.Metadata{Reference: s, Value: string(exifJson)})
	}
}
