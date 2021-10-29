package genericmodels

import (
	"database/sql"
)

type Directory struct {
	Id              string
	Path            string
	Dateadded       string
	Lastchecked     sql.NullString
	ParentDirectory sql.NullString
}

func (dir *Directory) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, dir.Id, dir.Path, dir.Dateadded, dir.Lastchecked, dir.ParentDirectory)
	if err != nil {
		return nil, res, err
	}
	return dir, res, nil
}

func (dir *Directory) ScanRow(row *sql.Rows) (IModel, error) {
	var temp Directory
	err := row.Scan(&temp.Id, &temp.Path, &temp.Dateadded, &temp.Lastchecked, &temp.ParentDirectory)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func (dir *Directory) GetId() interface{} {
	return dir.Id
}
