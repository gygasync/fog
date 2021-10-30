package services

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/genericmodels"
	"fog/db/repository"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

type IFileService interface {
	Add(file *genericmodels.File) error
	List(limit, offset uint) []*genericmodels.File
	GetFilesInDir(dir *genericmodels.Directory) ([]*genericmodels.File, error)
	Find(Id string) (*genericmodels.File, error)
}
type FileService struct {
	logger     common.Logger
	repository repository.IRepository
}

func NewFileService(logger common.Logger, repository repository.IRepository) *FileService {
	return &FileService{logger: logger, repository: repository}
}

func (s *FileService) Find(Id string) (*genericmodels.File, error) {
	res, err := s.repository.FindOne("Id", Id)
	if err != nil {
		return nil, err
	}

	return res.(*genericmodels.File), err
}

func (s *FileService) Add(file *genericmodels.File) error {
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
	file.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))

	_, err = s.repository.Add(file)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileService) derefArray(input []genericmodels.IModel) []*genericmodels.File {
	r := make([]*genericmodels.File, len(input))
	for i, _ := range input {
		a, ok := input[i].(*genericmodels.File)
		if !ok {
			return nil
		}
		r[i] = a
	}
	return r
}

func (s *FileService) List(limit, offset uint) []*genericmodels.File {
	result, err := s.repository.List(limit, offset)

	if err != nil {
		s.logger.Error("error when trying to execute file List", err)
		return nil
	}

	return s.derefArray(result)
}

func nullstr() sql.NullString {
	return sql.NullString{String: "", Valid: false}
}

func (s *FileService) getMimeType(file *genericmodels.File) sql.NullString {
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

func (s *FileService) GetFilesInDir(dir *genericmodels.Directory) ([]*genericmodels.File, error) {
	res, err := s.repository.FindMany("ParentDirectory", dir.Id)
	if err != nil {
		return nil, err
	}
	return s.derefArray(res), nil
}
