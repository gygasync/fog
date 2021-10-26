package services

import (
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
)

type ITagService interface {
	Add(tag *models.Tag) error
	List() ([]models.Tag, error)
	Get(id string) (*models.Tag, error)
}

type TagService struct {
	logger     common.Logger
	repository repository.ITagRepository
}

func NewTagService(logger common.Logger, repository repository.ITagRepository) *TagService {
	return &TagService{logger: logger, repository: repository}
}

func (s *TagService) Add(tag *models.Tag) error {
	_, err := s.repository.Add(tag)
	return err
}

func (s *TagService) List() ([]models.Tag, error) {
	return s.repository.GetAll()
}

func (s *TagService) Get(id string) (*models.Tag, error) {
	return s.repository.Get(id)
}
