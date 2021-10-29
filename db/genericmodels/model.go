package genericmodels

import "database/sql"

type IModel interface {
	ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error)
	ScanRow(row *sql.Rows) (IModel, error)
	GetId() interface{}
}
