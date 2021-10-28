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
	// Get(id interface{}) (genericmodels.IModel, error)
	Add(item genericmodels.IModel) (genericmodels.IModel, error)
	FindOne(column string, value interface{}) (genericmodels.IModel, error)
	FindMany(column string, value interface{}) ([]genericmodels.IModel, error)
}

type Repository struct {
	logger    common.Logger
	db        *sql.DB
	middleman genericmodels.IModel

	tableName   string
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
	}
}

func (r *Repository) FindOne(column string, value interface{}) (genericmodels.IModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT 1", r.tableName, column)
	err := r.middleman.ScanRow(r.db.QueryRow(query, value))

	if err != nil {
		return nil, err
	}

	return r.middleman, nil
}

func (r *Repository) Add(item genericmodels.IModel) error {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s %s)", r.tableName, r.fields, strings.Repeat(" ?,", r.numOfFields-1), "?")
	_, _, err := item.ExecuteQuery(query, r.db.Exec)

	if err != nil {
		return err
	}

	return nil
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
