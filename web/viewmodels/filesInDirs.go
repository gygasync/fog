package viewmodels

import "fog/db/models"

type FilesInDirs struct {
	ParentDirectoryId string
	Dirs              []models.Directory
	Files             []models.File
}
