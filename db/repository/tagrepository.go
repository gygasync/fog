package repository

import (
	"fog/common"
	"fog/db"
	"fog/db/models"
)

type ITagRepository interface {
	Add(tag *models.Tag) (*models.Tag, error)
	Get(id string) (*models.Tag, error)
	FindOne(column string, value interface{}) (*models.Tag, error)
	FindMany(column string, value interface{}) ([]models.Tag, error)
}

type TagRepository struct {
	logger common.Logger
	db     db.DbConfig

	tableName string
}

func NewTagRepository(logger common.Logger, db db.DbConfig) *TagRepository {
	return &TagRepository{logger: logger, db: db, tableName: "Tag"}
}

func (tags *TagRepository) FindOne(column string, value interface{}) (*models.Tag, error) {
	var tag models.Tag
	query := GenerateFindOneQuery(tags.tableName, column)
	row := tags.db.GetDB().QueryRow(query, value)
	err := row.Scan(&tag.Id, &tag.Name)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (tags *TagRepository) Get(id string) (*models.Tag, error) {
	return tags.FindOne("Id", id)
}

func (tags *TagRepository) Add(tag *models.Tag) (*models.Tag, error) {
	tag.Id = GenerateUuid()
	query := "INSERT INTO Tag (Id, Name) VALUES (?, ?)"
	_, err := tags.db.GetDB().Exec(query, tag.Id, tag.Name)

	if err != nil {
		return nil, err
	}

	newTag, err := tags.Get(tag.Id)

	if err != nil {
		return nil, err
	}

	return newTag, nil
}

func (tags *TagRepository) FindMany(column string, value interface{}) ([]models.Tag, error) {
	var result []models.Tag
	query := GenerateFindManyQuery(tags.tableName, column)
	rows, err := tags.db.GetDB().Query(query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			tags.logger.Error("could not bind remote data ", err)
			return nil, err
		}
		result = append(result, tag)
	}

	return result, nil
}
