package repository

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/genericmodels"
	"reflect"
	"strings"
)

type IRepository interface {
	Get(id interface{}) (genericmodels.IModel, error)
	Add(item genericmodels.IModel) (genericmodels.IModel, error)
	FindOne(column string, value interface{}) (genericmodels.IModel, error)
	FindMany(column string, value interface{}) ([]genericmodels.IModel, error)
	Delete(id interface{}) error
	Update(item genericmodels.IModel) error
}

type Repository struct {
	logger    common.Logger
	db        *sql.DB
	middleman genericmodels.IModel

	tableName   string
	idColumn    string
	fields      string
	numOfFields int
}

func NewRepository(logger common.Logger, db *sql.DB, middleman genericmodels.IModel) *Repository {
	modelFields := getModelFields(middleman)
	return &Repository{
		logger:      logger,
		db:          db,
		middleman:   middleman,
		tableName:   reflect.TypeOf(middleman).Elem().Name(),
		fields:      strings.Join(modelFields, ", "),
		numOfFields: len(modelFields),
		idColumn:    modelFields[0],
	}
}

func (r *Repository) FindOne(column string, value interface{}) (genericmodels.IModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT 1", r.tableName, column)
	rows, err := r.db.Query(query, value)
	if err != nil {
		return nil, err
	}

	err = r.middleman.ScanRow(rows)
	if err != nil {
		return nil, err
	}

	return r.middleman, nil
}

func (r *Repository) Add(item genericmodels.IModel) (genericmodels.IModel, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s %s)", r.tableName, r.fields, strings.Repeat(" ?,", r.numOfFields-1), "?")
	res, _, err := item.ExecuteQuery(query, r.db.Exec)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) FindMany(column string, value interface{}) ([]genericmodels.IModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", r.tableName, column)
	rows, err := r.db.Query(query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []genericmodels.IModel

	for rows.Next() {
		if err := r.middleman.ScanRow(rows); err != nil {
			return nil, err
		}
		result = append(result, r.middleman)
	}

	return result, nil
}

func (r *Repository) Get(id interface{}) (genericmodels.IModel, error) {
	return r.FindOne(r.idColumn, id)
}

func (r *Repository) Delete(id interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", r.tableName, r.idColumn)
	_, err := r.db.Query(query, id)

	return err
}

func (r *Repository) Update(item genericmodels.IModel) error {
	err := r.Delete(item.GetId())
	if err != nil {
		return err
	}
	_, err = r.Add(item)

	return err
}

func getModelFields(model genericmodels.IModel) []string {
	vp := reflect.ValueOf(model)
	v := reflect.Indirect(vp)
	modelDef := v.Type()
	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, modelDef.Field(i).Name)
	}

	return fields
}
