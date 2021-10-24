package services

import (
	"fmt"
	"fog/common"
	"fog/db/models"
	"fog/db/repository"

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
	directory.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
	return s.repository.Add(directory)
}
