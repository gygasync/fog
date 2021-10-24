package services

import (
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
)

type IDirectoryService interface {
	List(limit, offset uint) []models.Directory
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
