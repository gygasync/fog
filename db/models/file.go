package models

import "database/sql"

type File struct {
	Id              string
	Path            string
	ParentDirectory string
	Checksum        sql.NullString
	Lastchecked     sql.NullString
	MimeType        sql.NullString
}
