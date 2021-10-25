package viewmodels

import "fog/db/models"

type FilesInDirs struct {
	Dirs  []models.Directory
	Files []models.File
}
