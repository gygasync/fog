package services

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

type IFileService interface {
	Add(file models.File) error
	List(limit, offset uint) []models.File
	GetFilesInDir(dir *models.Directory) ([]models.File, error)
}
type FileService struct {
	logger     common.Logger
	repository repository.FileRepository
}

func NewFileService(logger common.Logger, repository repository.FileRepository) *FileService {
	return &FileService{logger: logger, repository: repository}
}

func (s *FileService) Add(file models.File) error {
	dir, err := os.Stat(file.Path)
	if err != nil || dir.IsDir() {
		s.logger.Warnf("file %s is not valid", file.Path)
		return fmt.Errorf("file %s is not valid", file.Path)
	}

	file.Path, err = filepath.Abs(file.Path)
	if err != nil {
		return err
	}

	_, err = s.repository.FindOne("Path", file.Path)

	if err != sql.ErrNoRows {
		if err == nil {
			s.logger.Warnf("path %s already exists", file.Path)
			return err
		}
		s.logger.Warnf("unable to add path %s %s", file.Path, err.Error())
		return err
	}

	// Do a mime type check
	file.MimeType = s.getMimeType(file)

	_, err = s.repository.Add(file)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileService) List(limit, offset uint) []models.File {
	result, err := s.repository.List(limit, offset)

	if err != nil {
		s.logger.Error("error when trying to execute file List", err)
		return nil
	}

	return result
}

func nullstr() sql.NullString {
	return sql.NullString{String: "", Valid: false}
}

func (s *FileService) getMimeType(file models.File) sql.NullString {
	buf, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return nullstr()
	}

	kind, err := filetype.Match(buf)
	if err != nil || kind == filetype.Unknown {
		return nullstr()
	}

	return sql.NullString{String: kind.MIME.Value, Valid: true}
}

func (s *FileService) GetFilesInDir(dir *models.Directory) ([]models.File, error) {
	return s.repository.FindMany("ParentDirectory", dir.Id)
}
