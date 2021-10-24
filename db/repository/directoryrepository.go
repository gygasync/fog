package repository

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db"
	"fog/db/models"

	"github.com/google/uuid"
)

type DirectoryRepository interface {
	Add(models.Directory) error
	Get(id string) (models.Directory, error)
	Delete(id string) error
	List(limit, offset uint) ([]models.Directory, error)
	FindOne(column string, value interface{}) (models.Directory, error)
	FindMany(column string, value interface{}) ([]models.Directory, error)
}

type Directories struct {
	db     db.DbConfig
	logger common.Logger
}

func NewDirectorySet(db db.DbConfig, logger common.Logger) *Directories {
	return &Directories{db: db, logger: logger}
}

func (dirs *Directories) Add(directory models.Directory) error {
	directory.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
	query := "INSERT INTO Directory (Id, Path) VALUES (?, ?)"
	_, err := dirs.db.GetDB().Exec(query, directory.Id, directory.Path, directory.Dateadded, directory.Lastchecked)

	return err
}

func (dirs *Directories) Get(id string) (models.Directory, error) {
	var dir models.Directory
	query := "SELECT * FROM Directory WHERE Id = ?"
	row := dirs.db.GetDB().QueryRow(query, id)
	err := row.Scan(&dir.Id, &dir.Path, &dir.Dateadded, &dir.Lastchecked)

	return dir, err
}

func (dirs *Directories) List(limit, offset uint) ([]models.Directory, error) {
	var result []models.Directory
	query := "SELECT * FROM Directory ORDER BY Id LIMIT ? OFFSET ?"
	rows, err := dirs.db.GetDB().Query(query, limit, offset)
	if err != nil {
		dirs.logger.Error("could not query the db ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dir models.Directory
		if err := rows.Scan(&dir.Id, &dir.Path, &dir.Dateadded, &dir.Lastchecked); err != nil {
			dirs.logger.Error("could not bind remote data ", err)
			return nil, err
		}
		result = append(result, dir)
	}

	return result, nil
}

func (dirs *Directories) Delete(id string) error {
	query := "DELETE FROM Directory WHERE Id = ?"
	_, err := dirs.db.GetDB().Exec(query, id)
	if err != nil {
		dirs.logger.Error("could not perform delete in Directory ", err)
		return err
	}

	return nil
}

func (dirs *Directories) FindOne(column string, value interface{}) (models.Directory, error) {
	var dir models.Directory
	query := fmt.Sprintf("SELECT * FROM Directory WHERE %s = ? LIMIT 1", column)
	row := dirs.db.GetDB().QueryRow(query, value)
	err := row.Scan(&dir.Id, &dir.Path, &dir.Dateadded, &dir.Lastchecked)

	return dir, err
}

func (dirs *Directories) FindMany(column string, value interface{}) ([]models.Directory, error) {
	var result []models.Directory
	query := fmt.Sprintf("SELECT * FROM Directory WHERE %s = ?", column)
	rows, err := dirs.db.GetDB().Query(query, value)

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		dirs.logger.Error("could not query the db ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dir models.Directory
		if err := rows.Scan(&dir.Id, &dir.Path, &dir.Dateadded, &dir.Lastchecked); err != nil {
			dirs.logger.Error("could not bind remote data ", err)
			return nil, err
		}
		result = append(result, dir)
	}

	return result, nil
}
