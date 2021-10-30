package genericmodels

import "database/sql"

type TagToTag struct {
	Id     int64
	Source string // Reference to Tag
	Target string // Reference to Tag
}

func (tag *TagToTag) ExecuteQuery(query string, f func(string, ...interface{}) (sql.Result, error)) (IModel, sql.Result, error) {
	res, err := f(query, tag.Id, tag.Source, tag.Target)
	if err != nil {
		return nil, res, err
	}
	return tag, res, nil
}

func (tag *TagToTag) ScanRow(row *sql.Rows) (IModel, error) {
	var temp TagToTag
	err := row.Scan(&temp.Id, &temp.Source, &temp.Target)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func (tag *TagToTag) GetId() interface{} {
	return tag.Id
}
