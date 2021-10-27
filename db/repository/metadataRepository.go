package repository

import (
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/models"

	"github.com/google/uuid"
)

type IMetadataRepository interface {
	Add(models.Metadata) (*models.Metadata, error)
	Get(id string) (*models.Metadata, error)
	FindOne(column string, value interface{}) (*models.Metadata, error)
	FindMany(column string, value interface{}) ([]models.Metadata, error)
}

type MetadataRepository struct {
	logger common.Logger
	db     db.DbConfig

	tableName string
}

func NewMetadataRepository(logger common.Logger, db db.DbConfig) *MetadataRepository {
	return &MetadataRepository{logger: logger, db: db, tableName: "Metadata"}
}

func (r *MetadataRepository) FindOne(column string, value interface{}) (*models.Metadata, error) {
	var meta models.Metadata
	query := GenerateFindOneQuery(r.tableName, column)
	row := r.db.GetDB().QueryRow(query, value)
	err := row.Scan(&meta.Id, &meta.MetaType, &meta.Reference, &meta.Value)

	if err != nil {
		return nil, err
	}

	return &meta, nil
}

func (r *MetadataRepository) Get(id string) (*models.Metadata, error) {
	return r.FindOne("Id", id)
}

func (r *MetadataRepository) Add(meta *models.Metadata) (*models.Metadata, error) {
	meta.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
	query := "INSERT INTO Metadata (Id, MetaType, Reference, Value) VALUES (?,?,?,?)"
	_, err := r.db.GetDB().Exec(query, meta.Id, meta.MetaType, meta.Reference, meta.Value)

	if err != nil {
		return nil, err
	}

	newMeta, err := r.Get(meta.Id)

	if err != nil {
		return nil, err
	}

	return newMeta, nil
}

func (r *MetadataRepository) FindMany(column string, value interface{}) ([]models.Metadata, error) {
	var result []models.Metadata
	query := GenerateFindManyQuery(r.tableName, column)
	rows, err := r.db.GetDB().Query(query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var meta models.Metadata
		if err := rows.Scan(&meta.Id, &meta.MetaType, &meta.Reference, &meta.Value); err != nil {
			r.logger.Error("could not bind remote data ", err)
			return nil, err
		}
		result = append(result, meta)
	}

	return result, nil
}
