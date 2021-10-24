package repository

import (
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/models"

	"github.com/google/uuid"
)

type FileRepository interface {
	Add(models.File) error
	Get(id string) (models.File, error)
	Delete(id string) error
	List(limit, offset uint) ([]models.File, error)
	FindOne(column string, value interface{}) (models.File, error)
}

type Files struct {
	logger common.Logger
	db     db.DbConfig
}

func NewFileRepository(logger common.Logger, db db.DbConfig) *Files {
	return &Files{db: db, logger: logger}
}

func (dirs *Files) Add(file models.File) error {
	file.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
	query := "INSERT INTO File (Id, Path, ParentDirectory, MimeType) VALUES (?, ?, ?, ?)"
	_, err := dirs.db.GetDB().Exec(query, file.Id, file.Path, file.ParentDirectory, file.MimeType)

	return err
}

func (files *Files) Get(id string) (models.File, error) {
	var file models.File
	query := "SELECT * FROM File WHERE Id = ?"
	row := files.db.GetDB().QueryRow(query, id)
	err := row.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked, &file.MimeType)

	return file, err
}

func (files *Files) List(limit, offset uint) ([]models.File, error) {
	var result []models.File
	query := "SELECT * FROM File ORDER BY Id LIMIT ? OFFSET ?"
	rows, err := files.db.GetDB().Query(query, limit, offset)
	if err != nil {
		files.logger.Error("Could not query the db ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var file models.File
		if err := rows.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked, &file.MimeType); err != nil {
			files.logger.Error("Could not bind remote data ", err)
			return nil, err
		}
		result = append(result, file)
	}

	return result, nil
}

func (files *Files) Delete(id string) error {
	query := "DELETE FROM File WHERE Id = ?"
	_, err := files.db.GetDB().Exec(query, id)
	if err != nil {
		files.logger.Error("Could not perform delete in Directory ", err)
		return err
	}

	return nil
}

func (files *Files) FindOne(column string, value interface{}) (models.File, error) {
	var file models.File
	query := fmt.Sprintf("SELECT * FROM File WHERE %s = ? LIMIT 1", column)
	row := files.db.GetDB().QueryRow(query, value)
	err := row.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked, &file.MimeType)

	return file, err
}
