package work

import (
	"encoding/json"
	"fog/common"
	"fog/db/genericmodels"
	"fog/services"
	"fog/work/definition"
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

func (d *directoryWorker) Work(work definition.Work) *definition.Response {
	d.logger.Infof("Directory work id %s started", work.Id)
	var def directoryWorkDefinition
	err := json.Unmarshal([]byte(work.Payload), &def)
	if err != nil {
		d.logger.Errorf("DirectoryWorker error, payload: %s, error:%s", work.Payload, err)
		return definition.NewResponseDefinition(work, "error", definition.GenerateErrorPayload(err))
	}

	err = d.service.Add(&genericmodels.Directory{Path: def.DirPath})
	if err != nil {
		d.logger.Errorf("DirectoryWorker error, payload: %s, error:%s", work.Payload, err)
		return definition.NewResponseDefinition(work, "error", definition.GenerateErrorPayload(err))
	}

	response := definition.NewResponseDefinition(work, work.WorkType, work.Payload)
	d.logger.Infof("Directory work id %s completed", response.Id)
	return response
}
