package genericmodels

import "database/sql"

type Reference struct {
	Id   int64
	Tag  string // Reference to Tag
	Item string // Reference to either File or Directory
}

func (ref *Reference) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, ref.Id, ref.Tag, ref.Item)
	if err != nil {
		return nil, res, err
	}
	return ref, res, nil
}

func (ref *Reference) ScanRow(row *sql.Rows) (IModel, error) {
	var temp Reference
	err := row.Scan(&ref.Id, &ref.Tag, &ref.Item)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func (ref *Reference) GetId() interface{} {
	return ref.Id
}
