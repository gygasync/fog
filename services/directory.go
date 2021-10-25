package services

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
	"fog/web/viewmodels"
	"io/ioutil"
	"os"
	"path/filepath"
)

type IDirectoryService interface {
	List(limit, offset uint) []models.Directory
	Add(directory models.Directory) error
	GetChildren(id string) (*viewmodels.FilesInDirs, error)
	FindChildren(directory *models.Directory) ([]models.Directory, error)
}

type DirectoryService struct {
	logger      common.Logger
	repository  repository.DirectoryRepository
	fileService IFileService
}

func NewDirectoryService(logger common.Logger, repository repository.DirectoryRepository, fileService IFileService) *DirectoryService {
	return &DirectoryService{logger: logger, repository: repository, fileService: fileService}
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

	if err != sql.ErrNoRows {
		if err == nil {
			s.logger.Warnf("path %s already exists", directory.Path)
			return err
		}
		s.logger.Warnf("unable to add path %s %s", directory.Path, err.Error())
		return err
	}
	newDir, err := s.repository.Add(directory)

	if err != nil {
		return err
	}

	err = s.workDirectory(newDir)
	if err != nil {
		s.logger.Error("error traversing directory", err)
	}

	return nil
}

func (s *DirectoryService) workDirectory(directory *models.Directory) error {
	files, err := ioutil.ReadDir(directory.Path)
	if err != nil {
		return err
	}

	for _, f := range files {
		path := filepath.Join(directory.Path, f.Name())
		if f.IsDir() {
			s.Add(models.Directory{Path: path, ParentDirectory: sql.NullString{String: directory.Id, Valid: true}})
		} else {
			s.fileService.Add(models.File{Path: path, ParentDirectory: directory.Id})
		}
	}

	return nil
}

func (s *DirectoryService) GetChildren(id string) (*viewmodels.FilesInDirs, error) {
	parent, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	dirs, err := s.FindChildren(parent)
	if err != nil {
		return nil, err
	}

	files, err := s.fileService.GetFilesInDir(parent)
	if err != nil {
		return nil, err
	}

	var result viewmodels.FilesInDirs
	result.Files = files
	result.Dirs = dirs
	result.ParentDirectoryId = parent.ParentDirectory.String

	return &result, nil
}

func (s *DirectoryService) FindChildren(directory *models.Directory) ([]models.Directory, error) {
	return s.repository.FindMany("ParentDirectory", directory.Id)
}
