package genericmodels

import (
	"database/sql"
)

type File struct {
	Id              string
	Path            string
	ParentDirectory string
	Checksum        sql.NullString
	Lastchecked     sql.NullString
	MimeType        sql.NullString
}

// Can't wait for generics
func (file *File) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, file.Id, file.Path, file.ParentDirectory, file.Checksum, file.Lastchecked, file.MimeType)
	if err != nil {
		return nil, res, err
	}
	return file, res, nil
}

func (file *File) ScanRow(row *sql.Rows) error {
	return row.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked, &file.MimeType)
}

func (file *File) GetId() interface{} {
	return file.Id
}
