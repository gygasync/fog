package genericmodels

import "database/sql"

type IModel interface {
	ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (interface{}, sql.Result, error)
	QueryRow(query string, value interface{}, row *sql.Row) error
	ToString() string
}
