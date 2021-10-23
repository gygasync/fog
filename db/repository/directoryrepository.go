package repository

import (
	"database/sql"
	"fog/common"
	"fog/db/models"
)

type DirectoryRepository interface {
	Add(models.Directory) error
	Get(id string) (models.Directory, error)
	Delete(id string) error
	List(limit, offset uint) ([]models.Directory, error)
}

type Directories struct {
	db     *sql.DB
	logger *common.StdLogger
}

func NewDirectorySet(db *sql.DB, logger *common.StdLogger) *Directories {
	return &Directories{db: db, logger: logger}
}

func (dirs *Directories) Add(directory models.Directory) error {
	query := "INSERT INTO Directory (Id, Path) VALUES (?, ?)"
	_, err := dirs.db.Exec(query, directory.Id, directory.Path, directory.Dateadded, directory.Lastchecked)

	return err
}

func (dirs *Directories) Get(id string) (models.Directory, error) {
	var dir models.Directory
	query := "SELECT * FROM Directory WHERE Id = ?"
	row := dirs.db.QueryRow(query, id)
	err := row.Scan(&dir.Id, &dir.Path, &dir.Dateadded, &dir.Lastchecked)

	return dir, err
}

func (dirs *Directories) List(limit, offset uint) ([]models.Directory, error) {
	var result []models.Directory
	query := "SELECT * FROM Directory ORDER BY Id LIMIT ? OFFSET ?"
	rows, err := dirs.db.Query(query, limit, offset)
	if err != nil {
		dirs.logger.Error("Could not query the db ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dir models.Directory
		if err := rows.Scan(&dir.Id, &dir.Path, &dir.Dateadded, &dir.Lastchecked); err != nil {
			dirs.logger.Error("Could not bind remote data ", err)
			return nil, err
		}
		result = append(result, dir)
	}

	return result, nil
}

func (dirs *Directories) Delete(id string) error {
	query := "DELETE FROM Directory WHERE Id = '?'"
	rows, err := dirs.db.Query(query, id)
	if err != nil {
		dirs.logger.Error("Could not perform delete in Directory ", err)
		return err
	}

	defer rows.Close()

	return nil
}
