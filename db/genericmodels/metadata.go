package genericmodels

import "database/sql"

type Metadata struct {
	Id        string
	MetaType  int64
	Reference string
	Value     string
}

func (meta *Metadata) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, meta.Id, meta.MetaType, meta.Reference, meta.Value)
	if err != nil {
		return nil, res, err
	}
	return meta, res, nil
}

func (meta *Metadata) ScanRow(row *sql.Rows) error {
	return row.Scan(&meta.Id, &meta.MetaType, &meta.Reference, &meta.Value)
}

func (meta *Metadata) GetId() interface{} {
	return meta.Id
}
