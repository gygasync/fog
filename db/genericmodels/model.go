package genericmodels

import "database/sql"

type IModel interface {
	ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (interface{}, sql.Result, error)
	ScanRow(row *sql.Rows) error
	ToString() string
}
