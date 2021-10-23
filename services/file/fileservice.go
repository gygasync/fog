package file

import (
	"database/sql"
	"fog/common"
)

type FileServiceStruct struct {
	logger *common.StdLogger
	conn   *sql.DB
}

type FileService interface {
	AddDirectory(dirPath string)
}

func NewFileService(logger *common.StdLogger, conn *sql.DB) *FileServiceStruct {
	return &FileServiceStruct{logger: logger, conn: conn}
}
