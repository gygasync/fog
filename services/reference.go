package services

import (
	"fog/common"
	"fog/db/models"
	"fog/db/repository"
)

type IReferenceService interface {
	TagFile(file *models.File, tag *models.Tag) error
	TagDirectory(file *models.Directory, tag *models.Tag) error
	GetItemTags(itemId string) ([]models.Tag, error)
	GetFileTags(file *models.File) ([]models.Tag, error)
	GetDirectoryTags(dir *models.Directory) ([]models.Tag, error)
}

type ReferenceService struct {
	logger     common.Logger
	repository repository.IReferenceRepository

	fileService IFileService
	tagService  ITagService
}

func NewReferenceService(
	logger common.Logger,
	repository repository.IReferenceRepository,
	fileService IFileService,
	tagService ITagService) *ReferenceService {
	return &ReferenceService{
		logger:      logger,
		repository:  repository,
		fileService: fileService,
		tagService:  tagService,
	}
}

func (s *ReferenceService) TagFile(file *models.File, tag *models.Tag) error {
	_, err := s.repository.Add(&models.Reference{Tag: tag.Id, Item: file.Id})
	return err
}

func (s *ReferenceService) TagDirectory(dir *models.Directory, tag *models.Tag) error {
	_, err := s.repository.Add(&models.Reference{Tag: tag.Id, Item: dir.Id})
	return err
}

func (s *ReferenceService) GetItemTags(itemId string) ([]models.Tag, error) {
	refs, err := s.repository.FindMany("Item", itemId)
	if err != nil {
		s.logger.Warn("error getting refs ", err)
		return nil, err
	}

	var result []models.Tag
	for _, ref := range refs {
		t, err := s.tagService.Get(ref.Tag)
		if err != nil {
			s.logger.Warnf("could not get tag for ref %d ", ref.Id)
		}
		result = append(result, *t)
	}

	return result, nil
}

func (s *ReferenceService) GetFileTags(file *models.File) ([]models.Tag, error) {
	return s.GetItemTags(file.Id)
}

func (s *ReferenceService) GetDirectoryTags(dir *models.Directory) ([]models.Tag, error) {
	return s.GetItemTags(dir.Id)
}
