package services

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type IDirectoryService interface {
	List(limit, offset uint) []models.Directory
	Add(directory models.Directory) error
}

type DirectoryService struct {
	logger     common.Logger
	repository repository.DirectoryRepository
}

func NewDirectoryService(logger common.Logger, repository repository.DirectoryRepository) *DirectoryService {
	return &DirectoryService{logger: logger, repository: repository}
}

func (s *DirectoryService) List(limit, offset uint) []models.Directory {
	result, err := s.repository.List(limit, offset)

	if err != nil {
		s.logger.Error("error when trying to execute directory List", err)
		return nil
	}

	return result
}

func (s *DirectoryService) Add(directory models.Directory) error {
	dir, err := os.Stat(directory.Path)
	if err != nil || !dir.IsDir() {
		s.logger.Warnf("directory %s is not valid", directory.Path)
		return fmt.Errorf("directory %s is not valid", directory.Path)
	}

	directory.Path, err = filepath.Abs(directory.Path)
	if err != nil {
		return err
	}

	_, err = s.repository.FindOne("Path", directory.Path)

	if err == sql.ErrNoRows {
		directory.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
		return s.repository.Add(directory)
	} else {
		if err == nil {
			s.logger.Warnf("path %s already exists", directory.Path)
			return err
		}
		s.logger.Warnf("unable to add path %s %s", directory.Path, err.Error())
		return err

	}
}
