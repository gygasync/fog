package models

import "database/sql"

type Directory struct {
	Id              string
	Path            string
	Dateadded       string
	Lastchecked     sql.NullString
	ParentDirectory sql.NullString
}

// type DirectoryDump struct {
// 	Id          string
// 	Path        string
// 	Dateadded   string
// 	Lastchecked string
// }

// func (dir *DirectoryDump) Import() Directory {
// 	id, _ := hex.DecodeString(dir.Id)
// 	return Directory{Id: id, Path: dir.Path, Dateadded: dir.Dateadded, Lastchecked: dir.Lastchecked}
// }

// func (dir *Directory) Export() DirectoryDump {
// 	return DirectoryDump{Id: fmt.Sprintf("0x%x", dir.Id), Path: dir.Path, Dateadded: dir.Dateadded, Lastchecked: dir.Lastchecked}
// }
