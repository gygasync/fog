package services

import (
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
)

type IFileService interface {
	Add(dirPath string) error
	List(limit, offset uint) []models.File
}
type FileService struct {
	logger     common.Logger
	repository repository.FileRepository
}

func NewFileService(logger common.Logger, repository repository.FileRepository) *FileService {
	return &FileService{logger: logger, repository: repository}
}

func (s *FileService) Add(dirPath string) error {
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
