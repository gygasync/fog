package services

import (
	"fog/common"
	"fog/db/genericmodels"
	"fog/db/repository"
)

type IMetadataTypeService interface {
	Add(meta *genericmodels.MetadataType) (*genericmodels.MetadataType, error)
	FindOne(column string, value interface{}) (*genericmodels.MetadataType, error)
}

type MetadataTypeService struct {
	logger     common.Logger
	repository repository.IRepository
}

func NewMetadataTypeService(logger common.Logger, repository repository.IRepository) *MetadataTypeService {
	return &MetadataTypeService{logger: logger, repository: repository}
}

func (s *MetadataTypeService) Add(meta *genericmodels.MetadataType) (*genericmodels.MetadataType, error) {
	res, err := s.repository.Add(meta)
	if err != nil {
		return nil, err
	}
	return res.(*genericmodels.MetadataType), nil
}

func (s *MetadataTypeService) FindOne(column string, value interface{}) (*genericmodels.MetadataType, error) {
	res, err := s.repository.FindOne(column, value)
	if err != nil {
		return nil, err
	}
	return res.(*genericmodels.MetadataType), nil
}
