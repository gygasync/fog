package repository

import (
	"fog/common"
	"fog/db"
	"fog/db/models"
)

type IReferenceRepository interface {
	Add(ref *models.Reference) (*models.Reference, error)
	Get(id int64) (*models.Tag, error)
	FindOne(column string, value interface{}) (*models.Reference, error)
	FindMany(column string, value interface{}) ([]models.Reference, error)
}

type ReferenceRepository struct {
	logger common.Logger
	db     db.DbConfig

	tableName string
}

func NewReferenceRepository(logger common.Logger, db db.DbConfig) *ReferenceRepository {
	return &ReferenceRepository{logger: logger, db: db, tableName: "Reference"}
}

func (refs *ReferenceRepository) FindOne(column string, value interface{}) (*models.Reference, error) {
	var ref models.Reference
	query := GenerateFindOneQuery(refs.tableName, column)
	row := refs.db.GetDB().QueryRow(query, value)
	err := row.Scan(&ref.Id, &ref.Tag, &ref.Item)

	if err != nil {
		return nil, err
	}

	return &ref, nil
}

func (refs *ReferenceRepository) Get(id int64) (*models.Reference, error) {
	return refs.FindOne("Id", id)
}

func (refs *ReferenceRepository) Add(ref *models.Reference) (*models.Reference, error) {
	query := "INSERT INTO Reference (Tag, Item) VALUES (?, ?)"
	res, err := refs.db.GetDB().Exec(query, ref.Tag, ref.Item)

	if err != nil {
		return nil, err
	}

	newId, _ := res.LastInsertId()

	newRef, err := refs.Get(newId)

	if err != nil {
		return nil, err
	}

	return newRef, nil
}

func (refs *ReferenceRepository) FindMany(column string, value interface{}) ([]models.Reference, error) {
	var result []models.Reference
	query := GenerateFindManyQuery(refs.tableName, column)
	rows, err := refs.db.GetDB().Query(query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ref models.Reference
		if err := rows.Scan(&ref.Id, &ref.Tag, &ref.Item); err != nil {
			refs.logger.Error("could not bind remote data ", err)
			return nil, err
		}
		result = append(result, ref)
	}

	return result, nil
}
