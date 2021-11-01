package routines

import (
	"encoding/json"
	"fog/common"
	"fog/db/genericmodels"
	"fog/services"
	"os"
	"time"

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

func NewExifWorker(logger common.Logger, workQueue []string, metaService services.IMetadataService, fileService services.IFileService) *ExifWorker {
	return &ExifWorker{logger: logger, workQueue: workQueue, metaService: metaService, fileService: fileService}
}

type ExifWorker struct {
	logger      common.Logger
	metaService services.IMetadataService
	fileService services.IFileService
	workQueue   []string
}

func (e *ExifWorker) Do() <-chan interface{} {

	r := make(chan interface{})
	go func() {
		e.logger.Info("Starting EXIF work")
		defer close(r)
		time.Sleep(time.Second * 3)
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
			res, err := e.metaService.Add(&genericmodels.Metadata{Reference: s, Value: string(exifJson)})
			r <- res
			e.logger.Info("Ending EXIF work")
		}
	}()

	return r
}
