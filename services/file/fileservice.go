package file

import (
	"fog/common"
	"fog/db"
)

type FileServiceStruct struct {
	logger common.Logger
	db     db.DbConfig
}

type FileService interface {
	AddDirectory(dirPath string)
}

func NewFileService(logger common.Logger, db db.DbConfig) *FileServiceStruct {
	return &FileServiceStruct{logger: logger, db: db}
}
