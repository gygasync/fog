package work

import "github.com/google/uuid"

type responseDefinition struct {
	id           string
	workId       string
	responseType string
	payload      string
}

func NewResponseDefinition(workDef workDefinition, payload string) *responseDefinition {
	return &responseDefinition{
		id:           uuid.NewString(),
		workId:       workDef.id,
		responseType: workDef.workType,
		payload:      payload,
	}
}
