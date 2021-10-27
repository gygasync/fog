package repository

import (
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/genericmodels"
	"reflect"
)

type IRepository interface {
	// Get(id interface{}) (genericmodels.IModel, error)
	// Add(item genericmodels.IModel) (genericmodels.IModel, error)
	FindOne(column string, value interface{}) (genericmodels.IModel, error)
	// FindMany(column string, value interface{}) ([]genericmodels.IModel, error)
}

type Repository struct {
	logger    common.Logger
	db        db.DbConfig
	tableName string
	mediator  genericmodels.IModel
}

func NewRepository(logger common.Logger, db db.DbConfig, mediator genericmodels.IModel) *Repository {
	return &Repository{logger: logger, db: db, mediator: mediator, tableName: reflect.TypeOf(mediator).Elem().Name()}
}

func (r *Repository) FindOne(column string, value interface{}) (genericmodels.IModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT 1", r.tableName, column)
	err := r.mediator.QueryRow(query, value, r.db.GetDB().QueryRow(query, value))

	if err != nil {
		return nil, err
	}

	return r.mediator, nil
}
