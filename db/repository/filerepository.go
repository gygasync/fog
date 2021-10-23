package repository

import (
	"database/sql"
	"fog/common"
	"fog/db/models"
)

type FileRepository interface {
	Add(models.File) error
	Get(id string) (models.File, error)
	Delete(id string) error
	List(limit, offset uint) ([]models.File, error)
}

type Files struct {
	db     *sql.DB
	logger *common.StdLogger
}

func NewFileSet(db *sql.DB, logger *common.StdLogger) *Files {
	return &Files{db: db, logger: logger}
}

func (dirs *Files) Add(file models.File) error {
	query := "INSERT INTO File (Id, Path, ParentDirectory) VALUES (?, ?, ?)"
	_, err := dirs.db.Exec(query, file.Id, file.Path, file.ParentDirectory)

	return err
}

func (files *Files) Get(id string) (models.File, error) {
	var file models.File
	query := "SELECT * FROM File WHERE Id = ?"
	row := files.db.QueryRow(query, id)
	err := row.Scan(&file.Id, &file.Path, &file.ParentDirectory, &file.Checksum, &file.Lastchecked)

	return file, err
}

func (files *Files) List(limit, offset uint) ([]models.File, error) {
	var result []models.File
	query := "SELECT * FROM File ORDER BY Id LIMIT ? OFFSET ?"
	rows, err := files.db.Query(query, limit, offset)
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
	_, err := files.db.Exec(query, id)
	if err != nil {
		files.logger.Error("Could not perform delete in Directory ", err)
		return err
	}

	return nil
}
