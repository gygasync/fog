package services

import (
	"database/sql"
	"fmt"
	"fog/common"
	"fog/db/genericmodels"
	"fog/db/repository"
	"fog/tasks"

	"github.com/google/uuid"
)

type IMetadataService interface {
	Add(meta *genericmodels.Metadata) (*genericmodels.Metadata, error)
}

type MetadataService struct {
	logger              common.Logger
	repository          repository.IRepository
	metadataTypeService IMetadataTypeService
	orchestartor        tasks.IOrchestrator
	metaName            string
	metaTypeId          int64
}

func NewMetadataService(
	logger common.Logger,
	repository repository.IRepository,
	metadataTypeService IMetadataTypeService,
	metaName string,
	orchestrator tasks.IOrchestrator) *MetadataService {
	metaType, err := metadataTypeService.FindOne("Name", metaName)

	var metaTypeId int64

	if err == nil {
		logger.Warnf("metaType %s already exists", metaType.Name)
		metaTypeId = metaType.Id
	} else {
		if err != sql.ErrNoRows {
			panic("Metadata Type error")
		}
		newM, err := metadataTypeService.Add(&genericmodels.MetadataType{Name: metaName})
		if err != nil {
			panic("Metadata Type insertion error")
		}
		metaTypeId = newM.Id
	}

	return &MetadataService{
		logger:              logger,
		repository:          repository,
		metadataTypeService: metadataTypeService,
		metaTypeId:          metaTypeId,
		metaName:            metaName,
		orchestartor:        orchestrator,
	}
}

func (s *MetadataService) Add(meta *genericmodels.Metadata) (*genericmodels.Metadata, error) {
	meta.Id = fmt.Sprintf("0x%x", [16]byte(uuid.New()))
	meta.MetaType = s.metaTypeId
	res, err := s.repository.Add(meta)
	if err != nil {
		return nil, err
	}

	return res.(*genericmodels.Metadata), nil
}
