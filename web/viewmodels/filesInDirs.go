package viewmodels

import "fog/db/genericmodels"

type FilesInDirs struct {
	ParentDirectoryId string
	Dirs              []*genericmodels.Directory
	Files             []*genericmodels.File
}
