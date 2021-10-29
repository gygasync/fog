package genericmodels

import "database/sql"

type Tag struct {
	Id   string
	Name string
}

func (tag *Tag) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, tag.Id, tag.Name)
	if err != nil {
		return nil, res, err
	}
	return tag, res, nil
}

func (tag *Tag) ScanRow(row *sql.Rows) error {
	return row.Scan(&tag.Id, &tag.Name)
}

func (tag *Tag) GetId() interface{} {
	return tag.Id
}
