package repository

import (
	"fog/common"
	"fog/db"
	"fog/db/models"
)

type FileRepository interface {
	Add(models.File) error
	Get(id string) (models.File, error)
	Delete(id string) error
	List(limit, offset uint) ([]models.File, error)
}

type Files struct {
	db     db.DbConfig
	logger common.Logger
}

func NewFileSet(db db.DbConfig, logger common.Logger) *Files {
	return &Files{db: db, logger: logger}
}

func (dirs *Files) Add(file models.File) error {
	query := "INSERT INTO File (Id, Path, ParentDirectory) VALUES (?, ?, ?)"
	_, err := dirs.db.GetDB().Exec(query, file.Id, file.Path, file.ParentDirectory)

	return err
}

func (files *Files) Get(id string) (models.File, error) {
	var file models.File
	query := "SELECT * FROM File WHERE Id = ?"
	row := files.db.GetDB().QueryRow(query, id)
	err := row.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked)

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
		if err := rows.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked); err != nil {
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
