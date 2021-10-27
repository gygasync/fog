package repository

import (
	"fog/common"
	"fog/db"
	"fog/db/models"
)

type IMetadataTypeRepository interface {
	Add(models.MetadataType) (*models.MetadataType, error)
	Get(id int64) (*models.MetadataType, error)
	FindOne(column string, value interface{}) (*models.MetadataType, error)
	FindMany(column string, value interface{}) ([]models.MetadataType, error)
}

type MetadataTypeRepository struct {
	logger common.Logger
	db     db.DbConfig

	tableName string
}

func NewMetadataTypeRepository(logger common.Logger, db db.DbConfig) *MetadataTypeRepository {
	return &MetadataTypeRepository{logger: logger, db: db, tableName: "MetadataType"}
}

func (r *MetadataTypeRepository) FindOne(column string, value interface{}) (*models.MetadataType, error) {
	var meta models.MetadataType
	query := GenerateFindOneQuery(r.tableName, column)
	row := r.db.GetDB().QueryRow(query, value)
	err := row.Scan(&meta.Id, &meta.Name)

	if err != nil {
		return nil, err
	}

	return &meta, nil
}

func (r *MetadataTypeRepository) Get(id int64) (*models.MetadataType, error) {
	return r.FindOne("Id", id)
}

func (r *MetadataTypeRepository) Add(meta *models.MetadataType) (*models.MetadataType, error) {
	query := "INSERT INTO Metadata (Name) VALUES (?)"
	res, err := r.db.GetDB().Exec(query, meta.Name)

	if err != nil {
		return nil, err
	}
	newId, _ := res.LastInsertId()

	newMeta, err := r.Get(newId)

	if err != nil {
		return nil, err
	}

	return newMeta, nil
}

func (r *MetadataTypeRepository) FindMany(column string, value interface{}) ([]models.MetadataType, error) {
	var result []models.MetadataType
	query := GenerateFindManyQuery(r.tableName, column)
	rows, err := r.db.GetDB().Query(query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var meta models.MetadataType
		if err := rows.Scan(&meta.Id, &meta.Name); err != nil {
			r.logger.Error("could not bind remote data ", err)
			return nil, err
		}
		result = append(result, meta)
	}

	return result, nil
}
