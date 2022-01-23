package services

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/genericmodels"
	"fog/db/repository"
	"fog/tasks"
	"fog/web/viewmodels"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type IDirectoryService interface {
	List(limit, offset uint) []*genericmodels.Directory
	Add(directory *genericmodels.Directory) error
	GetChildren(id string) (*viewmodels.FilesInDirs, error)
	FindChildren(directory *genericmodels.Directory) ([]*genericmodels.Directory, error)
}

type DirectoryService struct {
	logger       common.Logger
	repository   repository.IRepository
	fileService  IFileService
	orchestrator tasks.IOrchestrator
}

func NewDirectoryService(logger common.Logger, repository repository.IRepository, fileService IFileService, orchestrator tasks.IOrchestrator) *DirectoryService {
	return &DirectoryService{logger: logger, repository: repository, fileService: fileService, orchestrator: orchestrator}
}

func (s *DirectoryService) List(limit, offset uint) []*genericmodels.Directory {
	result, err := s.repository.List(limit, offset)

	if err != nil {
		s.logger.Error("error when trying to execute directory List", err)
		return nil
	}

	return s.derefArray(result)
}

func (s *DirectoryService) derefArray(input []genericmodels.IModel) []*genericmodels.Directory {
	r := make([]*genericmodels.Directory, len(input))
	for i, _ := range input {
		a, ok := input[i].(*genericmodels.Directory)
		if !ok {
			return nil
		}
		r[i] = a
	}
	return r
}

func (s *DirectoryService) Add(directory *genericmodels.Directory) error {
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
	directory.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
	directory.Dateadded = time.Now().Format(time.RFC3339)
	newDir, err := s.repository.Add(directory)

	if err != nil {
		return err
	}

	err = s.workDirectory(newDir.(*genericmodels.Directory))
	if err != nil {
		s.logger.Error("error traversing directory", err)
	}

	return nil
}

func (s *DirectoryService) workDirectory(directory *genericmodels.Directory) error {
	files, err := ioutil.ReadDir(directory.Path)
	if err != nil {
		return err
	}

	for _, f := range files {
		path := filepath.Join(directory.Path, f.Name())
		if f.IsDir() {
			s.Add(&genericmodels.Directory{Path: path, ParentDirectory: sql.NullString{String: directory.Id, Valid: true}})
		} else {
			s.fileService.Add(&genericmodels.File{Path: path, ParentDirectory: directory.Id})
		}
	}

	return nil
}

func (s *DirectoryService) GetChildren(id string) (*viewmodels.FilesInDirs, error) {
	parent, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	deref := parent.(*genericmodels.Directory)

	dirs, err := s.FindChildren(deref)
	if err != nil {
		return nil, err
	}

	files, err := s.fileService.GetFilesInDir(deref)
	if err != nil {
		return nil, err
	}

	var result viewmodels.FilesInDirs
	result.Files = files
	result.Dirs = dirs
	result.ParentDirectoryId = deref.ParentDirectory.String
	result.BasePath = deref.Path

	return &result, nil
}

func (s *DirectoryService) FindChildren(directory *genericmodels.Directory) ([]*genericmodels.Directory, error) {
	res, err := s.repository.FindMany("ParentDirectory", directory.Id)
	if err != nil {
		return nil, err
	}
	return s.derefArray(res), nil
}
