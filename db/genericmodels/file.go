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

func (file *File) ScanRow(row *sql.Rows) (IModel, error) {
	var temp File
	err := row.Scan(&temp.Id, &temp.Path, &temp.ParentDirectory, &temp.Checksum, &temp.Lastchecked, &temp.MimeType)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func (file *File) GetId() interface{} {
	return file.Id
}
