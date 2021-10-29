package genericmodels

import "database/sql"

type MetadataType struct {
	Id   int64
	Name string
}

func (meta *MetadataType) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, meta.Id, meta.Name)
	if err != nil {
		return nil, res, err
	}
	return meta, res, nil
}

func (meta *MetadataType) ScanRow(row *sql.Rows) error {
	return row.Scan(&meta.Id, &meta.Name)
}

func (meta *MetadataType) GetId() interface{} {
	return meta.Id
}
