package work

import (
	"encoding/json"
	"fog/common"
	"fog/db/genericmodels"
	"fog/services"
)

type directoryWorker struct {
	logger  common.Logger
	service services.IDirectoryService
}

type directoryWorkDefinition struct {
	DirPath string
}

func NewDirectoryWorker(logger common.Logger, service services.IDirectoryService) *directoryWorker {
	return &directoryWorker{
		logger:  logger,
		service: service,
	}
}

func (d *directoryWorker) GetWorkType() string {
	return "directory"
}

func (d *directoryWorker) Work(work workDefinition) *responseDefinition {
	var def directoryWorkDefinition
	err := json.Unmarshal([]byte(work.Payload), &def)
	if err != nil {
		d.logger.Errorf("DirectoryWorker error, payload: %s, error:%s", work.Payload, err)
		return NewResponseDefinition(work, "error", GenerateErrorPayload(err))
	}

	err = d.service.Add(&genericmodels.Directory{Path: def.DirPath})
	if err != nil {
		d.logger.Errorf("DirectoryWorker error, payload: %s, error:%s", work.Payload, err)
		return NewResponseDefinition(work, "error", GenerateErrorPayload(err))
	}

	return NewResponseDefinition(work, work.WorkType, work.Payload)
}
