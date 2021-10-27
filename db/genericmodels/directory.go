package genericmodels

import (
	"database/sql"
	"fmt"
)

type Directory struct {
	Id              string
	Path            string
	Dateadded       string
	Lastchecked     sql.NullString
	ParentDirectory sql.NullString
}

// Can't wait for generics
func (dir *Directory) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (interface{}, sql.Result, error) {
	res, err := f(query, &dir.Id, &dir.Path, &dir.ParentDirectory, &dir.Lastchecked, &dir.ParentDirectory)
	if err != nil {
		return nil, res, err
	}
	return &dir, res, nil
}

func (dir *Directory) QueryRow(query string, value interface{}, row *sql.Row) error {
	return row.Scan(&dir.Id, &dir.Path, &dir.ParentDirectory, &dir.Lastchecked, &dir.ParentDirectory)
}

func (dir *Directory) ToString() string {
	return fmt.Sprintf("%s %s", dir.Id, dir.Path)
}
